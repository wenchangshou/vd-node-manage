// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: pb/task.proto

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

type TaskType int32

const (
	TaskType_INITIALIZE TaskType = 0
	TaskType_EXECUTE    TaskType = 1
	TaskType_DONE       TaskType = 2
	TaskType_ERROR      TaskType = 3
	TaskType_ALL        TaskType = 4
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "INITIALIZE",
		1: "EXECUTE",
		2: "DONE",
		3: "ERROR",
		4: "ALL",
	}
	TaskType_value = map[string]int32{
		"INITIALIZE": 0,
		"EXECUTE":    1,
		"DONE":       2,
		"ERROR":      3,
		"ALL":        4,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_task_proto_enumTypes[0].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_pb_task_proto_enumTypes[0]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_pb_task_proto_rawDescGZIP(), []int{0}
}

type TaskOperatorType int32

const (
	TaskOperatorType_INSTALL_PROJECT  TaskOperatorType = 0
	TaskOperatorType_INSTALL_RESOURCE TaskOperatorType = 1
	TaskOperatorType_UPGRADE_PROJECT  TaskOperatorType = 2
	TaskOperatorType_DELETE_RESOURCE  TaskOperatorType = 3
	TaskOperatorType_DELETE_PROJECT   TaskOperatorType = 4
	TaskOperatorType_UNKNOWN          TaskOperatorType = 5
)

// Enum value maps for TaskOperatorType.
var (
	TaskOperatorType_name = map[int32]string{
		0: "INSTALL_PROJECT",
		1: "INSTALL_RESOURCE",
		2: "UPGRADE_PROJECT",
		3: "DELETE_RESOURCE",
		4: "DELETE_PROJECT",
		5: "UNKNOWN",
	}
	TaskOperatorType_value = map[string]int32{
		"INSTALL_PROJECT":  0,
		"INSTALL_RESOURCE": 1,
		"UPGRADE_PROJECT":  2,
		"DELETE_RESOURCE":  3,
		"DELETE_PROJECT":   4,
		"UNKNOWN":          5,
	}
)

func (x TaskOperatorType) Enum() *TaskOperatorType {
	p := new(TaskOperatorType)
	*p = x
	return p
}

func (x TaskOperatorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskOperatorType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_task_proto_enumTypes[1].Descriptor()
}

func (TaskOperatorType) Type() protoreflect.EnumType {
	return &file_pb_task_proto_enumTypes[1]
}

func (x TaskOperatorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskOperatorType.Descriptor instead.
func (TaskOperatorType) EnumDescriptor() ([]byte, []int) {
	return file_pb_task_proto_rawDescGZIP(), []int{1}
}

type SetTaskStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type TaskType `protobuf:"varint,2,opt,name=type,proto3,enum=TaskType" json:"type,omitempty"`
}

func (x *SetTaskStatusRequest) Reset() {
	*x = SetTaskStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTaskStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTaskStatusRequest) ProtoMessage() {}

func (x *SetTaskStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTaskStatusRequest.ProtoReflect.Descriptor instead.
func (*SetTaskStatusRequest) Descriptor() ([]byte, []int) {
	return file_pb_task_proto_rawDescGZIP(), []int{0}
}

func (x *SetTaskStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetTaskStatusRequest) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_INITIALIZE
}

type TasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Task `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TasksResponse) Reset() {
	*x = TasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksResponse) ProtoMessage() {}

func (x *TasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksResponse.ProtoReflect.Descriptor instead.
func (*TasksResponse) Descriptor() ([]byte, []int) {
	return file_pb_task_proto_rawDescGZIP(), []int{1}
}

func (x *TasksResponse) GetItems() []*Task {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac      string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	TaskType TaskType `protobuf:"varint,2,opt,name=taskType,proto3,enum=TaskType" json:"taskType,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_pb_task_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaskRequest) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *GetTaskRequest) GetTaskType() TaskType {
	if x != nil {
		return x.TaskType
	}
	return TaskType_INITIALIZE
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Options  string           `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	Action   TaskOperatorType `protobuf:"varint,3,opt,name=action,proto3,enum=TaskOperatorType" json:"action,omitempty"`
	Status   int32            `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Depend   int32            `protobuf:"varint,5,opt,name=Depend,proto3" json:"Depend,omitempty"`
	Schedule int32            `protobuf:"varint,6,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Active   bool             `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_pb_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_pb_task_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Task) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

