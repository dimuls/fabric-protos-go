// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: transientstore/transientstore.proto

package transientstore

import (
	rwset "github.com/dimuls/fabric-protos-go/ledger/rwset"
	peer "github.com/dimuls/fabric-protos-go/peer"
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

// TxPvtReadWriteSetWithConfigInfo encapsulates the transaction's private
// read-write set and additional information about the configurations such as
// the latest collection config when the transaction is simulated
type TxPvtReadWriteSetWithConfigInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndorsedAt        uint64                                   `protobuf:"varint,1,opt,name=endorsed_at,json=endorsedAt,proto3" json:"endorsed_at,omitempty"`
	PvtRwset          *rwset.TxPvtReadWriteSet                 `protobuf:"bytes,2,opt,name=pvt_rwset,json=pvtRwset,proto3" json:"pvt_rwset,omitempty"`
	CollectionConfigs map[string]*peer.CollectionConfigPackage `protobuf:"bytes,3,rep,name=collection_configs,json=collectionConfigs,proto3" json:"collection_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TxPvtReadWriteSetWithConfigInfo) Reset() {
	*x = TxPvtReadWriteSetWithConfigInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transientstore_transientstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxPvtReadWriteSetWithConfigInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxPvtReadWriteSetWithConfigInfo) ProtoMessage() {}

func (x *TxPvtReadWriteSetWithConfigInfo) ProtoReflect() protoreflect.Message {
	mi := &file_transientstore_transientstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxPvtReadWriteSetWithConfigInfo.ProtoReflect.Descriptor instead.
func (*TxPvtReadWriteSetWithConfigInfo) Descriptor() ([]byte, []int) {
	return file_transientstore_transientstore_proto_rawDescGZIP(), []int{0}
}

func (x *TxPvtReadWriteSetWithConfigInfo) GetEndorsedAt() uint64 {
	if x != nil {
		return x.EndorsedAt
	}
	return 0
}

func (x *TxPvtReadWriteSetWithConfigInfo) GetPvtRwset() *rwset.TxPvtReadWriteSet {
	if x != nil {
		return x.PvtRwset
	}
	return nil
}

func (x *TxPvtReadWriteSetWithConfigInfo) GetCollectionConfigs() map[string]*peer.CollectionConfigPackage {
	if x != nil {
		return x.CollectionConfigs
	}
	return nil
}

var File_transientstore_transientstore_proto protoreflect.FileDescriptor

var file_transientstore_transientstore_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x18, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x77,
	0x73, 0x65, 0x74, 0x2f, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x1f, 0x54, 0x78, 0x50, 0x76, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x70,
	0x76, 0x74, 0x5f, 0x72, 0x77, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x54, 0x78, 0x50, 0x76, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x52, 0x08, 0x70, 0x76, 0x74, 0x52, 0x77, 0x73,
	0x65, 0x74, 0x12, 0x75, 0x0a, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x54, 0x78, 0x50, 0x76, 0x74, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x65, 0x0a, 0x16, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x61, 0x0a, 0x2c, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6d,
	0x75, 0x6c, 0x73, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transientstore_transientstore_proto_rawDescOnce sync.Once
	file_transientstore_transientstore_proto_rawDescData = file_transientstore_transientstore_proto_rawDesc
)

func file_transientstore_transientstore_proto_rawDescGZIP() []byte {
	file_transientstore_transientstore_proto_rawDescOnce.Do(func() {
		file_transientstore_transientstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_transientstore_transientstore_proto_rawDescData)
	})
	return file_transientstore_transientstore_proto_rawDescData
}

var file_transientstore_transientstore_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transientstore_transientstore_proto_goTypes = []interface{}{
	(*TxPvtReadWriteSetWithConfigInfo)(nil), // 0: transientstore.TxPvtReadWriteSetWithConfigInfo
	nil,                                     // 1: transientstore.TxPvtReadWriteSetWithConfigInfo.CollectionConfigsEntry
	(*rwset.TxPvtReadWriteSet)(nil),         // 2: rwset.TxPvtReadWriteSet
	(*peer.CollectionConfigPackage)(nil),    // 3: protos.CollectionConfigPackage
}
var file_transientstore_transientstore_proto_depIdxs = []int32{
	2, // 0: transientstore.TxPvtReadWriteSetWithConfigInfo.pvt_rwset:type_name -> rwset.TxPvtReadWriteSet
	1, // 1: transientstore.TxPvtReadWriteSetWithConfigInfo.collection_configs:type_name -> transientstore.TxPvtReadWriteSetWithConfigInfo.CollectionConfigsEntry
	3, // 2: transientstore.TxPvtReadWriteSetWithConfigInfo.CollectionConfigsEntry.value:type_name -> protos.CollectionConfigPackage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transientstore_transientstore_proto_init() }
func file_transientstore_transientstore_proto_init() {
	if File_transientstore_transientstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transientstore_transientstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxPvtReadWriteSetWithConfigInfo); i {
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
			RawDescriptor: file_transientstore_transientstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transientstore_transientstore_proto_goTypes,
		DependencyIndexes: file_transientstore_transientstore_proto_depIdxs,
		MessageInfos:      file_transientstore_transientstore_proto_msgTypes,
	}.Build()
	File_transientstore_transientstore_proto = out.File
	file_transientstore_transientstore_proto_rawDesc = nil
	file_transientstore_transientstore_proto_goTypes = nil
	file_transientstore_transientstore_proto_depIdxs = nil
}
