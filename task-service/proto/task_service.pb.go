// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/task_service.proto

package task

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskStatus int32

const (
	TaskStatus_TASK_STATUS_NO_STATUS     TaskStatus = 0
	TaskStatus_TASK_STATUS_NOT_STARTED   TaskStatus = 1
	TaskStatus_TASK_STATUS_STARTED       TaskStatus = 2
	TaskStatus_TASK_STATUS_PENDING       TaskStatus = 3
	TaskStatus_TASK_STATUS_IS_TRIGGERED  TaskStatus = 4
	TaskStatus_TASK_STATUS_NOT_TRIGGERED TaskStatus = 5
	TaskStatus_TASK_STATUS_COMPLETED     TaskStatus = 6
	TaskStatus_TASK_STATUS_IN_PROGRESS   TaskStatus = 7
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "TASK_STATUS_NO_STATUS",
		1: "TASK_STATUS_NOT_STARTED",
		2: "TASK_STATUS_STARTED",
		3: "TASK_STATUS_PENDING",
		4: "TASK_STATUS_IS_TRIGGERED",
		5: "TASK_STATUS_NOT_TRIGGERED",
		6: "TASK_STATUS_COMPLETED",
		7: "TASK_STATUS_IN_PROGRESS",
	}
	TaskStatus_value = map[string]int32{
		"TASK_STATUS_NO_STATUS":     0,
		"TASK_STATUS_NOT_STARTED":   1,
		"TASK_STATUS_STARTED":       2,
		"TASK_STATUS_PENDING":       3,
		"TASK_STATUS_IS_TRIGGERED":  4,
		"TASK_STATUS_NOT_TRIGGERED": 5,
		"TASK_STATUS_COMPLETED":     6,
		"TASK_STATUS_IN_PROGRESS":   7,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_task_service_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_proto_task_service_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{0}
}

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	UserId        int32                  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status        TaskStatus             `protobuf:"varint,5,opt,name=status,proto3,enum=task.TaskStatus" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	TriggerAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=trigger_at,json=triggerAt,proto3" json:"trigger_at,omitempty"`
	Triggered     bool                   `protobuf:"varint,9,opt,name=triggered,proto3" json:"triggered,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_proto_task_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[0]
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
	return file_proto_task_service_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Task) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TASK_STATUS_NO_STATUS
}

func (x *Task) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Task) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Task) GetTriggerAt() *timestamppb.Timestamp {
	if x != nil {
		return x.TriggerAt
	}
	return nil
}

func (x *Task) GetTriggered() bool {
	if x != nil {
		return x.Triggered
	}
	return false
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_proto_task_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskCreated   *Task                  `protobuf:"bytes,1,opt,name=task_created,json=taskCreated,proto3" json:"task_created,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	mi := &file_proto_task_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskResponse) GetTaskCreated() *Task {
	if x != nil {
		return x.TaskCreated
	}
	return nil
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_proto_task_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[3]
	if x != nil {
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
	return file_proto_task_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetTaskRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	mi := &file_proto_task_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UpdatedTask   *Task                  `protobuf:"bytes,1,opt,name=updated_task,json=updatedTask,proto3" json:"updated_task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_proto_task_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTaskRequest) GetUpdatedTask() *Task {
	if x != nil {
		return x.UpdatedTask
	}
	return nil
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	mi := &file_proto_task_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{6}
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_proto_task_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteTaskRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	mi := &file_proto_task_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{8}
}

type DeleteAllTaskOfUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllTaskOfUserRequest) Reset() {
	*x = DeleteAllTaskOfUserRequest{}
	mi := &file_proto_task_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllTaskOfUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllTaskOfUserRequest) ProtoMessage() {}

func (x *DeleteAllTaskOfUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllTaskOfUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllTaskOfUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAllTaskOfUserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteAllTaskOfUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllTaskOfUserResponse) Reset() {
	*x = DeleteAllTaskOfUserResponse{}
	mi := &file_proto_task_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllTaskOfUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllTaskOfUserResponse) ProtoMessage() {}

func (x *DeleteAllTaskOfUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllTaskOfUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteAllTaskOfUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{10}
}

type ListTasksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	UserId        int32                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status        TaskStatus             `protobuf:"varint,4,opt,name=status,proto3,enum=task.TaskStatus" json:"status,omitempty"`
	SortByDate    bool                   `protobuf:"varint,5,opt,name=sort_by_date,json=sortByDate,proto3" json:"sort_by_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksRequest) Reset() {
	*x = ListTasksRequest{}
	mi := &file_proto_task_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksRequest) ProtoMessage() {}

