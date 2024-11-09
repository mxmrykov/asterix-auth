// protoc --go_out=./internal/proto/gen --go-grpc_out=./internal/proto/gen internal/proto/ast.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: internal/proto/ast.proto

package __

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

type GetIAIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login       string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	ConfirmCode string `protobuf:"bytes,2,opt,name=confirmCode,proto3" json:"confirmCode,omitempty"`
}

func (x *GetIAIDRequest) Reset() {
	*x = GetIAIDRequest{}
	mi := &file_internal_proto_ast_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIAIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIAIDRequest) ProtoMessage() {}

func (x *GetIAIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_ast_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIAIDRequest.ProtoReflect.Descriptor instead.
func (*GetIAIDRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_ast_proto_rawDescGZIP(), []int{0}
}

func (x *GetIAIDRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetIAIDRequest) GetConfirmCode() string {
	if x != nil {
		return x.ConfirmCode
	}
	return ""
}

type GetIAIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Has  bool   `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
	IAID string `protobuf:"bytes,2,opt,name=IAID,proto3" json:"IAID,omitempty"`
	ASID string `protobuf:"bytes,3,opt,name=ASID,proto3" json:"ASID,omitempty"`
}

func (x *GetIAIDResponse) Reset() {
	*x = GetIAIDResponse{}
	mi := &file_internal_proto_ast_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIAIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIAIDResponse) ProtoMessage() {}

func (x *GetIAIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_ast_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIAIDResponse.ProtoReflect.Descriptor instead.
func (*GetIAIDResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_ast_proto_rawDescGZIP(), []int{1}
}

func (x *GetIAIDResponse) GetHas() bool {
	if x != nil {
		return x.Has
	}
	return false
}

func (x *GetIAIDResponse) GetIAID() string {
	if x != nil {
		return x.IAID
	}
	return ""
}

func (x *GetIAIDResponse) GetASID() string {
	if x != nil {
		return x.ASID
	}
	return ""
}

var File_internal_proto_ast_proto protoreflect.FileDescriptor

var file_internal_proto_ast_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x49, 0x41, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x41, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x61, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x41, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x41, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x41, 0x53, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x53, 0x49,
	0x44, 0x32, 0x33, 0x0a, 0x03, 0x41, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x41, 0x49, 0x44, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x41, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x41, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_ast_proto_rawDescOnce sync.Once
	file_internal_proto_ast_proto_rawDescData = file_internal_proto_ast_proto_rawDesc
)

func file_internal_proto_ast_proto_rawDescGZIP() []byte {
	file_internal_proto_ast_proto_rawDescOnce.Do(func() {
		file_internal_proto_ast_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_ast_proto_rawDescData)
	})
	return file_internal_proto_ast_proto_rawDescData
}

var file_internal_proto_ast_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_ast_proto_goTypes = []any{
	(*GetIAIDRequest)(nil),  // 0: GetIAIDRequest
	(*GetIAIDResponse)(nil), // 1: GetIAIDResponse
}
var file_internal_proto_ast_proto_depIdxs = []int32{
	0, // 0: Ast.GetIAID:input_type -> GetIAIDRequest
	1, // 1: Ast.GetIAID:output_type -> GetIAIDResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_ast_proto_init() }
func file_internal_proto_ast_proto_init() {
	if File_internal_proto_ast_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_ast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_ast_proto_goTypes,
		DependencyIndexes: file_internal_proto_ast_proto_depIdxs,
		MessageInfos:      file_internal_proto_ast_proto_msgTypes,
	}.Build()
	File_internal_proto_ast_proto = out.File
	file_internal_proto_ast_proto_rawDesc = nil
	file_internal_proto_ast_proto_goTypes = nil
	file_internal_proto_ast_proto_depIdxs = nil
}