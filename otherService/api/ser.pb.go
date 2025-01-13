// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: api/ser.proto

package api

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

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	mi := &file_api_ser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ser_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_api_ser_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *MessageRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response string          `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_api_ser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ser_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_api_ser_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponse) GetRequest() *MessageRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *MessageResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type MessageKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageKey) Reset() {
	*x = MessageKey{}
	mi := &file_api_ser_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageKey) ProtoMessage() {}

func (x *MessageKey) ProtoReflect() protoreflect.Message {
	mi := &file_api_ser_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageKey.ProtoReflect.Descriptor instead.
func (*MessageKey) Descriptor() ([]byte, []int) {
	return file_api_ser_proto_rawDescGZIP(), []int{2}
}

func (x *MessageKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MessageKey) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Keys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []*MessageKey `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
}

func (x *Keys) Reset() {
	*x = Keys{}
	mi := &file_api_ser_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Keys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keys) ProtoMessage() {}

func (x *Keys) ProtoReflect() protoreflect.Message {
	mi := &file_api_ser_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keys.ProtoReflect.Descriptor instead.
func (*Keys) Descriptor() ([]byte, []int) {
	return file_api_ser_proto_rawDescGZIP(), []int{3}
}

func (x *Keys) GetKey() []*MessageKey {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_api_ser_proto protoreflect.FileDescriptor

var file_api_ser_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x3e, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x38, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x04,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4b,
	0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0xec, 0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4b, 0x65, 0x79, 0x73,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x33, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4b, 0x65, 0x79, 0x73,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6e, 0x65, 0x78, 0x6c, 0x65, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ser_proto_rawDescOnce sync.Once
	file_api_ser_proto_rawDescData = file_api_ser_proto_rawDesc
)

func file_api_ser_proto_rawDescGZIP() []byte {
	file_api_ser_proto_rawDescOnce.Do(func() {
		file_api_ser_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ser_proto_rawDescData)
	})
	return file_api_ser_proto_rawDescData
}

var file_api_ser_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_ser_proto_goTypes = []any{
	(*MessageRequest)(nil),  // 0: api.MessageRequest
	(*MessageResponse)(nil), // 1: api.MessageResponse
	(*MessageKey)(nil),      // 2: api.MessageKey
	(*Keys)(nil),            // 3: api.Keys
}
var file_api_ser_proto_depIdxs = []int32{
	0, // 0: api.MessageResponse.request:type_name -> api.MessageRequest
	2, // 1: api.Keys.key:type_name -> api.MessageKey
	0, // 2: api.RouteGuide.GetMessage:input_type -> api.MessageRequest
	2, // 3: api.RouteGuide.GetMessageLists:input_type -> api.MessageKey
	0, // 4: api.RouteGuide.AddMessage:input_type -> api.MessageRequest
	2, // 5: api.RouteGuide.AddMessageLists:input_type -> api.MessageKey
	1, // 6: api.RouteGuide.GetMessage:output_type -> api.MessageResponse
	3, // 7: api.RouteGuide.GetMessageLists:output_type -> api.Keys
	1, // 8: api.RouteGuide.AddMessage:output_type -> api.MessageResponse
	3, // 9: api.RouteGuide.AddMessageLists:output_type -> api.Keys
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ser_proto_init() }
func file_api_ser_proto_init() {
	if File_api_ser_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_ser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ser_proto_goTypes,
		DependencyIndexes: file_api_ser_proto_depIdxs,
		MessageInfos:      file_api_ser_proto_msgTypes,
	}.Build()
	File_api_ser_proto = out.File
	file_api_ser_proto_rawDesc = nil
	file_api_ser_proto_goTypes = nil
	file_api_ser_proto_depIdxs = nil
}