func (x *ListTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksRequest.ProtoReflect.Descriptor instead.
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{11}
}

func (x *ListTasksRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTasksRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTasksRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListTasksRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_TASK_STATUS_NO_STATUS
}

func (x *ListTasksRequest) GetSortByDate() bool {
	if x != nil {
		return x.SortByDate
	}
	return false
}

type ListTasksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksResponse) Reset() {
	*x = ListTasksResponse{}
	mi := &file_proto_task_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksResponse) ProtoMessage() {}

func (x *ListTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksResponse.ProtoReflect.Descriptor instead.
func (*ListTasksResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_proto_rawDescGZIP(), []int{12}
}

func (x *ListTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_proto_task_service_proto protoreflect.FileDescriptor

const file_proto_task_service_proto_rawDesc = "" +
	"\n" +
	"\x18proto/task_service.proto\x12\x04task\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe0\x02\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x17\n" +
	"\auser_id\x18\x04 \x01(\x05R\x06userId\x12(\n" +
	"\x06status\x18\x05 \x01(\x0e2\x10.task.TaskStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x129\n" +
	"\n" +
	"trigger_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\ttriggerAt\x12\x1c\n" +
	"\ttriggered\x18\t \x01(\bR\ttriggered\"3\n" +
	"\x11CreateTaskRequest\x12\x1e\n" +
	"\x04task\x18\x01 \x01(\v2\n" +
	".task.TaskR\x04task\"C\n" +
	"\x12CreateTaskResponse\x12-\n" +
	"\ftask_created\x18\x01 \x01(\v2\n" +
	".task.TaskR\vtaskCreated\"9\n" +
	"\x0eGetTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\"1\n" +
	"\x0fGetTaskResponse\x12\x1e\n" +
	"\x04task\x18\x01 \x01(\v2\n" +
	".task.TaskR\x04task\"B\n" +
	"\x11UpdateTaskRequest\x12-\n" +
	"\fupdated_task\x18\x01 \x01(\v2\n" +
	".task.TaskR\vupdatedTask\"\x14\n" +
	"\x12UpdateTaskResponse\"<\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x05R\x06userId\"\x14\n" +
	"\x12DeleteTaskResponse\"5\n" +
	"\x1aDeleteAllTaskOfUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x05R\x06userId\"\x1d\n" +
	"\x1bDeleteAllTaskOfUserResponse\"\xa1\x01\n" +
	"\x10ListTasksRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\x05R\x06userId\x12(\n" +
	"\x06status\x18\x04 \x01(\x0e2\x10.task.TaskStatusR\x06status\x12 \n" +
	"\fsort_by_date\x18\x05 \x01(\bR\n" +
	"sortByDate\"5\n" +
	"\x11ListTasksResponse\x12 \n" +
	"\x05tasks\x18\x01 \x03(\v2\n" +
	".task.TaskR\x05tasks*\xeb\x01\n" +
	"\n" +
	"TaskStatus\x12\x19\n" +
	"\x15TASK_STATUS_NO_STATUS\x10\x00\x12\x1b\n" +
	"\x17TASK_STATUS_NOT_STARTED\x10\x01\x12\x17\n" +
	"\x13TASK_STATUS_STARTED\x10\x02\x12\x17\n" +
	"\x13TASK_STATUS_PENDING\x10\x03\x12\x1c\n" +
	"\x18TASK_STATUS_IS_TRIGGERED\x10\x04\x12\x1d\n" +
	"\x19TASK_STATUS_NOT_TRIGGERED\x10\x05\x12\x19\n" +
	"\x15TASK_STATUS_COMPLETED\x10\x06\x12\x1b\n" +
	"\x17TASK_STATUS_IN_PROGRESS\x10\a2\xa2\x03\n" +
	"\vTaskService\x12?\n" +
	"\n" +
	"CreateTask\x12\x17.task.CreateTaskRequest\x1a\x18.task.CreateTaskResponse\x126\n" +
	"\aGetTask\x12\x14.task.GetTaskRequest\x1a\x15.task.GetTaskResponse\x12?\n" +
	"\n" +
	"UpdateTask\x12\x17.task.UpdateTaskRequest\x1a\x18.task.UpdateTaskResponse\x12?\n" +
	"\n" +
	"DeleteTask\x12\x17.task.DeleteTaskRequest\x1a\x18.task.DeleteTaskResponse\x12Z\n" +
	"\x13DeleteAllTaskOfUser\x12 .task.DeleteAllTaskOfUserRequest\x1a!.task.DeleteAllTaskOfUserResponse\x12<\n" +
	"\tListTasks\x12\x16.task.ListTasksRequest\x1a\x17.task.ListTasksResponseB(Z&github.com/your-org/shared-proto/;taskb\x06proto3"

var (
	file_proto_task_service_proto_rawDescOnce sync.Once
	file_proto_task_service_proto_rawDescData []byte
)

func file_proto_task_service_proto_rawDescGZIP() []byte {
	file_proto_task_service_proto_rawDescOnce.Do(func() {
		file_proto_task_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_task_service_proto_rawDesc), len(file_proto_task_service_proto_rawDesc)))
	})
	return file_proto_task_service_proto_rawDescData
}

