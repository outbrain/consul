// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package types

import (
	"math"

	"github.com/hashicorp/go-multierror"

	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func RegisterDNSPolicy(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbcatalog.DNSPolicyType,
		Proto:    &pbcatalog.DNSPolicy{},
		Scope:    resource.ScopeNamespace,
		Validate: ValidateDNSPolicy,
		ACLs:     ACLHooksForWorkloadSelectingType[*pbcatalog.DNSPolicy](),
	})
}

func ValidateDNSPolicy(res *pbresource.Resource) error {
	var policy pbcatalog.DNSPolicy

	if err := res.Data.UnmarshalTo(&policy); err != nil {
		return resource.NewErrDataParse(&policy, err)
	}

	var err error
	// Ensure that this resource isn't useless and is attempting to
	// select at least one workload.
	if selErr := ValidateSelector(policy.Workloads, false); selErr != nil {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "workloads",
			Wrapped: selErr,
		})
	}

	// Validate the weights
	if weightErr := validateDNSPolicyWeights(policy.Weights); weightErr != nil {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "weights",
			Wrapped: weightErr,
		})
	}

	return err
}

func validateDNSPolicyWeights(weights *pbcatalog.Weights) error {
	// Non nil weights are required
	if weights == nil {
		return resource.ErrMissing
	}

	var err error
	if weights.Passing < 1 || weights.Passing > math.MaxUint16 {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "passing",
			Wrapped: errDNSPassingWeightOutOfRange,
		})
	}

	// Each weight is an unsigned integer so we don't need to
	// check for negative weights.
	if weights.Warning > math.MaxUint16 {
		err = multierror.Append(err, resource.ErrInvalidField{
			Name:    "warning",
			Wrapped: errDNSWarningWeightOutOfRange,
		})
	}

	return err
}
