// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internalapi/counter_service.proto

package internalapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDialogCountersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogId     string   `protobuf:"bytes,1,opt,name=dialog_id,json=dialogId,proto3" json:"dialog_id,omitempty"`
	Participants []string `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *CreateDialogCountersRequest) Reset() {
	*x = CreateDialogCountersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_counter_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDialogCountersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDialogCountersRequest) ProtoMessage() {}

func (x *CreateDialogCountersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_counter_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDialogCountersRequest.ProtoReflect.Descriptor instead.
func (*CreateDialogCountersRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_counter_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDialogCountersRequest) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

func (x *CreateDialogCountersRequest) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

type IncreaseDialogCountersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogId    string   `protobuf:"bytes,1,opt,name=dialog_id,json=dialogId,proto3" json:"dialog_id,omitempty"`
	SenderId    string   `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	FollowersId []string `protobuf:"bytes,3,rep,name=followers_id,json=followersId,proto3" json:"followers_id,omitempty"`
}

func (x *IncreaseDialogCountersRequest) Reset() {
	*x = IncreaseDialogCountersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_counter_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseDialogCountersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseDialogCountersRequest) ProtoMessage() {}

func (x *IncreaseDialogCountersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_counter_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseDialogCountersRequest.ProtoReflect.Descriptor instead.
func (*IncreaseDialogCountersRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_counter_service_proto_rawDescGZIP(), []int{1}
}

func (x *IncreaseDialogCountersRequest) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

func (x *IncreaseDialogCountersRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *IncreaseDialogCountersRequest) GetFollowersId() []string {
	if x != nil {
		return x.FollowersId
	}
	return nil
}

type FlushDialogCountersForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogId string `protobuf:"bytes,1,opt,name=dialog_id,json=dialogId,proto3" json:"dialog_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FlushDialogCountersForUserRequest) Reset() {
	*x = FlushDialogCountersForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_counter_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushDialogCountersForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushDialogCountersForUserRequest) ProtoMessage() {}

func (x *FlushDialogCountersForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_counter_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushDialogCountersForUserRequest.ProtoReflect.Descriptor instead.
func (*FlushDialogCountersForUserRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_counter_service_proto_rawDescGZIP(), []int{2}
}

func (x *FlushDialogCountersForUserRequest) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

func (x *FlushDialogCountersForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetDialogCounterForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogId string `protobuf:"bytes,1,opt,name=dialog_id,json=dialogId,proto3" json:"dialog_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetDialogCounterForUserRequest) Reset() {
	*x = GetDialogCounterForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_counter_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDialogCounterForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDialogCounterForUserRequest) ProtoMessage() {}

func (x *GetDialogCounterForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_counter_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDialogCounterForUserRequest.ProtoReflect.Descriptor instead.
func (*GetDialogCounterForUserRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_counter_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetDialogCounterForUserRequest) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

func (x *GetDialogCounterForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetDialogCounterForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter int64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *GetDialogCounterForUserResponse) Reset() {
	*x = GetDialogCounterForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_counter_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDialogCounterForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDialogCounterForUserResponse) ProtoMessage() {}

func (x *GetDialogCounterForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_counter_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDialogCounterForUserResponse.ProtoReflect.Descriptor instead.
func (*GetDialogCounterForUserResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_counter_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetDialogCounterForUserResponse) GetCounter() int64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

var File_internalapi_counter_service_proto protoreflect.FileDescriptor

var file_internalapi_counter_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x7c, 0x0a, 0x1d, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x21, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x32, 0xd5, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x30,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x16, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x6e, 0x0a, 0x1a, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x44, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x36, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x33,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x74, 0x68, 0x30, 0x6c,
	0x65, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_counter_service_proto_rawDescOnce sync.Once
	file_internalapi_counter_service_proto_rawDescData = file_internalapi_counter_service_proto_rawDesc
)

func file_internalapi_counter_service_proto_rawDescGZIP() []byte {
	file_internalapi_counter_service_proto_rawDescOnce.Do(func() {
		file_internalapi_counter_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_counter_service_proto_rawDescData)
	})
	return file_internalapi_counter_service_proto_rawDescData
}

var file_internalapi_counter_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internalapi_counter_service_proto_goTypes = []interface{}{
	(*CreateDialogCountersRequest)(nil),       // 0: counter.internalapi.CreateDialogCountersRequest
	(*IncreaseDialogCountersRequest)(nil),     // 1: counter.internalapi.IncreaseDialogCountersRequest
	(*FlushDialogCountersForUserRequest)(nil), // 2: counter.internalapi.FlushDialogCountersForUserRequest
	(*GetDialogCounterForUserRequest)(nil),    // 3: counter.internalapi.GetDialogCounterForUserRequest
	(*GetDialogCounterForUserResponse)(nil),   // 4: counter.internalapi.GetDialogCounterForUserResponse
	(*emptypb.Empty)(nil),                     // 5: google.protobuf.Empty
}
var file_internalapi_counter_service_proto_depIdxs = []int32{
	0, // 0: counter.internalapi.CounterService.CreateDialogCounters:input_type -> counter.internalapi.CreateDialogCountersRequest
	1, // 1: counter.internalapi.CounterService.IncreaseDialogCounters:input_type -> counter.internalapi.IncreaseDialogCountersRequest
	2, // 2: counter.internalapi.CounterService.FlushDialogCountersForUser:input_type -> counter.internalapi.FlushDialogCountersForUserRequest
	3, // 3: counter.internalapi.CounterService.GetDialogCounterForUser:input_type -> counter.internalapi.GetDialogCounterForUserRequest
	5, // 4: counter.internalapi.CounterService.CreateDialogCounters:output_type -> google.protobuf.Empty
	5, // 5: counter.internalapi.CounterService.IncreaseDialogCounters:output_type -> google.protobuf.Empty
	5, // 6: counter.internalapi.CounterService.FlushDialogCountersForUser:output_type -> google.protobuf.Empty
	4, // 7: counter.internalapi.CounterService.GetDialogCounterForUser:output_type -> counter.internalapi.GetDialogCounterForUserResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internalapi_counter_service_proto_init() }
func file_internalapi_counter_service_proto_init() {
	if File_internalapi_counter_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalapi_counter_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDialogCountersRequest); i {
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
		file_internalapi_counter_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseDialogCountersRequest); i {
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
		file_internalapi_counter_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushDialogCountersForUserRequest); i {
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
		file_internalapi_counter_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDialogCounterForUserRequest); i {
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
		file_internalapi_counter_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDialogCounterForUserResponse); i {
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
			RawDescriptor: file_internalapi_counter_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_counter_service_proto_goTypes,
		DependencyIndexes: file_internalapi_counter_service_proto_depIdxs,
		MessageInfos:      file_internalapi_counter_service_proto_msgTypes,
	}.Build()
	File_internalapi_counter_service_proto = out.File
	file_internalapi_counter_service_proto_rawDesc = nil
	file_internalapi_counter_service_proto_goTypes = nil
	file_internalapi_counter_service_proto_depIdxs = nil
}
