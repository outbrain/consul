// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package types

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"

	"github.com/hashicorp/consul/internal/resource"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func RegisterGRPCRoute(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbmesh.GRPCRouteType,
		Proto:    &pbmesh.GRPCRoute{},
		Scope:    resource.ScopeNamespace,
		Mutate:   MutateGRPCRoute,
		Validate: ValidateGRPCRoute,
		ACLs:     xRouteACLHooks[*pbmesh.GRPCRoute](),
	})
}

func MutateGRPCRoute(res *pbresource.Resource) error {
	var route pbmesh.GRPCRoute

	if err := res.Data.UnmarshalTo(&route); err != nil {
		return resource.NewErrDataParse(&route, err)
	}

	changed := false

	if mutateParentRefs(res.Id.Tenancy, route.ParentRefs) {
		changed = true
	}

	for _, rule := range route.Rules {
		for _, backend := range rule.BackendRefs {
			if backend.BackendRef == nil || backend.BackendRef.Ref == nil {
				continue
			}
			if mutateXRouteRef(res.Id.Tenancy, backend.BackendRef.Ref) {
				changed = true
			}
		}
	}

	if !changed {
		return nil
	}

	return res.Data.MarshalFrom(&route)
}

func ValidateGRPCRoute(res *pbresource.Resource) error {
	var route pbmesh.GRPCRoute

	if err := res.Data.UnmarshalTo(&route); err != nil {
		return resource.NewErrDataParse(&route, err)
	}

	var merr error
	if err := validateParentRefs(res.Id, route.ParentRefs); err != nil {
		merr = multierror.Append(merr, err)
	}

	if len(route.Hostnames) > 0 {
		merr = multierror.Append(merr, resource.ErrInvalidField{
			Name:    "hostnames",
			Wrapped: errors.New("should not populate hostnames"),
		})
	}

	for i, rule := range route.Rules {
		wrapRuleErr := func(err error) error {
			return resource.ErrInvalidListElement{
				Name:    "rules",
				Index:   i,
				Wrapped: err,
			}
		}

		for j, match := range rule.Matches {
			wrapMatchErr := func(err error) error {
				return wrapRuleErr(resource.ErrInvalidListElement{
					Name:    "matches",
					Index:   j,
					Wrapped: err,
				})
			}

			if match.Method != nil {
				switch match.Method.Type {
				case pbmesh.GRPCMethodMatchType_GRPC_METHOD_MATCH_TYPE_UNSPECIFIED:
					merr = multierror.Append(merr, wrapMatchErr(
						resource.ErrInvalidField{
							Name:    "type",
							Wrapped: resource.ErrMissing,
						},
					))
				case pbmesh.GRPCMethodMatchType_GRPC_METHOD_MATCH_TYPE_EXACT:
				case pbmesh.GRPCMethodMatchType_GRPC_METHOD_MATCH_TYPE_REGEX:
				default:
					merr = multierror.Append(merr, wrapMatchErr(
						resource.ErrInvalidField{
							Name:    "type",
							Wrapped: fmt.Errorf("not a supported enum value: %v", match.Method.Type),
						},
					))
				}
				if match.Method.Service == "" && match.Method.Method == "" {
					merr = multierror.Append(merr, wrapMatchErr(
						resource.ErrInvalidField{
							Name:    "service",
							Wrapped: errors.New("at least one of \"service\" or \"method\" must be set"),
						},
					))
				}
			}

			for k, header := range match.Headers {
				wrapMatchHeaderErr := func(err error) error {
					return wrapRuleErr(resource.ErrInvalidListElement{
						Name:    "headers",
						Index:   k,
						Wrapped: err,
					})
				}

				if err := validateHeaderMatchType(header.Type); err != nil {
					merr = multierror.Append(merr, wrapMatchHeaderErr(
						resource.ErrInvalidField{
							Name:    "type",
							Wrapped: err,
						}),
					)
				}

				if header.Name == "" {
					merr = multierror.Append(merr, wrapMatchHeaderErr(
						resource.ErrInvalidField{
							Name:    "name",
							Wrapped: resource.ErrMissing,
						}),
					)
				}
			}
		}

		for j, filter := range rule.Filters {
			wrapFilterErr := func(err error) error {
				return wrapRuleErr(resource.ErrInvalidListElement{
					Name:    "filters",
					Index:   j,
					Wrapped: err,
				})
			}
			set := 0
			if filter.RequestHeaderModifier != nil {
				set++
			}
			if filter.ResponseHeaderModifier != nil {
				set++
			}
			if filter.UrlRewrite != nil {
				set++
				if filter.UrlRewrite.PathPrefix == "" {
					merr = multierror.Append(merr, wrapFilterErr(
						resource.ErrInvalidField{
							Name: "url_rewrite",
							Wrapped: resource.ErrInvalidField{
								Name:    "path_prefix",
								Wrapped: errors.New("field should not be empty if enclosing section is set"),
							},
						},
					))
				}
			}
			if set != 1 {
				merr = multierror.Append(merr, wrapFilterErr(
					errors.New("exactly one of request_header_modifier, response_header_modifier, or url_rewrite is required"),
				))
			}
		}

		if len(rule.BackendRefs) == 0 {
			merr = multierror.Append(merr, wrapRuleErr(
				resource.ErrInvalidField{
					Name:    "backend_refs",
					Wrapped: resource.ErrEmpty,
				},
			))
		}
		for j, hbref := range rule.BackendRefs {
			wrapBackendRefErr := func(err error) error {
				return wrapRuleErr(resource.ErrInvalidListElement{
					Name:    "backend_refs",
					Index:   j,
					Wrapped: err,
				})
			}

			wrapBackendRefFieldErr := func(err error) error {
				return wrapBackendRefErr(resource.ErrInvalidField{
					Name:    "backend_ref",
					Wrapped: err,
				})
			}
			if err := validateBackendRef(hbref.BackendRef, wrapBackendRefFieldErr); err != nil {
				merr = multierror.Append(merr, err)
			}

			if len(hbref.Filters) > 0 {
				merr = multierror.Append(merr, wrapBackendRefErr(
					resource.ErrInvalidField{
						Name:    "filters",
						Wrapped: errors.New("filters are not supported at this level yet"),
					},
				))
			}
		}

		if rule.Timeouts != nil {
			for _, err := range validateHTTPTimeouts(rule.Timeouts) {
				merr = multierror.Append(merr, wrapRuleErr(
					resource.ErrInvalidField{
						Name:    "timeouts",
						Wrapped: err,
					},
				))
			}
		}
		if rule.Retries != nil {
			for _, err := range validateHTTPRetries(rule.Retries) {
				merr = multierror.Append(merr, wrapRuleErr(
					resource.ErrInvalidField{
						Name:    "retries",
						Wrapped: err,
					},
				))
			}
		}
	}

	return merr
}
