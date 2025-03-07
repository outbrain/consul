// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package types

import (
	"math"
	"sort"

	"github.com/hashicorp/go-multierror"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func RegisterWorkload(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbcatalog.WorkloadType,
		Proto:    &pbcatalog.Workload{},
		Scope:    resource.ScopeNamespace,
		Validate: ValidateWorkload,
		ACLs: &resource.ACLHooks{
			Read:  aclReadHookWorkload,
			Write: aclWriteHookWorkload,
			List:  resource.NoOpACLListHook,
		},
	})
}

func ValidateWorkload(res *pbresource.Resource) error {
	var workload pbcatalog.Workload

	if err := res.Data.UnmarshalTo(&workload); err != nil {
		return resource.NewErrDataParse(&workload, err)
	}

	var err error

	// Validate that the workload has at least one port
	if len(workload.Ports) < 1 {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "ports",
			Wrapped: resource.ErrEmpty,
		})
	}

	var meshPorts []string

	// Validate the Workload Ports
	for portName, port := range workload.Ports {
		if portNameErr := ValidatePortName(portName); portNameErr != nil {
			err = multierror.Append(err, resource.ErrInvalidMapKey{
				Map:     "ports",
				Key:     portName,
				Wrapped: portNameErr,
			})
		}

		// disallow port 0 for now
		if port.Port < 1 || port.Port > math.MaxUint16 {
			err = multierror.Append(err, resource.ErrInvalidMapValue{
				Map: "ports",
				Key: portName,
				Wrapped: resource.ErrInvalidField{
					Name:    "port",
					Wrapped: errInvalidPhysicalPort,
				},
			})
		}

		if protoErr := validateProtocol(port.Protocol); protoErr != nil {
			err = multierror.Append(err, resource.ErrInvalidMapValue{
				Map: "ports",
				Key: portName,
				Wrapped: resource.ErrInvalidField{
					Name:    "protocol",
					Wrapped: protoErr,
				},
			})
		}

		// Collect the list of mesh ports
		if port.Protocol == pbcatalog.Protocol_PROTOCOL_MESH {
			meshPorts = append(meshPorts, portName)
		}
	}

	if len(meshPorts) > 1 {
		sort.Strings(meshPorts)
		err = multierror.Append(err, resource.ErrInvalidField{
			Name: "ports",
			Wrapped: errTooMuchMesh{
				Ports: meshPorts,
			},
		})
	}

	// If the workload is mesh enabled then a valid identity must be provided.
	// If not mesh enabled but a non-empty identity is provided then we still
	// validate that its valid.
	if len(meshPorts) > 0 && workload.Identity == "" {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "identity",
			Wrapped: resource.ErrMissing,
		})
	} else if workload.Identity != "" && !isValidDNSLabel(workload.Identity) {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "identity",
			Wrapped: errNotDNSLabel,
		})
	}

	// Validate workload locality
	if workload.Locality != nil && workload.Locality.Region == "" && workload.Locality.Zone != "" {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "locality",
			Wrapped: errLocalityZoneNoRegion,
		})
	}

	// Node associations are optional but if present the name should
	// be a valid DNS label.
	if workload.NodeName != "" {
		if !isValidDNSLabel(workload.NodeName) {
			err = multierror.Append(err, resource.ErrInvalidField{
				Name:    "node_name",
				Wrapped: errNotDNSLabel,
			})
		}
	}

	if len(workload.Addresses) < 1 {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "addresses",
			Wrapped: resource.ErrEmpty,
		})
	}

	// Validate Workload Addresses
	for idx, addr := range workload.Addresses {
		if addrErr := validateWorkloadAddress(addr, workload.Ports); addrErr != nil {
			err = multierror.Append(err, resource.ErrInvalidListElement{
				Name:    "addresses",
				Index:   idx,
				Wrapped: addrErr,
			})
		}
	}

	return err
}

func aclReadHookWorkload(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, id *pbresource.ID, _ *pbresource.Resource) error {
	return authorizer.ToAllowAuthorizer().ServiceReadAllowed(id.GetName(), authzContext)
}

func aclWriteHookWorkload(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, res *pbresource.Resource) error {
	decodedWorkload, err := resource.Decode[*pbcatalog.Workload](res)
	if err != nil {
		return resource.ErrNeedResource
	}

	// First check service:write on the workload name.
	err = authorizer.ToAllowAuthorizer().ServiceWriteAllowed(res.GetId().GetName(), authzContext)
	if err != nil {
		return err
	}

	// Check node:read permissions if node is specified.
	if decodedWorkload.GetData().GetNodeName() != "" {
		return authorizer.ToAllowAuthorizer().NodeReadAllowed(decodedWorkload.GetData().GetNodeName(), authzContext)
	}

	// Check identity:read permissions if identity is specified.
	if decodedWorkload.GetData().GetIdentity() != "" {
		return authorizer.ToAllowAuthorizer().IdentityReadAllowed(decodedWorkload.GetData().GetIdentity(), authzContext)
	}

	return nil
}
