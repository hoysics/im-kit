// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/protocol/operation.proto

package protocol

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

type Operation int32

const (
	Operation_Heartbeat  Operation = 0
	Operation_ChangeRoom Operation = 1
	Operation_SubTopic   Operation = 2
	Operation_UnsubTopic Operation = 3
	Operation_PassOnMsg  Operation = 4
	Operation_DeliverMsg Operation = 5
	Operation_BeginEvent Operation = 6
	Operation_EndEvent   Operation = 7
	Operation_Disconnect Operation = 8
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "Heartbeat",
		1: "ChangeRoom",
		2: "SubTopic",
		3: "UnsubTopic",
		4: "PassOnMsg",
		5: "DeliverMsg",
		6: "BeginEvent",
		7: "EndEvent",
		8: "Disconnect",
	}
	Operation_value = map[string]int32{
		"Heartbeat":  0,
		"ChangeRoom": 1,
		"SubTopic":   2,
		"UnsubTopic": 3,
		"PassOnMsg":  4,
		"DeliverMsg": 5,
		"BeginEvent": 6,
		"EndEvent":   7,
		"Disconnect": 8,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_api_protocol_operation_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_api_protocol_operation_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_api_protocol_operation_proto_rawDescGZIP(), []int{0}
}

var File_api_protocol_operation_proto protoreflect.FileDescriptor

var file_api_protocol_operation_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2a, 0x95, 0x01, 0x0a, 0x09,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x61, 0x73, 0x73, 0x4f, 0x6e,
	0x4d, 0x73, 0x67, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x10, 0x08, 0x42, 0x40, 0x0a, 0x0b, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x6f, 0x79, 0x73, 0x69, 0x63, 0x73, 0x2f, 0x69, 0x6d, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protocol_operation_proto_rawDescOnce sync.Once
	file_api_protocol_operation_proto_rawDescData = file_api_protocol_operation_proto_rawDesc
)

func file_api_protocol_operation_proto_rawDescGZIP() []byte {
	file_api_protocol_operation_proto_rawDescOnce.Do(func() {
		file_api_protocol_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protocol_operation_proto_rawDescData)
	})
	return file_api_protocol_operation_proto_rawDescData
}

var file_api_protocol_operation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_protocol_operation_proto_goTypes = []interface{}{
	(Operation)(0), // 0: im.protocol.Operation
}
var file_api_protocol_operation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_protocol_operation_proto_init() }
func file_api_protocol_operation_proto_init() {
	if File_api_protocol_operation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_protocol_operation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protocol_operation_proto_goTypes,
		DependencyIndexes: file_api_protocol_operation_proto_depIdxs,
		EnumInfos:         file_api_protocol_operation_proto_enumTypes,
	}.Build()
	File_api_protocol_operation_proto = out.File
	file_api_protocol_operation_proto_rawDesc = nil
	file_api_protocol_operation_proto_goTypes = nil
	file_api_protocol_operation_proto_depIdxs = nil
}
