// protoc will be automatically invoked when save.
// see settings.json

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: revp.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamData_Error int32

const (
	StreamData_NO_ERROR StreamData_Error = 0
	StreamData_EOF      StreamData_Error = 1
)

// Enum value maps for StreamData_Error.
var (
	StreamData_Error_name = map[int32]string{
		0: "NO_ERROR",
		1: "EOF",
	}
	StreamData_Error_value = map[string]int32{
		"NO_ERROR": 0,
		"EOF":      1,
	}
)

func (x StreamData_Error) Enum() *StreamData_Error {
	p := new(StreamData_Error)
	*p = x
	return p
}

func (x StreamData_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamData_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_revp_proto_enumTypes[0].Descriptor()
}

func (StreamData_Error) Type() protoreflect.EnumType {
	return &file_revp_proto_enumTypes[0]
}

func (x StreamData_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamData_Error.Descriptor instead.
func (StreamData_Error) EnumDescriptor() ([]byte, []int) {
	return file_revp_proto_rawDescGZIP(), []int{0, 0}
}

type StreamData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret *string           `protobuf:"bytes,1,opt,name=secret,proto3,oneof" json:"secret,omitempty"`
	Data   []byte            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Err    *StreamData_Error `protobuf:"varint,3,opt,name=err,proto3,enum=api.v1alpha1.revp.StreamData_Error,oneof" json:"err,omitempty"`
}

func (x *StreamData) Reset() {
	*x = StreamData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamData) ProtoMessage() {}

func (x *StreamData) ProtoReflect() protoreflect.Message {
	mi := &file_revp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamData.ProtoReflect.Descriptor instead.
func (*StreamData) Descriptor() ([]byte, []int) {
	return file_revp_proto_rawDescGZIP(), []int{0}
}

func (x *StreamData) GetSecret() string {
	if x != nil && x.Secret != nil {
		return *x.Secret
	}
	return ""
}

func (x *StreamData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StreamData) GetErr() StreamData_Error {
	if x != nil && x.Err != nil {
		return *x.Err
	}
	return StreamData_NO_ERROR
}

type StreamExampleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamExampleData) Reset() {
	*x = StreamExampleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamExampleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamExampleData) ProtoMessage() {}

