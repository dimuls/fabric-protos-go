// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: peer/configuration.proto

package peer

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

// AnchorPeers simply represents list of anchor peers which is used in ConfigurationItem
type AnchorPeers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnchorPeers []*AnchorPeer `protobuf:"bytes,1,rep,name=anchor_peers,json=anchorPeers,proto3" json:"anchor_peers,omitempty"`
}

func (x *AnchorPeers) Reset() {
	*x = AnchorPeers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnchorPeers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnchorPeers) ProtoMessage() {}

func (x *AnchorPeers) ProtoReflect() protoreflect.Message {
	mi := &file_peer_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnchorPeers.ProtoReflect.Descriptor instead.
func (*AnchorPeers) Descriptor() ([]byte, []int) {
	return file_peer_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *AnchorPeers) GetAnchorPeers() []*AnchorPeer {
	if x != nil {
		return x.AnchorPeers
	}
	return nil
}

// AnchorPeer message structure which provides information about anchor peer, it includes host name,
// port number and peer certificate.
type AnchorPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`  // DNS host name of the anchor peer
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"` // The port number
}

func (x *AnchorPeer) Reset() {
	*x = AnchorPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnchorPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnchorPeer) ProtoMessage() {}

func (x *AnchorPeer) ProtoReflect() protoreflect.Message {
	mi := &file_peer_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnchorPeer.ProtoReflect.Descriptor instead.
func (*AnchorPeer) Descriptor() ([]byte, []int) {
	return file_peer_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *AnchorPeer) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AnchorPeer) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// APIResource represents an API resource in the peer whose ACL
// is determined by the policy_ref field
type APIResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyRef string `protobuf:"bytes,1,opt,name=policy_ref,json=policyRef,proto3" json:"policy_ref,omitempty"` // The policy name to use for this API
}

func (x *APIResource) Reset() {
	*x = APIResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIResource) ProtoMessage() {}

func (x *APIResource) ProtoReflect() protoreflect.Message {
	mi := &file_peer_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIResource.ProtoReflect.Descriptor instead.
func (*APIResource) Descriptor() ([]byte, []int) {
	return file_peer_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *APIResource) GetPolicyRef() string {
	if x != nil {
		return x.PolicyRef
	}
	return ""
}

// ACLs provides mappings for resources in a channel. APIResource encapsulates
// reference to a policy used to determine ACL for the resource
type ACLs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acls map[string]*APIResource `protobuf:"bytes,1,rep,name=acls,proto3" json:"acls,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ACLs) Reset() {
	*x = ACLs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_configuration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLs) ProtoMessage() {}

func (x *ACLs) ProtoReflect() protoreflect.Message {
	mi := &file_peer_configuration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLs.ProtoReflect.Descriptor instead.
func (*ACLs) Descriptor() ([]byte, []int) {
	return file_peer_configuration_proto_rawDescGZIP(), []int{3}
}

func (x *ACLs) GetAcls() map[string]*APIResource {
	if x != nil {
		return x.Acls
	}
	return nil
}

var File_peer_configuration_proto protoreflect.FileDescriptor

var file_peer_configuration_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x22, 0x44, 0x0a, 0x0b, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x12, 0x35, 0x0a, 0x0c, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0b, 0x61, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x41, 0x6e, 0x63, 0x68,
	0x6f, 0x72, 0x50, 0x65, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x2c,
	0x0a, 0x0b, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x66, 0x22, 0x80, 0x01, 0x0a,
	0x04, 0x41, 0x43, 0x4c, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x61, 0x63, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x43, 0x4c,
	0x73, 0x2e, 0x41, 0x63, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x63, 0x6c,
	0x73, 0x1a, 0x4c, 0x0a, 0x09, 0x41, 0x63, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x4d, 0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x70, 0x65, 0x65, 0x72, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x6d, 0x75, 0x6c, 0x73, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_configuration_proto_rawDescOnce sync.Once
	file_peer_configuration_proto_rawDescData = file_peer_configuration_proto_rawDesc
)

func file_peer_configuration_proto_rawDescGZIP() []byte {
	file_peer_configuration_proto_rawDescOnce.Do(func() {
		file_peer_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_configuration_proto_rawDescData)
	})
	return file_peer_configuration_proto_rawDescData
}

var file_peer_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_peer_configuration_proto_goTypes = []interface{}{
	(*AnchorPeers)(nil), // 0: protos.AnchorPeers
	(*AnchorPeer)(nil),  // 1: protos.AnchorPeer
	(*APIResource)(nil), // 2: protos.APIResource
	(*ACLs)(nil),        // 3: protos.ACLs
	nil,                 // 4: protos.ACLs.AclsEntry
}
var file_peer_configuration_proto_depIdxs = []int32{
	1, // 0: protos.AnchorPeers.anchor_peers:type_name -> protos.AnchorPeer
	4, // 1: protos.ACLs.acls:type_name -> protos.ACLs.AclsEntry
	2, // 2: protos.ACLs.AclsEntry.value:type_name -> protos.APIResource
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_peer_configuration_proto_init() }
func file_peer_configuration_proto_init() {
	if File_peer_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnchorPeers); i {
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
		file_peer_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnchorPeer); i {
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
		file_peer_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIResource); i {
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
		file_peer_configuration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACLs); i {
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
			RawDescriptor: file_peer_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_configuration_proto_goTypes,
		DependencyIndexes: file_peer_configuration_proto_depIdxs,
		MessageInfos:      file_peer_configuration_proto_msgTypes,
	}.Build()
	File_peer_configuration_proto = out.File
	file_peer_configuration_proto_rawDesc = nil
	file_peer_configuration_proto_goTypes = nil
	file_peer_configuration_proto_depIdxs = nil
}
