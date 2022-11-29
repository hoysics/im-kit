// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/comet/comet.proto

package comet

import (
	protocol "github.com/hoysics/im-kit/api/protocol"
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

type PushMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  *protocol.Major `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Cids []string        `protobuf:"bytes,3,rep,name=cids,proto3" json:"cids,omitempty"`
}

func (x *PushMsgRequest) Reset() {
	*x = PushMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgRequest) ProtoMessage() {}

func (x *PushMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgRequest.ProtoReflect.Descriptor instead.
func (*PushMsgRequest) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgRequest) GetMsg() *protocol.Major {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *PushMsgRequest) GetCids() []string {
	if x != nil {
		return x.Cids
	}
	return nil
}

type PushMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushMsgReply) Reset() {
	*x = PushMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReply) ProtoMessage() {}

func (x *PushMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReply.ProtoReflect.Descriptor instead.
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{1}
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg   *protocol.Major `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Speed int32           `protobuf:"varint,2,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastRequest) GetMsg() *protocol.Major {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *BroadcastRequest) GetSpeed() int32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

type BroadcastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastReply) Reset() {
	*x = BroadcastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReply) ProtoMessage() {}

func (x *BroadcastReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReply.ProtoReflect.Descriptor instead.
func (*BroadcastReply) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{3}
}

type BroadcastInRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  *protocol.Major `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Room string          `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *BroadcastInRoomRequest) Reset() {
	*x = BroadcastInRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastInRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastInRoomRequest) ProtoMessage() {}

func (x *BroadcastInRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastInRoomRequest.ProtoReflect.Descriptor instead.
func (*BroadcastInRoomRequest) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastInRoomRequest) GetMsg() *protocol.Major {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *BroadcastInRoomRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type BroadcastInRoomReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastInRoomReply) Reset() {
	*x = BroadcastInRoomReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastInRoomReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastInRoomReply) ProtoMessage() {}

func (x *BroadcastInRoomReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastInRoomReply.ProtoReflect.Descriptor instead.
func (*BroadcastInRoomReply) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{5}
}

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid       string  `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Room      string  `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`                                    //access point
	SubTopics []int32 `protobuf:"varint,3,rep,packed,name=sub_topics,json=subTopics,proto3" json:"sub_topics,omitempty"` //msg op subscribed, not proto.op
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_comet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_comet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_api_comet_comet_proto_rawDescGZIP(), []int{6}
}

func (x *ClientInfo) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *ClientInfo) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *ClientInfo) GetSubTopics() []int32 {
	if x != nil {
		return x.SubTopics
	}
	return nil
}

var File_api_comet_comet_proto protoreflect.FileDescriptor

var file_api_comet_comet_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65,
	0x74, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4e, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x52, 0x0a, 0x16, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d,
	0x61, 0x6a, 0x6f, 0x72, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x16, 0x0a,
	0x14, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x51, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62,
	0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x32, 0x3b, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x65,
	0x74, 0x12, 0x32, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x69, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x1a, 0x12, 0x2e,
	0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x61, 0x6a, 0x6f,
	0x72, 0x28, 0x01, 0x30, 0x01, 0x32, 0xdd, 0x01, 0x0a, 0x06, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x2e, 0x69, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a,
	0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x69, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65,
	0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x53, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x20, 0x2e, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74,
	0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3b, 0x0a, 0x0c, 0x69, 0x6d, 0x2d, 0x6b, 0x69, 0x74, 0x2e,
	0x63, 0x6f, 0x6d, 0x65, 0x74, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x79, 0x73, 0x69, 0x63, 0x73, 0x2f, 0x69, 0x6d, 0x2d, 0x6b,
	0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x3b, 0x63, 0x6f, 0x6d,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_comet_comet_proto_rawDescOnce sync.Once
	file_api_comet_comet_proto_rawDescData = file_api_comet_comet_proto_rawDesc
)

func file_api_comet_comet_proto_rawDescGZIP() []byte {
	file_api_comet_comet_proto_rawDescOnce.Do(func() {
		file_api_comet_comet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_comet_comet_proto_rawDescData)
	})
	return file_api_comet_comet_proto_rawDescData
}

var file_api_comet_comet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_comet_comet_proto_goTypes = []interface{}{
	(*PushMsgRequest)(nil),         // 0: im.comet.PushMsgRequest
	(*PushMsgReply)(nil),           // 1: im.comet.PushMsgReply
	(*BroadcastRequest)(nil),       // 2: im.comet.BroadcastRequest
	(*BroadcastReply)(nil),         // 3: im.comet.BroadcastReply
	(*BroadcastInRoomRequest)(nil), // 4: im.comet.BroadcastInRoomRequest
	(*BroadcastInRoomReply)(nil),   // 5: im.comet.BroadcastInRoomReply
	(*ClientInfo)(nil),             // 6: im.comet.ClientInfo
	(*protocol.Major)(nil),         // 7: im.protocol.Major
}
var file_api_comet_comet_proto_depIdxs = []int32{
	7, // 0: im.comet.PushMsgRequest.msg:type_name -> im.protocol.Major
	7, // 1: im.comet.BroadcastRequest.msg:type_name -> im.protocol.Major
	7, // 2: im.comet.BroadcastInRoomRequest.msg:type_name -> im.protocol.Major
	7, // 3: im.comet.Comet.Mail:input_type -> im.protocol.Major
	0, // 4: im.comet.Broker.PushMsg:input_type -> im.comet.PushMsgRequest
	2, // 5: im.comet.Broker.Broadcast:input_type -> im.comet.BroadcastRequest
	4, // 6: im.comet.Broker.BroadcastInRoom:input_type -> im.comet.BroadcastInRoomRequest
	7, // 7: im.comet.Comet.Mail:output_type -> im.protocol.Major
	1, // 8: im.comet.Broker.PushMsg:output_type -> im.comet.PushMsgReply
	3, // 9: im.comet.Broker.Broadcast:output_type -> im.comet.BroadcastReply
	5, // 10: im.comet.Broker.BroadcastInRoom:output_type -> im.comet.BroadcastInRoomReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_comet_comet_proto_init() }
func file_api_comet_comet_proto_init() {
	if File_api_comet_comet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_comet_comet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgRequest); i {
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
		file_api_comet_comet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReply); i {
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
		file_api_comet_comet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRequest); i {
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
		file_api_comet_comet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReply); i {
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
		file_api_comet_comet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastInRoomRequest); i {
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
		file_api_comet_comet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastInRoomReply); i {
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
		file_api_comet_comet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
			RawDescriptor: file_api_comet_comet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_comet_comet_proto_goTypes,
		DependencyIndexes: file_api_comet_comet_proto_depIdxs,
		MessageInfos:      file_api_comet_comet_proto_msgTypes,
	}.Build()
	File_api_comet_comet_proto = out.File
	file_api_comet_comet_proto_rawDesc = nil
	file_api_comet_comet_proto_goTypes = nil
	file_api_comet_comet_proto_depIdxs = nil
}