// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.2
// source: event.proto

package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventType     string               `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	AggregateId   string               `protobuf:"bytes,3,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateType string               `protobuf:"bytes,4,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	EventData     []byte               `protobuf:"bytes,5,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	Channel       string               `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Event) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *Event) GetAggregateType() string {
	if x != nil {
		return x.AggregateType
	}
	return ""
}

func (x *Event) GetEventData() []byte {
	if x != nil {
		return x.EventData
	}
	return nil
}

func (x *Event) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Event) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Event) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type RoomCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SagaId string `protobuf:"bytes,1,opt,name=saga_id,json=sagaId,proto3" json:"saga_id,omitempty"`
	Room   *Room  `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *RoomCreated) Reset() {
	*x = RoomCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreated) ProtoMessage() {}

func (x *RoomCreated) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreated.ProtoReflect.Descriptor instead.
func (*RoomCreated) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *RoomCreated) GetSagaId() string {
	if x != nil {
		return x.SagaId
	}
	return ""
}

func (x *RoomCreated) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type CreateRoomFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SagaId  string `protobuf:"bytes,1,opt,name=saga_id,json=sagaId,proto3" json:"saga_id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateRoomFailed) Reset() {
	*x = CreateRoomFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomFailed) ProtoMessage() {}

func (x *CreateRoomFailed) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomFailed.ProtoReflect.Descriptor instead.
func (*CreateRoomFailed) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoomFailed) GetSagaId() string {
	if x != nil {
		return x.SagaId
	}
	return ""
}

func (x *CreateRoomFailed) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SagaId string `protobuf:"bytes,1,opt,name=saga_id,json=sagaId,proto3" json:"saga_id,omitempty"`
	PostId int64  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateRoom) Reset() {
	*x = CreateRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoom) ProtoMessage() {}

func (x *CreateRoom) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoom.ProtoReflect.Descriptor instead.
func (*CreateRoom) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoom) GetSagaId() string {
	if x != nil {
		return x.SagaId
	}
	return ""
}

func (x *CreateRoom) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *CreateRoom) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PostDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *PostDeleted) Reset() {
	*x = PostDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDeleted) ProtoMessage() {}

func (x *PostDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDeleted.ProtoReflect.Descriptor instead.
func (*PostDeleted) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{4}
}

func (x *PostDeleted) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type PostRejected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SagaId       string `protobuf:"bytes,1,opt,name=saga_id,json=sagaId,proto3" json:"saga_id,omitempty"`
	Post         *Post  `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *PostRejected) Reset() {
	*x = PostRejected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRejected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRejected) ProtoMessage() {}

func (x *PostRejected) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRejected.ProtoReflect.Descriptor instead.
func (*PostRejected) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{5}
}

func (x *PostRejected) GetSagaId() string {
	if x != nil {
		return x.SagaId
	}
	return ""
}

func (x *PostRejected) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *PostRejected) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type PostApproved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SagaId string `protobuf:"bytes,1,opt,name=saga_id,json=sagaId,proto3" json:"saga_id,omitempty"`
	Post   *Post  `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *PostApproved) Reset() {
	*x = PostApproved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostApproved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostApproved) ProtoMessage() {}

func (x *PostApproved) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostApproved.ProtoReflect.Descriptor instead.
func (*PostApproved) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{6}
}

func (x *PostApproved) GetSagaId() string {
	if x != nil {
		return x.SagaId
	}
	return ""
}

func (x *PostApproved) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type ApplyPostCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyPost *ApplyPost `protobuf:"bytes,1,opt,name=apply_post,json=applyPost,proto3" json:"apply_post,omitempty"`
}

func (x *ApplyPostCreated) Reset() {
	*x = ApplyPostCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyPostCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyPostCreated) ProtoMessage() {}

func (x *ApplyPostCreated) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyPostCreated.ProtoReflect.Descriptor instead.
func (*ApplyPostCreated) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{7}
}

func (x *ApplyPostCreated) GetApplyPost() *ApplyPost {
	if x != nil {
		return x.ApplyPost
	}
	return nil
}

type ApplyPostDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyPost *ApplyPost `protobuf:"bytes,1,opt,name=apply_post,json=applyPost,proto3" json:"apply_post,omitempty"`
}

func (x *ApplyPostDeleted) Reset() {
	*x = ApplyPostDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyPostDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyPostDeleted) ProtoMessage() {}

func (x *ApplyPostDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyPostDeleted.ProtoReflect.Descriptor instead.
func (*ApplyPostDeleted) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{8}
}

func (x *ApplyPostDeleted) GetApplyPost() *ApplyPost {
	if x != nil {
		return x.ApplyPost
	}
	return nil
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x46, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x67, 0x61, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x67, 0x61, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x22, 0x45, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x67, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x67, 0x61, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x67, 0x61, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x67, 0x61, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73,
	0x74, 0x22, 0x6c, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x67, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x67, 0x61, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x47, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x61, 0x67, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x61, 0x67, 0x61, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x0a,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x10,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x2e, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x6f, 0x73, 0x74,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_event_proto_goTypes = []interface{}{
	(*Event)(nil),               // 0: event.Event
	(*RoomCreated)(nil),         // 1: event.RoomCreated
	(*CreateRoomFailed)(nil),    // 2: event.CreateRoomFailed
	(*CreateRoom)(nil),          // 3: event.CreateRoom
	(*PostDeleted)(nil),         // 4: event.PostDeleted
	(*PostRejected)(nil),        // 5: event.PostRejected
	(*PostApproved)(nil),        // 6: event.PostApproved
	(*ApplyPostCreated)(nil),    // 7: event.ApplyPostCreated
	(*ApplyPostDeleted)(nil),    // 8: event.ApplyPostDeleted
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*Room)(nil),                // 10: chat.Room
	(*Post)(nil),                // 11: post.Post
	(*ApplyPost)(nil),           // 12: post.ApplyPost
}
var file_event_proto_depIdxs = []int32{
	9,  // 0: event.Event.created_at:type_name -> google.protobuf.Timestamp
	9,  // 1: event.Event.updated_at:type_name -> google.protobuf.Timestamp
	10, // 2: event.RoomCreated.room:type_name -> chat.Room
	11, // 3: event.PostDeleted.post:type_name -> post.Post
	11, // 4: event.PostRejected.post:type_name -> post.Post
	11, // 5: event.PostApproved.post:type_name -> post.Post
	12, // 6: event.ApplyPostCreated.apply_post:type_name -> post.ApplyPost
	12, // 7: event.ApplyPostDeleted.apply_post:type_name -> post.ApplyPost
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	file_chat_proto_init()
	file_post_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomCreated); i {
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
		file_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomFailed); i {
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
		file_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoom); i {
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
		file_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostDeleted); i {
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
		file_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRejected); i {
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
		file_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostApproved); i {
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
		file_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyPostCreated); i {
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
		file_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyPostDeleted); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
