// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package catalog

import (
	"github.com/hashicorp/consul/internal/catalog/internal/controllers"
	"github.com/hashicorp/consul/internal/catalog/internal/controllers/endpoints"
	"github.com/hashicorp/consul/internal/catalog/internal/controllers/failover"
	"github.com/hashicorp/consul/internal/catalog/internal/controllers/nodehealth"
	"github.com/hashicorp/consul/internal/catalog/internal/controllers/workloadhealth"
	"github.com/hashicorp/consul/internal/catalog/internal/indexers"
	"github.com/hashicorp/consul/internal/catalog/internal/types"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/controller/cache"
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

var (
	// Controller Statuses
	NodeHealthStatusKey              = nodehealth.StatusKey
	NodeHealthStatusConditionHealthy = nodehealth.StatusConditionHealthy
	NodeHealthConditions             = nodehealth.Conditions

	WorkloadHealthStatusKey              = workloadhealth.StatusKey
	WorkloadHealthStatusConditionHealthy = workloadhealth.StatusConditionHealthy
	WorkloadHealthConditions             = workloadhealth.WorkloadConditions
	WorkloadAndNodeHealthConditions      = workloadhealth.NodeAndWorkloadConditions

	EndpointsStatusKey                       = endpoints.StatusKey
	EndpointsStatusConditionEndpointsManaged = endpoints.StatusConditionEndpointsManaged
	EndpointsStatusConditionManaged          = endpoints.ConditionManaged
	EndpointsStatusConditionUnmanaged        = endpoints.ConditionUnmanaged

	FailoverStatusKey                                              = failover.StatusKey
	FailoverStatusConditionAccepted                                = failover.StatusConditionAccepted
	FailoverStatusConditionAcceptedOKReason                        = failover.OKReason
	FailoverStatusConditionAcceptedMissingServiceReason            = failover.MissingServiceReason
	FailoverStatusConditionAcceptedUnknownPortReason               = failover.UnknownPortReason
	FailoverStatusConditionAcceptedMissingDestinationServiceReason = failover.MissingDestinationServiceReason
	FailoverStatusConditionAcceptedUnknownDestinationPortReason    = failover.UnknownDestinationPortReason
	FailoverStatusConditionAcceptedUsingMeshDestinationPortReason  = failover.UsingMeshDestinationPortReason
)

// RegisterTypes adds all resource types within the "catalog" API group
// to the given type registry
func RegisterTypes(r resource.Registry) {
	types.Register(r)
}

// RegisterControllers registers controllers for the catalog types with
// the given controller Manager.
func RegisterControllers(mgr *controller.Manager) {
	controllers.Register(mgr)
}

// SimplifyFailoverPolicy fully populates the PortConfigs map and clears the
// Configs map using the provided Service.
func SimplifyFailoverPolicy(svc *pbcatalog.Service, failover *pbcatalog.FailoverPolicy) *pbcatalog.FailoverPolicy {
	return types.SimplifyFailoverPolicy(svc, failover)
}

// ValidateLocalServiceRefNoSection ensures the following:
//
// - ref is non-nil
// - type is ServiceType
// - section is empty
// - tenancy is set and partition/namespace are both non-empty
// - peer_name must be "local"
//
// Each possible validation error is wrapped in the wrapErr function before
// being collected in a multierror.Error.
func ValidateLocalServiceRefNoSection(ref *pbresource.Reference, wrapErr func(error) error) error {
	return types.ValidateLocalServiceRefNoSection(ref, wrapErr)
}

func FailoverDestinationsIndex() *cache.Index {
	return indexers.FailoverDestinationsIndex()
}
