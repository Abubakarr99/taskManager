// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: taskmanager.proto

package proto

import (
	date "google.golang.org/genproto/googleapis/type/date"
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

// describes the level of urgency of the task
type Urgency int32

const (
	Urgency_VERY_HIGH Urgency = 0
	Urgency_MODERATE  Urgency = 1
	Urgency_LOW       Urgency = 2
)

// Enum value maps for Urgency.
var (
	Urgency_name = map[int32]string{
		0: "VERY_HIGH",
		1: "MODERATE",
		2: "LOW",
	}
	Urgency_value = map[string]int32{
		"VERY_HIGH": 0,
		"MODERATE":  1,
		"LOW":       2,
	}
)

func (x Urgency) Enum() *Urgency {
	p := new(Urgency)
	*p = x
	return p
}

func (x Urgency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Urgency) Descriptor() protoreflect.EnumDescriptor {
	return file_taskmanager_proto_enumTypes[0].Descriptor()
}

func (Urgency) Type() protoreflect.EnumType {
	return &file_taskmanager_proto_enumTypes[0]
}

func (x Urgency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Urgency.Descriptor instead.
func (Urgency) EnumDescriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{0}
}

// Task represents a task
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the id for the task. A UUIDv4 for the task
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the title/,name of the task
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// urgency is the level of urgency of the task
	Urgency Urgency `protobuf:"varint,3,opt,name=urgency,proto3,enum=taskmanager.Urgency" json:"urgency,omitempty"`
	// Due date is when the task is due to be completed
	DueDate *date.Date `protobuf:"bytes,4,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_taskmanager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[0]
	if x != nil {
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
	return file_taskmanager_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetUrgency() Urgency {
	if x != nil {
		return x.Urgency
	}
	return Urgency_VERY_HIGH
}

func (x *Task) GetDueDate() *date.Date {
	if x != nil {
		return x.DueDate
	}
	return nil
}

// Represents a range of dates.
type DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When to start the range, this is inclusive.
	Start *date.Date `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// When to end the range, this is exclusive.
	End *date.Date `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	mi := &file_taskmanager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{1}
}

func (x *DateRange) GetStart() *date.Date {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *DateRange) GetEnd() *date.Date {
	if x != nil {
		return x.End
	}
	return nil
}

// the request used to add a Task to the system
type AddTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *AddTaskReq) Reset() {
	*x = AddTaskReq{}
	mi := &file_taskmanager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskReq) ProtoMessage() {}

func (x *AddTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskReq.ProtoReflect.Descriptor instead.
func (*AddTaskReq) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{2}
}

func (x *AddTaskReq) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// the response to the AddTaskReq
type AddTasksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the IDs of the task to add
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *AddTasksResp) Reset() {
	*x = AddTasksResp{}
	mi := &file_taskmanager_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTasksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTasksResp) ProtoMessage() {}

func (x *AddTasksResp) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTasksResp.ProtoReflect.Descriptor instead.
func (*AddTasksResp) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{3}
}

func (x *AddTasksResp) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// the request to update tasks in the system
type UpdateTasksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *UpdateTasksReq) Reset() {
	*x = UpdateTasksReq{}
	mi := &file_taskmanager_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTasksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTasksReq) ProtoMessage() {}

func (x *UpdateTasksReq) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTasksReq.ProtoReflect.Descriptor instead.
func (*UpdateTasksReq) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTasksReq) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// the response to UpdateTasks()
type UpdateTasksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTasksResp) Reset() {
	*x = UpdateTasksResp{}
	mi := &file_taskmanager_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTasksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTasksResp) ProtoMessage() {}

func (x *UpdateTasksResp) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTasksResp.ProtoReflect.Descriptor instead.
func (*UpdateTasksResp) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{5}
}

// indicates the tasks to delete
type DeleteTasksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteTasksReq) Reset() {
	*x = DeleteTasksReq{}
	mi := &file_taskmanager_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTasksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTasksReq) ProtoMessage() {}

func (x *DeleteTasksReq) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTasksReq.ProtoReflect.Descriptor instead.
func (*DeleteTasksReq) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTasksReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// the response to DeleteTask()
type DeleteTaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTaskResp) Reset() {
	*x = DeleteTaskResp{}
	mi := &file_taskmanager_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResp) ProtoMessage() {}

func (x *DeleteTaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResp.ProtoReflect.Descriptor instead.
func (*DeleteTaskResp) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{7}
}

// the request to search for tasks
type SearchTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filter by urgency level
	Urgency []Urgency `protobuf:"varint,1,rep,packed,name=urgency,proto3,enum=taskmanager.Urgency" json:"urgency,omitempty"`
	// filter by due date
	DueRange *DateRange `protobuf:"bytes,2,opt,name=due_range,json=dueRange,proto3" json:"due_range,omitempty"`
}

func (x *SearchTaskReq) Reset() {
	*x = SearchTaskReq{}
	mi := &file_taskmanager_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTaskReq) ProtoMessage() {}

func (x *SearchTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_taskmanager_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTaskReq.ProtoReflect.Descriptor instead.
func (*SearchTaskReq) Descriptor() ([]byte, []int) {
	return file_taskmanager_proto_rawDescGZIP(), []int{8}
}

func (x *SearchTaskReq) GetUrgency() []Urgency {
	if x != nil {
		return x.Urgency
	}
	return nil
}

func (x *SearchTaskReq) GetDueRange() *DateRange {
	if x != nil {
		return x.DueRange
	}
	return nil
}

var File_taskmanager_proto protoreflect.FileDescriptor

var file_taskmanager_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x07,
	0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x07, 0x64, 0x75,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x59, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x22, 0x35, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x27,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x20, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x39, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x22, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x74, 0x0a,
	0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2e,
	0x0a, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x33,
	0x0a, 0x09, 0x64, 0x75, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x64, 0x75, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x2a, 0x2f, 0x0a, 0x07, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0d,
	0x0a, 0x09, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4c,
	0x4f, 0x57, 0x10, 0x02, 0x32, 0xa8, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x12, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1a, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x62,
	0x75, 0x62, 0x61, 0x6b, 0x61, 0x72, 0x72, 0x39, 0x39, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_taskmanager_proto_rawDescOnce sync.Once
	file_taskmanager_proto_rawDescData = file_taskmanager_proto_rawDesc
)

func file_taskmanager_proto_rawDescGZIP() []byte {
	file_taskmanager_proto_rawDescOnce.Do(func() {
		file_taskmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_taskmanager_proto_rawDescData)
	})
	return file_taskmanager_proto_rawDescData
}

var file_taskmanager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_taskmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_taskmanager_proto_goTypes = []any{
	(Urgency)(0),            // 0: taskmanager.Urgency
	(*Task)(nil),            // 1: taskmanager.Task
	(*DateRange)(nil),       // 2: taskmanager.DateRange
	(*AddTaskReq)(nil),      // 3: taskmanager.AddTaskReq
	(*AddTasksResp)(nil),    // 4: taskmanager.AddTasksResp
	(*UpdateTasksReq)(nil),  // 5: taskmanager.UpdateTasksReq
	(*UpdateTasksResp)(nil), // 6: taskmanager.UpdateTasksResp
	(*DeleteTasksReq)(nil),  // 7: taskmanager.DeleteTasksReq
	(*DeleteTaskResp)(nil),  // 8: taskmanager.DeleteTaskResp
	(*SearchTaskReq)(nil),   // 9: taskmanager.SearchTaskReq
	(*date.Date)(nil),       // 10: google.type.Date
}
var file_taskmanager_proto_depIdxs = []int32{
	0,  // 0: taskmanager.Task.urgency:type_name -> taskmanager.Urgency
	10, // 1: taskmanager.Task.due_date:type_name -> google.type.Date
	10, // 2: taskmanager.DateRange.start:type_name -> google.type.Date
	10, // 3: taskmanager.DateRange.end:type_name -> google.type.Date
	1,  // 4: taskmanager.AddTaskReq.tasks:type_name -> taskmanager.Task
	1,  // 5: taskmanager.UpdateTasksReq.tasks:type_name -> taskmanager.Task
	0,  // 6: taskmanager.SearchTaskReq.urgency:type_name -> taskmanager.Urgency
	2,  // 7: taskmanager.SearchTaskReq.due_range:type_name -> taskmanager.DateRange
	3,  // 8: taskmanager.TaskManager.AddTasks:input_type -> taskmanager.AddTaskReq
	5,  // 9: taskmanager.TaskManager.UpdateTasks:input_type -> taskmanager.UpdateTasksReq
	7,  // 10: taskmanager.TaskManager.DeleteTasks:input_type -> taskmanager.DeleteTasksReq
	9,  // 11: taskmanager.TaskManager.SearchTasks:input_type -> taskmanager.SearchTaskReq
	4,  // 12: taskmanager.TaskManager.AddTasks:output_type -> taskmanager.AddTasksResp
	6,  // 13: taskmanager.TaskManager.UpdateTasks:output_type -> taskmanager.UpdateTasksResp
	8,  // 14: taskmanager.TaskManager.DeleteTasks:output_type -> taskmanager.DeleteTaskResp
	1,  // 15: taskmanager.TaskManager.SearchTasks:output_type -> taskmanager.Task
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_taskmanager_proto_init() }
func file_taskmanager_proto_init() {
	if File_taskmanager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_taskmanager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taskmanager_proto_goTypes,
		DependencyIndexes: file_taskmanager_proto_depIdxs,
		EnumInfos:         file_taskmanager_proto_enumTypes,
		MessageInfos:      file_taskmanager_proto_msgTypes,
	}.Build()
	File_taskmanager_proto = out.File
	file_taskmanager_proto_rawDesc = nil
	file_taskmanager_proto_goTypes = nil
	file_taskmanager_proto_depIdxs = nil
}
