// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protos/control.proto

package playerRpc

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

type RpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *RpcRequest) Reset() {
	*x = RpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcRequest) ProtoMessage() {}

func (x *RpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcRequest.ProtoReflect.Descriptor instead.
func (*RpcRequest) Descriptor() ([]byte, []int) {
	return file_protos_control_proto_rawDescGZIP(), []int{0}
}

func (x *RpcRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type RpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Payload string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RpcResponse) Reset() {
	*x = RpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcResponse) ProtoMessage() {}

func (x *RpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcResponse.ProtoReflect.Descriptor instead.
func (*RpcResponse) Descriptor() ([]byte, []int) {
	return file_protos_control_proto_rawDescGZIP(), []int{1}
}

func (x *RpcResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RpcResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RpcResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_protos_control_proto_rawDescGZIP(), []int{2}
}

type SimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SimpleResponse) Reset() {
	*x = SimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleResponse) ProtoMessage() {}

func (x *SimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleResponse.ProtoReflect.Descriptor instead.
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return file_protos_control_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SimpleResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RpcGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RpcGetRequest) Reset() {
	*x = RpcGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcGetRequest) ProtoMessage() {}

func (x *RpcGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcGetRequest.ProtoReflect.Descriptor instead.
func (*RpcGetRequest) Descriptor() ([]byte, []int) {
	return file_protos_control_proto_rawDescGZIP(), []int{4}
}

func (x *RpcGetRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type RpcGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Payload string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RpcGetResponse) Reset() {
	*x = RpcGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_control_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcGetResponse) ProtoMessage() {}

func (x *RpcGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_control_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcGetResponse.ProtoReflect.Descriptor instead.
func (*RpcGetResponse) Descriptor() ([]byte, []int) {
	return file_protos_control_proto_rawDescGZIP(), []int{5}
}

func (x *RpcGetResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RpcGetResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RpcGetResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_protos_control_proto protoreflect.FileDescriptor

var file_protos_control_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x22,
	0x20, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x4d, 0x0a, 0x0b, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x36, 0x0a, 0x0e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x70, 0x63, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x50, 0x0a, 0x0e, 0x52, 0x70, 0x63, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xe5, 0x01, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x31, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x15, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x52, 0x70, 0x63, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x52,
	0x70, 0x63, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x2f, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_control_proto_rawDescOnce sync.Once
	file_protos_control_proto_rawDescData = file_protos_control_proto_rawDesc
)

func file_protos_control_proto_rawDescGZIP() []byte {
	file_protos_control_proto_rawDescOnce.Do(func() {
		file_protos_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_control_proto_rawDescData)
	})
	return file_protos_control_proto_rawDescData
}

var file_protos_control_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_control_proto_goTypes = []interface{}{
	(*RpcRequest)(nil),     // 0: control.RpcRequest
	(*RpcResponse)(nil),    // 1: control.RpcResponse
	(*EmptyMessage)(nil),   // 2: control.EmptyMessage
	(*SimpleResponse)(nil), // 3: control.SimpleResponse
	(*RpcGetRequest)(nil),  // 4: control.RpcGetRequest
	(*RpcGetResponse)(nil), // 5: control.RpcGetResponse
}
var file_protos_control_proto_depIdxs = []int32{
	0, // 0: control.RpcCall.Call:input_type -> control.RpcRequest
	2, // 1: control.RpcCall.Close:input_type -> control.EmptyMessage
	2, // 2: control.RpcCall.Ping:input_type -> control.EmptyMessage
	4, // 3: control.RpcCall.Get:input_type -> control.RpcGetRequest
	1, // 4: control.RpcCall.Call:output_type -> control.RpcResponse
	3, // 5: control.RpcCall.Close:output_type -> control.SimpleResponse
	3, // 6: control.RpcCall.Ping:output_type -> control.SimpleResponse
	5, // 7: control.RpcCall.Get:output_type -> control.RpcGetResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_control_proto_init() }
func file_protos_control_proto_init() {
	if File_protos_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcRequest); i {
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
		file_protos_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcResponse); i {
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
		file_protos_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
		file_protos_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleResponse); i {
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
		file_protos_control_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcGetRequest); i {
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
		file_protos_control_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcGetResponse); i {
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
			RawDescriptor: file_protos_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_control_proto_goTypes,
		DependencyIndexes: file_protos_control_proto_depIdxs,
		MessageInfos:      file_protos_control_proto_msgTypes,
	}.Build()
	File_protos_control_proto = out.File
	file_protos_control_proto_rawDesc = nil
	file_protos_control_proto_goTypes = nil
	file_protos_control_proto_depIdxs = nil
}