func (x *Task) GetAction() TaskOperatorType {
	if x != nil {
		return x.Action
	}
	return TaskOperatorType_INSTALL_PROJECT
}

func (x *Task) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Task) GetDepend() int32 {
	if x != nil {
		return x.Depend
	}
	return 0
}

func (x *Task) GetSchedule() int32 {
	if x != nil {
		return x.Schedule
	}
	return 0
}

func (x *Task) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

var File_pb_task_proto protoreflect.FileDescriptor

var file_pb_task_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x45, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x25, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x22,
	0xbf, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x2a, 0x45, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f,
	0x4e, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x04, 0x2a, 0x88, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73,
	0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x50, 0x47, 0x52,
	0x41, 0x44, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f,
	0x4a, 0x45, 0x43, 0x54, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x05, 0x32, 0x8d, 0x01, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x12, 0x0f,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x0d, 0x73, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x15, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x65, 0x6e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x73, 0x68, 0x6f, 0x75, 0x32, 0x2f,
	0x76, 0x64, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_task_proto_rawDescOnce sync.Once
	file_pb_task_proto_rawDescData = file_pb_task_proto_rawDesc
)

func file_pb_task_proto_rawDescGZIP() []byte {
	file_pb_task_proto_rawDescOnce.Do(func() {
		file_pb_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_task_proto_rawDescData)
	})
	return file_pb_task_proto_rawDescData
}

var file_pb_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_task_proto_goTypes = []interface{}{
	(TaskType)(0),                // 0: TaskType
	(TaskOperatorType)(0),        // 1: TaskOperatorType
	(*SetTaskStatusRequest)(nil), // 2: SetTaskStatusRequest
	(*TasksResponse)(nil),        // 3: TasksResponse
	(*GetTaskRequest)(nil),       // 4: GetTaskRequest
	(*Task)(nil),                 // 5: Task
	(*wrapperspb.BoolValue)(nil), // 6: google.protobuf.BoolValue
}
var file_pb_task_proto_depIdxs = []int32{
	0, // 0: SetTaskStatusRequest.type:type_name -> TaskType
	5, // 1: TasksResponse.items:type_name -> Task
	0, // 2: GetTaskRequest.taskType:type_name -> TaskType
	1, // 3: Task.action:type_name -> TaskOperatorType
	4, // 4: TaskManagement.getTaskByComputerMac:input_type -> GetTaskRequest
	2, // 5: TaskManagement.setTaskStatus:input_type -> SetTaskStatusRequest
	3, // 6: TaskManagement.getTaskByComputerMac:output_type -> TasksResponse
	6, // 7: TaskManagement.setTaskStatus:output_type -> google.protobuf.BoolValue
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_task_proto_init() }
func file_pb_task_proto_init() {
	if File_pb_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTaskStatusRequest); i {
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
		file_pb_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksResponse); i {
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
		file_pb_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_pb_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
			RawDescriptor: file_pb_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_task_proto_goTypes,
		DependencyIndexes: file_pb_task_proto_depIdxs,
		EnumInfos:         file_pb_task_proto_enumTypes,
		MessageInfos:      file_pb_task_proto_msgTypes,
	}.Build()
	File_pb_task_proto = out.File
	file_pb_task_proto_rawDesc = nil
	file_pb_task_proto_goTypes = nil
	file_pb_task_proto_depIdxs = nil
}
