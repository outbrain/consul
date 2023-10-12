// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/pbproxystate/traffic_permissions.proto

package pbproxystate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TrafficPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowPermissions []*Permission `protobuf:"bytes,1,rep,name=allow_permissions,json=allowPermissions,proto3" json:"allow_permissions,omitempty"`
	DenyPermissions  []*Permission `protobuf:"bytes,2,rep,name=deny_permissions,json=denyPermissions,proto3" json:"deny_permissions,omitempty"`
	// default_allow determines if the workload is in default allow mode. This is determined
	// by combining the cluster's default allow setting with the is_default property on
	// computed traffic permissions.
	DefaultAllow bool `protobuf:"varint,4,opt,name=default_allow,json=defaultAllow,proto3" json:"default_allow,omitempty"`
}

func (x *TrafficPermissions) Reset() {
	*x = TrafficPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficPermissions) ProtoMessage() {}

func (x *TrafficPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficPermissions.ProtoReflect.Descriptor instead.
func (*TrafficPermissions) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *TrafficPermissions) GetAllowPermissions() []*Permission {
	if x != nil {
		return x.AllowPermissions
	}
	return nil
}

func (x *TrafficPermissions) GetDenyPermissions() []*Permission {
	if x != nil {
		return x.DenyPermissions
	}
	return nil
}

func (x *TrafficPermissions) GetDefaultAllow() bool {
	if x != nil {
		return x.DefaultAllow
	}
	return false
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Principals []*Principal `protobuf:"bytes,1,rep,name=principals,proto3" json:"principals,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetPrincipals() []*Principal {
	if x != nil {
		return x.Principals
	}
	return nil
}

type Principal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spiffe         *Spiffe   `protobuf:"bytes,1,opt,name=spiffe,proto3" json:"spiffe,omitempty"`
	ExcludeSpiffes []*Spiffe `protobuf:"bytes,2,rep,name=exclude_spiffes,json=excludeSpiffes,proto3" json:"exclude_spiffes,omitempty"`
}

func (x *Principal) Reset() {
	*x = Principal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *Principal) GetSpiffe() *Spiffe {
	if x != nil {
		return x.Spiffe
	}
	return nil
}

func (x *Principal) GetExcludeSpiffes() []*Spiffe {
	if x != nil {
		return x.ExcludeSpiffes
	}
	return nil
}

type Spiffe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// regex is the regular expression for matching spiffe ids.
	Regex string `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
	// xfcc_regex specifies that Envoy needs to find the spiffe id in an xfcc header.
	// It is currently unused, but considering this is important for to avoid breaking changes.
	XfccRegex string `protobuf:"bytes,2,opt,name=xfcc_regex,json=xfccRegex,proto3" json:"xfcc_regex,omitempty"`
}

func (x *Spiffe) Reset() {
	*x = Spiffe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spiffe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spiffe) ProtoMessage() {}

func (x *Spiffe) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spiffe.ProtoReflect.Descriptor instead.
func (*Spiffe) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *Spiffe) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

func (x *Spiffe) GetXfccRegex() string {
	if x != nil {
		return x.XfccRegex
	}
	return ""
}

var File_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x81, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x63, 0x0a, 0x11, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x61, 0x0a, 0x10, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x64, 0x65, 0x6e, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x63, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x22, 0xb4, 0x01, 0x0a,
	0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x70,
	0x69, 0x66, 0x66, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x70, 0x69, 0x66, 0x66, 0x65, 0x52, 0x06,
	0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x12, 0x5b, 0x0a, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x70, 0x69,
	0x66, 0x66, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x70, 0x69, 0x66,
	0x66, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x06, 0x53, 0x70, 0x69, 0x66, 0x66, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x78, 0x66, 0x63, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x78, 0x66, 0x63, 0x63, 0x52, 0x65, 0x67,
	0x65, 0x78, 0x42, 0xdd, 0x02, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x17, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0xa2, 0x02, 0x05, 0x48, 0x43, 0x4d, 0x56, 0x50, 0xaa, 0x02,
	0x2a, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50,
	0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0xca, 0x02, 0x2a, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d,
	0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x50, 0x62, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0xe2, 0x02, 0x36, 0x48, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68,
	0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x50, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3a, 0x3a, 0x50, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescData = file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDesc
)

func file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDescData
}

var file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_goTypes = []interface{}{
	(*TrafficPermissions)(nil), // 0: hashicorp.consul.mesh.v2beta1.pbproxystate.TrafficPermissions
	(*Permission)(nil),         // 1: hashicorp.consul.mesh.v2beta1.pbproxystate.Permission
	(*Principal)(nil),          // 2: hashicorp.consul.mesh.v2beta1.pbproxystate.Principal
	(*Spiffe)(nil),             // 3: hashicorp.consul.mesh.v2beta1.pbproxystate.Spiffe
}
var file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_depIdxs = []int32{
	1, // 0: hashicorp.consul.mesh.v2beta1.pbproxystate.TrafficPermissions.allow_permissions:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Permission
	1, // 1: hashicorp.consul.mesh.v2beta1.pbproxystate.TrafficPermissions.deny_permissions:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Permission
	2, // 2: hashicorp.consul.mesh.v2beta1.pbproxystate.Permission.principals:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Principal
	3, // 3: hashicorp.consul.mesh.v2beta1.pbproxystate.Principal.spiffe:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Spiffe
	3, // 4: hashicorp.consul.mesh.v2beta1.pbproxystate.Principal.exclude_spiffes:type_name -> hashicorp.consul.mesh.v2beta1.pbproxystate.Spiffe
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_init() }
func file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_init() {
	if File_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficPermissions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Principal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spiffe); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_depIdxs,
		MessageInfos:      file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_msgTypes,
	}.Build()
	File_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto = out.File
	file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_rawDesc = nil
	file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_goTypes = nil
	file_pbmesh_v2beta1_pbproxystate_traffic_permissions_proto_depIdxs = nil
}
