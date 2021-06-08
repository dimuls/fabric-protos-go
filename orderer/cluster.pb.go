// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: orderer/cluster.proto

package orderer

import (
	common "github.com/dimuls/fabric-protos-go/common"
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

// StepRequest wraps a message that is sent to a cluster member.
type StepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*StepRequest_ConsensusRequest
	//	*StepRequest_SubmitRequest
	Payload isStepRequest_Payload `protobuf_oneof:"payload"`
}

func (x *StepRequest) Reset() {
	*x = StepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepRequest) ProtoMessage() {}

func (x *StepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepRequest.ProtoReflect.Descriptor instead.
func (*StepRequest) Descriptor() ([]byte, []int) {
	return file_orderer_cluster_proto_rawDescGZIP(), []int{0}
}

func (m *StepRequest) GetPayload() isStepRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *StepRequest) GetConsensusRequest() *ConsensusRequest {
	if x, ok := x.GetPayload().(*StepRequest_ConsensusRequest); ok {
		return x.ConsensusRequest
	}
	return nil
}

func (x *StepRequest) GetSubmitRequest() *SubmitRequest {
	if x, ok := x.GetPayload().(*StepRequest_SubmitRequest); ok {
		return x.SubmitRequest
	}
	return nil
}

type isStepRequest_Payload interface {
	isStepRequest_Payload()
}

type StepRequest_ConsensusRequest struct {
	// consensus_request is a consensus specific message.
	ConsensusRequest *ConsensusRequest `protobuf:"bytes,1,opt,name=consensus_request,json=consensusRequest,proto3,oneof"`
}

type StepRequest_SubmitRequest struct {
	// submit_request is a relay of a transaction.
	SubmitRequest *SubmitRequest `protobuf:"bytes,2,opt,name=submit_request,json=submitRequest,proto3,oneof"`
}

func (*StepRequest_ConsensusRequest) isStepRequest_Payload() {}

func (*StepRequest_SubmitRequest) isStepRequest_Payload() {}

// StepResponse is a message received from a cluster member.
type StepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*StepResponse_SubmitRes
	Payload isStepResponse_Payload `protobuf_oneof:"payload"`
}

func (x *StepResponse) Reset() {
	*x = StepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepResponse) ProtoMessage() {}

func (x *StepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepResponse.ProtoReflect.Descriptor instead.
func (*StepResponse) Descriptor() ([]byte, []int) {
	return file_orderer_cluster_proto_rawDescGZIP(), []int{1}
}

func (m *StepResponse) GetPayload() isStepResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *StepResponse) GetSubmitRes() *SubmitResponse {
	if x, ok := x.GetPayload().(*StepResponse_SubmitRes); ok {
		return x.SubmitRes
	}
	return nil
}

type isStepResponse_Payload interface {
	isStepResponse_Payload()
}

type StepResponse_SubmitRes struct {
	SubmitRes *SubmitResponse `protobuf:"bytes,1,opt,name=submit_res,json=submitRes,proto3,oneof"`
}

func (*StepResponse_SubmitRes) isStepResponse_Payload() {}

// ConsensusRequest is a consensus specific message sent to a cluster member.
type ConsensusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel  string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Payload  []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ConsensusRequest) Reset() {
	*x = ConsensusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusRequest) ProtoMessage() {}

func (x *ConsensusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusRequest.ProtoReflect.Descriptor instead.
func (*ConsensusRequest) Descriptor() ([]byte, []int) {
	return file_orderer_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ConsensusRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *ConsensusRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ConsensusRequest) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// SubmitRequest wraps a transaction to be sent for ordering.
type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// last_validation_seq denotes the last
	// configuration sequence at which the
	// sender validated this message.
	LastValidationSeq uint64 `protobuf:"varint,2,opt,name=last_validation_seq,json=lastValidationSeq,proto3" json:"last_validation_seq,omitempty"`
	// content is the fabric transaction
	// that is forwarded to the cluster member.
	Payload *common.Envelope `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_orderer_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *SubmitRequest) GetLastValidationSeq() uint64 {
	if x != nil {
		return x.LastValidationSeq
	}
	return 0
}

func (x *SubmitRequest) GetPayload() *common.Envelope {
	if x != nil {
		return x.Payload
	}
	return nil
}

// SubmitResponse returns a success
// or failure status to the sender.
type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// Status code, which may be used to programatically respond to success/failure.
	Status common.Status `protobuf:"varint,2,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	// Info string which may contain additional information about the returned status.
	Info string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_orderer_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitResponse) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *SubmitResponse) GetStatus() common.Status {
	if x != nil {
		return x.Status
	}
	return common.Status_UNKNOWN
}

func (x *SubmitResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_orderer_cluster_proto protoreflect.FileDescriptor

var file_orderer_cluster_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3f, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x53, 0x0a, 0x0c, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x62, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x85, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6c,
	0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x71,
	0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x66, 0x0a, 0x0e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0x42, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x37, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x53, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e,
	0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x6d, 0x75, 0x6c, 0x73, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_cluster_proto_rawDescOnce sync.Once
	file_orderer_cluster_proto_rawDescData = file_orderer_cluster_proto_rawDesc
)

func file_orderer_cluster_proto_rawDescGZIP() []byte {
	file_orderer_cluster_proto_rawDescOnce.Do(func() {
		file_orderer_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_cluster_proto_rawDescData)
	})
	return file_orderer_cluster_proto_rawDescData
}

var file_orderer_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orderer_cluster_proto_goTypes = []interface{}{
	(*StepRequest)(nil),      // 0: orderer.StepRequest
	(*StepResponse)(nil),     // 1: orderer.StepResponse
	(*ConsensusRequest)(nil), // 2: orderer.ConsensusRequest
	(*SubmitRequest)(nil),    // 3: orderer.SubmitRequest
	(*SubmitResponse)(nil),   // 4: orderer.SubmitResponse
	(*common.Envelope)(nil),  // 5: common.Envelope
	(common.Status)(0),       // 6: common.Status
}
var file_orderer_cluster_proto_depIdxs = []int32{
	2, // 0: orderer.StepRequest.consensus_request:type_name -> orderer.ConsensusRequest
	3, // 1: orderer.StepRequest.submit_request:type_name -> orderer.SubmitRequest
	4, // 2: orderer.StepResponse.submit_res:type_name -> orderer.SubmitResponse
	5, // 3: orderer.SubmitRequest.payload:type_name -> common.Envelope
	6, // 4: orderer.SubmitResponse.status:type_name -> common.Status
	0, // 5: orderer.Cluster.Step:input_type -> orderer.StepRequest
	1, // 6: orderer.Cluster.Step:output_type -> orderer.StepResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orderer_cluster_proto_init() }
func file_orderer_cluster_proto_init() {
	if File_orderer_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepRequest); i {
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
		file_orderer_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepResponse); i {
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
		file_orderer_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusRequest); i {
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
		file_orderer_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_orderer_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
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
	file_orderer_cluster_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StepRequest_ConsensusRequest)(nil),
		(*StepRequest_SubmitRequest)(nil),
	}
	file_orderer_cluster_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StepResponse_SubmitRes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderer_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderer_cluster_proto_goTypes,
		DependencyIndexes: file_orderer_cluster_proto_depIdxs,
		MessageInfos:      file_orderer_cluster_proto_msgTypes,
	}.Build()
	File_orderer_cluster_proto = out.File
	file_orderer_cluster_proto_rawDesc = nil
	file_orderer_cluster_proto_goTypes = nil
	file_orderer_cluster_proto_depIdxs = nil
}
