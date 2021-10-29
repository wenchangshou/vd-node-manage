// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: rpc/server/pb/layout.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Window struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	X         int32             `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y         int32             `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Z         int32             `protobuf:"varint,4,opt,name=z,proto3" json:"z,omitempty"`
	Width     int32             `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height    int32             `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	Service   string            `protobuf:"bytes,7,opt,name=service,proto3" json:"service,omitempty"`
	Arguments map[string]string `protobuf:"bytes,8,rep,name=Arguments,proto3" json:"Arguments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Source    string            `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	Style     map[string]string `protobuf:"bytes,10,rep,name=Style,proto3" json:"Style,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Window) Reset() {
	*x = Window{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_server_pb_layout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Window) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Window) ProtoMessage() {}

func (x *Window) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_server_pb_layout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Window.ProtoReflect.Descriptor instead.
func (*Window) Descriptor() ([]byte, []int) {
	return file_rpc_server_pb_layout_proto_rawDescGZIP(), []int{0}
}

func (x *Window) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Window) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Window) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Window) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Window) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Window) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Window) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Window) GetArguments() map[string]string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Window) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Window) GetStyle() map[string]string {
	if x != nil {
		return x.Style
	}
	return nil
}

type OpenMultiScreenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kill    bool      `protobuf:"varint,2,opt,name=kill,proto3" json:"kill,omitempty"`
	Windows []*Window `protobuf:"bytes,3,rep,name=windows,proto3" json:"windows,omitempty"`
}

func (x *OpenMultiScreenRequest) Reset() {
	*x = OpenMultiScreenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_server_pb_layout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenMultiScreenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenMultiScreenRequest) ProtoMessage() {}

func (x *OpenMultiScreenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_server_pb_layout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenMultiScreenRequest.ProtoReflect.Descriptor instead.
func (*OpenMultiScreenRequest) Descriptor() ([]byte, []int) {
	return file_rpc_server_pb_layout_proto_rawDescGZIP(), []int{1}
}

func (x *OpenMultiScreenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OpenMultiScreenRequest) GetKill() bool {
	if x != nil {
		return x.Kill
	}
	return false
}

func (x *OpenMultiScreenRequest) GetWindows() []*Window {
	if x != nil {
		return x.Windows
	}
	return nil
}

type CloseMultiScreenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CloseMultiScreenRequest) Reset() {
	*x = CloseMultiScreenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_server_pb_layout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseMultiScreenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseMultiScreenRequest) ProtoMessage() {}

func (x *CloseMultiScreenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_server_pb_layout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseMultiScreenRequest.ProtoReflect.Descriptor instead.
func (*CloseMultiScreenRequest) Descriptor() ([]byte, []int) {
	return file_rpc_server_pb_layout_proto_rawDescGZIP(), []int{2}
}

func (x *CloseMultiScreenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_rpc_server_pb_layout_proto protoreflect.FileDescriptor

var file_rpc_server_pb_layout_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x02, 0x0a,
	0x06, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x7a, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x41, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x2e,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x53, 0x74, 0x79, 0x6c,
	0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x38, 0x0a, 0x0a, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5f, 0x0a, 0x16, 0x4f, 0x70, 0x65,
	0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x21, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x22, 0x29, 0x0a, 0x17, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa0, 0x01, 0x0a, 0x0c, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x48,
	0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x12, 0x18, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x73,
	0x68, 0x6f, 0x75, 0x32, 0x2f, 0x76, 0x64, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_server_pb_layout_proto_rawDescOnce sync.Once
	file_rpc_server_pb_layout_proto_rawDescData = file_rpc_server_pb_layout_proto_rawDesc
)

func file_rpc_server_pb_layout_proto_rawDescGZIP() []byte {
	file_rpc_server_pb_layout_proto_rawDescOnce.Do(func() {
		file_rpc_server_pb_layout_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_server_pb_layout_proto_rawDescData)
	})
	return file_rpc_server_pb_layout_proto_rawDescData
}

var file_rpc_server_pb_layout_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_server_pb_layout_proto_goTypes = []interface{}{
	(*Window)(nil),                  // 0: Window
	(*OpenMultiScreenRequest)(nil),  // 1: OpenMultiScreenRequest
	(*CloseMultiScreenRequest)(nil), // 2: CloseMultiScreenRequest
	nil,                             // 3: Window.ArgumentsEntry
	nil,                             // 4: Window.StyleEntry
	(*wrapperspb.BoolValue)(nil),    // 5: google.protobuf.BoolValue
}
var file_rpc_server_pb_layout_proto_depIdxs = []int32{
	3, // 0: Window.Arguments:type_name -> Window.ArgumentsEntry
	4, // 1: Window.Style:type_name -> Window.StyleEntry
	0, // 2: OpenMultiScreenRequest.windows:type_name -> Window
	1, // 3: LayoutManage.OpenMultiScreen:input_type -> OpenMultiScreenRequest
	2, // 4: LayoutManage.CloseMultiScreen:input_type -> CloseMultiScreenRequest
	5, // 5: LayoutManage.OpenMultiScreen:output_type -> google.protobuf.BoolValue
	5, // 6: LayoutManage.CloseMultiScreen:output_type -> google.protobuf.BoolValue
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_server_pb_layout_proto_init() }
func file_rpc_server_pb_layout_proto_init() {
	if File_rpc_server_pb_layout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_server_pb_layout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Window); i {
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
		file_rpc_server_pb_layout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenMultiScreenRequest); i {
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
		file_rpc_server_pb_layout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseMultiScreenRequest); i {
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
			RawDescriptor: file_rpc_server_pb_layout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_server_pb_layout_proto_goTypes,
		DependencyIndexes: file_rpc_server_pb_layout_proto_depIdxs,
		MessageInfos:      file_rpc_server_pb_layout_proto_msgTypes,
	}.Build()
	File_rpc_server_pb_layout_proto = out.File
	file_rpc_server_pb_layout_proto_rawDesc = nil
	file_rpc_server_pb_layout_proto_goTypes = nil
	file_rpc_server_pb_layout_proto_depIdxs = nil
}
