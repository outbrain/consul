// Code generated by protoc-gen-resource-types. DO NOT EDIT.

package catalogv2beta1

import (
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	GroupName = "catalog"
	Version   = "v2beta1"
)

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.DNSPolicy
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const DNSPolicyKind = "DNSPolicy"

var DNSPolicyType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         DNSPolicyKind,
}

func (_ *DNSPolicy) GetResourceType() *pbresource.Type {
	return DNSPolicyType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.FailoverPolicy
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const FailoverPolicyKind = "FailoverPolicy"

var FailoverPolicyType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         FailoverPolicyKind,
}

func (_ *FailoverPolicy) GetResourceType() *pbresource.Type {
	return FailoverPolicyType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.HealthChecks
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const HealthChecksKind = "HealthChecks"

var HealthChecksType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         HealthChecksKind,
}

func (_ *HealthChecks) GetResourceType() *pbresource.Type {
	return HealthChecksType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.HealthStatus
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const HealthStatusKind = "HealthStatus"

var HealthStatusType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         HealthStatusKind,
}

func (_ *HealthStatus) GetResourceType() *pbresource.Type {
	return HealthStatusType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.Node
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const NodeKind = "Node"

var NodeType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         NodeKind,
}

func (_ *Node) GetResourceType() *pbresource.Type {
	return NodeType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.Service
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const ServiceKind = "Service"

var ServiceType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         ServiceKind,
}

func (_ *Service) GetResourceType() *pbresource.Type {
	return ServiceType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.ServiceEndpoints
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const ServiceEndpointsKind = "ServiceEndpoints"

var ServiceEndpointsType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         ServiceEndpointsKind,
}

func (_ *ServiceEndpoints) GetResourceType() *pbresource.Type {
	return ServiceEndpointsType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.VirtualIPs
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const VirtualIPsKind = "VirtualIPs"

var VirtualIPsType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         VirtualIPsKind,
}

func (_ *VirtualIPs) GetResourceType() *pbresource.Type {
	return VirtualIPsType
}

/* ---------------------------------------------------------------------------
 * hashicorp.consul.catalog.v2beta1.Workload
 *
 * This following section contains constants variables and utility methods
 * for interacting with this kind of resource.
 * -------------------------------------------------------------------------*/

const WorkloadKind = "Workload"

var WorkloadType = &pbresource.Type{
	Group:        GroupName,
	GroupVersion: Version,
	Kind:         WorkloadKind,
}

func (_ *Workload) GetResourceType() *pbresource.Type {
	return WorkloadType
}
