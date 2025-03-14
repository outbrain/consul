// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package types

import (
	"github.com/hashicorp/go-multierror"

	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func RegisterHealthChecks(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbcatalog.HealthChecksType,
		Proto:    &pbcatalog.HealthChecks{},
		Scope:    resource.ScopeNamespace,
		Validate: ValidateHealthChecks,
		ACLs:     ACLHooksForWorkloadSelectingType[*pbcatalog.HealthChecks](),
	})
}

func ValidateHealthChecks(res *pbresource.Resource) error {
	var checks pbcatalog.HealthChecks

	if err := res.Data.UnmarshalTo(&checks); err != nil {
		return resource.NewErrDataParse(&checks, err)
	}

	var err error

	// Validate the workload selector
	if selErr := ValidateSelector(checks.Workloads, false); selErr != nil {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "workloads",
			Wrapped: selErr,
		})
	}

	// Validate each check
	for idx, check := range checks.HealthChecks {
		if checkErr := validateCheck(check); checkErr != nil {
			err = multierror.Append(err, resource.ErrInvalidListElement{
				Name:    "checks",
				Index:   idx,
				Wrapped: checkErr,
			})
		}
	}

	return err
}

func validateCheck(check *pbcatalog.HealthCheck) error {
	var err error
	// Validate the check name
	if check.Name == "" {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "name",
			Wrapped: resource.ErrMissing,
		})
	} else if !isValidDNSLabel(check.Name) {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "name",
			Wrapped: errNotDNSLabel,
		})
	}

	// Validate the definition
	if check.Definition == nil {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "definition",
			Wrapped: resource.ErrMissing,
		})
	}

	// In theory it would be nice to validate the individual check definition.
	// However whether a check is valid will be up for interpretation by
	// the check executor. The executor may default some addressing,
	// allow for templating etc. Therefore we cannot really know at admission
	// time whether the check will be executable. Therefore it is expected
	// that check executors will update the status of the resource to note
	// whether it was valid for that executor.

	return err
}