var file_proto_task_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_task_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_task_service_proto_goTypes = []any{
	(TaskStatus)(0),                     // 0: task.TaskStatus
	(*Task)(nil),                        // 1: task.Task
	(*CreateTaskRequest)(nil),           // 2: task.CreateTaskRequest
	(*CreateTaskResponse)(nil),          // 3: task.CreateTaskResponse
	(*GetTaskRequest)(nil),              // 4: task.GetTaskRequest
	(*GetTaskResponse)(nil),             // 5: task.GetTaskResponse
	(*UpdateTaskRequest)(nil),           // 6: task.UpdateTaskRequest
	(*UpdateTaskResponse)(nil),          // 7: task.UpdateTaskResponse
	(*DeleteTaskRequest)(nil),           // 8: task.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),          // 9: task.DeleteTaskResponse
	(*DeleteAllTaskOfUserRequest)(nil),  // 10: task.DeleteAllTaskOfUserRequest
	(*DeleteAllTaskOfUserResponse)(nil), // 11: task.DeleteAllTaskOfUserResponse
	(*ListTasksRequest)(nil),            // 12: task.ListTasksRequest
	(*ListTasksResponse)(nil),           // 13: task.ListTasksResponse
	(*timestamppb.Timestamp)(nil),       // 14: google.protobuf.Timestamp
}
var file_proto_task_service_proto_depIdxs = []int32{
	0,  // 0: task.Task.status:type_name -> task.TaskStatus
	14, // 1: task.Task.created_at:type_name -> google.protobuf.Timestamp
	14, // 2: task.Task.updated_at:type_name -> google.protobuf.Timestamp
	14, // 3: task.Task.trigger_at:type_name -> google.protobuf.Timestamp
	1,  // 4: task.CreateTaskRequest.task:type_name -> task.Task
	1,  // 5: task.CreateTaskResponse.task_created:type_name -> task.Task
	1,  // 6: task.GetTaskResponse.task:type_name -> task.Task
	1,  // 7: task.UpdateTaskRequest.updated_task:type_name -> task.Task
	0,  // 8: task.ListTasksRequest.status:type_name -> task.TaskStatus
	1,  // 9: task.ListTasksResponse.tasks:type_name -> task.Task
	2,  // 10: task.TaskService.CreateTask:input_type -> task.CreateTaskRequest
	4,  // 11: task.TaskService.GetTask:input_type -> task.GetTaskRequest
	6,  // 12: task.TaskService.UpdateTask:input_type -> task.UpdateTaskRequest
	8,  // 13: task.TaskService.DeleteTask:input_type -> task.DeleteTaskRequest
	10, // 14: task.TaskService.DeleteAllTaskOfUser:input_type -> task.DeleteAllTaskOfUserRequest
	12, // 15: task.TaskService.ListTasks:input_type -> task.ListTasksRequest
	3,  // 16: task.TaskService.CreateTask:output_type -> task.CreateTaskResponse
	5,  // 17: task.TaskService.GetTask:output_type -> task.GetTaskResponse
	7,  // 18: task.TaskService.UpdateTask:output_type -> task.UpdateTaskResponse
	9,  // 19: task.TaskService.DeleteTask:output_type -> task.DeleteTaskResponse
	11, // 20: task.TaskService.DeleteAllTaskOfUser:output_type -> task.DeleteAllTaskOfUserResponse
	13, // 21: task.TaskService.ListTasks:output_type -> task.ListTasksResponse
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_task_service_proto_init() }
func file_proto_task_service_proto_init() {
	if File_proto_task_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_task_service_proto_rawDesc), len(file_proto_task_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_service_proto_goTypes,
		DependencyIndexes: file_proto_task_service_proto_depIdxs,
		EnumInfos:         file_proto_task_service_proto_enumTypes,
		MessageInfos:      file_proto_task_service_proto_msgTypes,
	}.Build()
	File_proto_task_service_proto = out.File
	file_proto_task_service_proto_goTypes = nil
	file_proto_task_service_proto_depIdxs = nil
}