func (x *StreamExampleData) ProtoReflect() protoreflect.Message {
	mi := &file_revp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamExampleData.ProtoReflect.Descriptor instead.
func (*StreamExampleData) Descriptor() ([]byte, []int) {
	return file_revp_proto_rawDescGZIP(), []int{1}
}

func (x *StreamExampleData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count              int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	ElapseMilliseconds int64 `protobuf:"varint,2,opt,name=elapse_milliseconds,json=elapseMilliseconds,proto3" json:"elapse_milliseconds,omitempty"`
}

func (x *StreamReq) Reset() {
	*x = StreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReq) ProtoMessage() {}

func (x *StreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_revp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReq.ProtoReflect.Descriptor instead.
func (*StreamReq) Descriptor() ([]byte, []int) {
	return file_revp_proto_rawDescGZIP(), []int{2}
}

func (x *StreamReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *StreamReq) GetElapseMilliseconds() int64 {
	if x != nil {
		return x.ElapseMilliseconds
	}
	return 0
}

type StreamExampleSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *StreamExampleSummary) Reset() {
	*x = StreamExampleSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamExampleSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamExampleSummary) ProtoMessage() {}

func (x *StreamExampleSummary) ProtoReflect() protoreflect.Message {
	mi := &file_revp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamExampleSummary.ProtoReflect.Descriptor instead.
func (*StreamExampleSummary) Descriptor() ([]byte, []int) {
	return file_revp_proto_rawDescGZIP(), []int{3}
}

func (x *StreamExampleSummary) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_revp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_revp_proto_rawDescGZIP(), []int{4}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_revp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_revp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_revp_proto_rawDescGZIP(), []int{5}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_revp_proto protoreflect.FileDescriptor

var file_revp_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a,
	0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x01,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x88, 0x01, 0x01, 0x22, 0x1e, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x45, 0x4f, 0x46, 0x10, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x65, 0x72, 0x72, 0x22, 0x27, 0x0a, 0x11, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x52, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x65, 0x6c, 0x61, 0x70, 0x73,
	0x65, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26,
	0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x54, 0x0a, 0x04, 0x52, 0x65, 0x76, 0x70, 0x12, 0x4c,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0xb3, 0x02, 0x0a,
	0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x61,
	0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x76, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x56, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72,
	0x65, 0x76, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x12, 0x67, 0x0a, 0x13, 0x42, 0x69, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x72, 0x65, 0x76, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x32, 0x57, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x4c, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x76, 0x70, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6b,
	0x69, 0x64, 0x2f, 0x72, 0x65, 0x76, 0x70, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_revp_proto_rawDescOnce sync.Once
	file_revp_proto_rawDescData = file_revp_proto_rawDesc
)

func file_revp_proto_rawDescGZIP() []byte {
	file_revp_proto_rawDescOnce.Do(func() {
		file_revp_proto_rawDescData = protoimpl.X.CompressGZIP(file_revp_proto_rawDescData)
	})
	return file_revp_proto_rawDescData
}

var file_revp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_revp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_revp_proto_goTypes = []interface{}{
	(StreamData_Error)(0),        // 0: api.v1alpha1.revp.StreamData.Error
	(*StreamData)(nil),           // 1: api.v1alpha1.revp.StreamData
	(*StreamExampleData)(nil),    // 2: api.v1alpha1.revp.StreamExampleData
	(*StreamReq)(nil),            // 3: api.v1alpha1.revp.StreamReq
	(*StreamExampleSummary)(nil), // 4: api.v1alpha1.revp.StreamExampleSummary
	(*HelloRequest)(nil),         // 5: api.v1alpha1.revp.HelloRequest
	(*HelloReply)(nil),           // 6: api.v1alpha1.revp.HelloReply
}
var file_revp_proto_depIdxs = []int32{
	0, // 0: api.v1alpha1.revp.StreamData.err:type_name -> api.v1alpha1.revp.StreamData.Error
	1, // 1: api.v1alpha1.revp.Revp.Stream:input_type -> api.v1alpha1.revp.StreamData
	2, // 2: api.v1alpha1.revp.StreamExample.ClientStream:input_type -> api.v1alpha1.revp.StreamExampleData
	3, // 3: api.v1alpha1.revp.StreamExample.ServerStream:input_type -> api.v1alpha1.revp.StreamReq
	2, // 4: api.v1alpha1.revp.StreamExample.BidirectionalStream:input_type -> api.v1alpha1.revp.StreamExampleData
	5, // 5: api.v1alpha1.revp.Greeter.SayHello:input_type -> api.v1alpha1.revp.HelloRequest
	1, // 6: api.v1alpha1.revp.Revp.Stream:output_type -> api.v1alpha1.revp.StreamData
	4, // 7: api.v1alpha1.revp.StreamExample.ClientStream:output_type -> api.v1alpha1.revp.StreamExampleSummary
	2, // 8: api.v1alpha1.revp.StreamExample.ServerStream:output_type -> api.v1alpha1.revp.StreamExampleData
	2, // 9: api.v1alpha1.revp.StreamExample.BidirectionalStream:output_type -> api.v1alpha1.revp.StreamExampleData
	6, // 10: api.v1alpha1.revp.Greeter.SayHello:output_type -> api.v1alpha1.revp.HelloReply
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_revp_proto_init() }
func file_revp_proto_init() {
	if File_revp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_revp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamData); i {
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
		file_revp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamExampleData); i {
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
		file_revp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReq); i {
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
		file_revp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamExampleSummary); i {
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
		file_revp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_revp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
	file_revp_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_revp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_revp_proto_goTypes,
		DependencyIndexes: file_revp_proto_depIdxs,
		EnumInfos:         file_revp_proto_enumTypes,
		MessageInfos:      file_revp_proto_msgTypes,
	}.Build()
	File_revp_proto = out.File
	file_revp_proto_rawDesc = nil
	file_revp_proto_goTypes = nil
	file_revp_proto_depIdxs = nil
}
