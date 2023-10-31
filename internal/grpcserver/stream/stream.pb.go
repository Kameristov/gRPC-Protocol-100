// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: stream/stream.proto

package stream

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_stream_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_stream_stream_proto_rawDescGZIP(), []int{0}
}

// The response message
type ResponseScale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type    string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Subtype string `protobuf:"bytes,4,opt,name=subtype,proto3" json:"subtype,omitempty"`
}

func (x *ResponseScale) Reset() {
	*x = ResponseScale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseScale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseScale) ProtoMessage() {}

func (x *ResponseScale) ProtoReflect() protoreflect.Message {
	mi := &file_stream_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseScale.ProtoReflect.Descriptor instead.
func (*ResponseScale) Descriptor() ([]byte, []int) {
	return file_stream_stream_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseScale) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ResponseScale) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseScale) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResponseScale) GetSubtype() string {
	if x != nil {
		return x.Subtype
	}
	return ""
}

type RequestTareValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestTareValue) Reset() {
	*x = RequestTareValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTareValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTareValue) ProtoMessage() {}

func (x *RequestTareValue) ProtoReflect() protoreflect.Message {
	mi := &file_stream_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTareValue.ProtoReflect.Descriptor instead.
func (*RequestTareValue) Descriptor() ([]byte, []int) {
	return file_stream_stream_proto_rawDescGZIP(), []int{2}
}

func (x *RequestTareValue) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RequestScale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Subtype string `protobuf:"bytes,3,opt,name=subtype,proto3" json:"subtype,omitempty"`
}

func (x *RequestScale) Reset() {
	*x = RequestScale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestScale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestScale) ProtoMessage() {}

func (x *RequestScale) ProtoReflect() protoreflect.Message {
	mi := &file_stream_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestScale.ProtoReflect.Descriptor instead.
func (*RequestScale) Descriptor() ([]byte, []int) {
	return file_stream_stream_proto_rawDescGZIP(), []int{3}
}

func (x *RequestScale) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RequestScale) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RequestScale) GetSubtype() string {
	if x != nil {
		return x.Subtype
	}
	return ""
}

type ResponseSetScale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResponseSetScale) Reset() {
	*x = ResponseSetScale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseSetScale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseSetScale) ProtoMessage() {}

func (x *ResponseSetScale) ProtoReflect() protoreflect.Message {
	mi := &file_stream_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseSetScale.ProtoReflect.Descriptor instead.
func (*ResponseSetScale) Descriptor() ([]byte, []int) {
	return file_stream_stream_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseSetScale) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ResponseInstantWeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseInstantWeight) Reset() {
	*x = ResponseInstantWeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_stream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInstantWeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInstantWeight) ProtoMessage() {}

func (x *ResponseInstantWeight) ProtoReflect() protoreflect.Message {
	mi := &file_stream_stream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInstantWeight.ProtoReflect.Descriptor instead.
func (*ResponseInstantWeight) Descriptor() ([]byte, []int) {
	return file_stream_stream_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseInstantWeight) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ResponseInstantWeight) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_stream_stream_proto protoreflect.FileDescriptor

var file_stream_stream_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6d, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x61, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x56, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x47, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x88,
	0x03, 0x0a, 0x0e, 0x41, 0x70, 0x69, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x12, 0x4c, 0x0a, 0x17, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x75, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x34, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x74, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x54, 0x61, 0x72, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x18, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x53,
	0x65, 0x74, 0x5a, 0x65, 0x72, 0x6f, 0x12, 0x0d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_stream_proto_rawDescOnce sync.Once
	file_stream_stream_proto_rawDescData = file_stream_stream_proto_rawDesc
)

func file_stream_stream_proto_rawDescGZIP() []byte {
	file_stream_stream_proto_rawDescOnce.Do(func() {
		file_stream_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_stream_proto_rawDescData)
	})
	return file_stream_stream_proto_rawDescData
}

var file_stream_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stream_stream_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: stream.Empty
	(*ResponseScale)(nil),         // 1: stream.ResponseScale
	(*RequestTareValue)(nil),      // 2: stream.RequestTareValue
	(*RequestScale)(nil),          // 3: stream.RequestScale
	(*ResponseSetScale)(nil),      // 4: stream.ResponseSetScale
	(*ResponseInstantWeight)(nil), // 5: stream.ResponseInstantWeight
}
var file_stream_stream_proto_depIdxs = []int32{
	3, // 0: stream.ApiCallerScale.ScalesMessageOutChannel:input_type -> stream.RequestScale
	0, // 1: stream.ApiCallerScale.SetTare:input_type -> stream.Empty
	2, // 2: stream.ApiCallerScale.SetTareValue:input_type -> stream.RequestTareValue
	0, // 3: stream.ApiCallerScale.SetZero:input_type -> stream.Empty
	0, // 4: stream.ApiCallerScale.GetInstantWeight:input_type -> stream.Empty
	0, // 5: stream.ApiCallerScale.GetState:input_type -> stream.Empty
	1, // 6: stream.ApiCallerScale.ScalesMessageOutChannel:output_type -> stream.ResponseScale
	4, // 7: stream.ApiCallerScale.SetTare:output_type -> stream.ResponseSetScale
	4, // 8: stream.ApiCallerScale.SetTareValue:output_type -> stream.ResponseSetScale
	4, // 9: stream.ApiCallerScale.SetZero:output_type -> stream.ResponseSetScale
	5, // 10: stream.ApiCallerScale.GetInstantWeight:output_type -> stream.ResponseInstantWeight
	1, // 11: stream.ApiCallerScale.GetState:output_type -> stream.ResponseScale
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_stream_proto_init() }
func file_stream_stream_proto_init() {
	if File_stream_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_stream_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseScale); i {
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
		file_stream_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTareValue); i {
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
		file_stream_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestScale); i {
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
		file_stream_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseSetScale); i {
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
		file_stream_stream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInstantWeight); i {
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
			RawDescriptor: file_stream_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_stream_proto_goTypes,
		DependencyIndexes: file_stream_stream_proto_depIdxs,
		MessageInfos:      file_stream_stream_proto_msgTypes,
	}.Build()
	File_stream_stream_proto = out.File
	file_stream_stream_proto_rawDesc = nil
	file_stream_stream_proto_goTypes = nil
	file_stream_stream_proto_depIdxs = nil
}
