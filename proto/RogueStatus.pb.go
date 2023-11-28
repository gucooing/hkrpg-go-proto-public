// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: RogueStatus.proto

package proto

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

type RogueStatus int32

const (
	RogueStatus_ROGUE_STATUS_NONE    RogueStatus = 0
	RogueStatus_ROGUE_STATUS_DOING   RogueStatus = 1
	RogueStatus_ROGUE_STATUS_PENDING RogueStatus = 2
	RogueStatus_ROGUE_STATUS_ENDLESS RogueStatus = 3
	RogueStatus_ROGUE_STATUS_FINISH  RogueStatus = 4
)

// Enum value maps for RogueStatus.
var (
	RogueStatus_name = map[int32]string{
		0: "ROGUE_STATUS_NONE",
		1: "ROGUE_STATUS_DOING",
		2: "ROGUE_STATUS_PENDING",
		3: "ROGUE_STATUS_ENDLESS",
		4: "ROGUE_STATUS_FINISH",
	}
	RogueStatus_value = map[string]int32{
		"ROGUE_STATUS_NONE":    0,
		"ROGUE_STATUS_DOING":   1,
		"ROGUE_STATUS_PENDING": 2,
		"ROGUE_STATUS_ENDLESS": 3,
		"ROGUE_STATUS_FINISH":  4,
	}
)

func (x RogueStatus) Enum() *RogueStatus {
	p := new(RogueStatus)
	*p = x
	return p
}

func (x RogueStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueStatus_proto_enumTypes[0].Descriptor()
}

func (RogueStatus) Type() protoreflect.EnumType {
	return &file_RogueStatus_proto_enumTypes[0]
}

func (x RogueStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueStatus.Descriptor instead.
func (RogueStatus) EnumDescriptor() ([]byte, []int) {
	return file_RogueStatus_proto_rawDescGZIP(), []int{0}
}

var File_RogueStatus_proto protoreflect.FileDescriptor

var file_RogueStatus_proto_rawDesc = []byte{
	0x0a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x89, 0x01, 0x0a, 0x0b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f,
	0x47, 0x55, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x44,
	0x4c, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x04, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_RogueStatus_proto_rawDescOnce sync.Once
	file_RogueStatus_proto_rawDescData = file_RogueStatus_proto_rawDesc
)

func file_RogueStatus_proto_rawDescGZIP() []byte {
	file_RogueStatus_proto_rawDescOnce.Do(func() {
		file_RogueStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueStatus_proto_rawDescData)
	})
	return file_RogueStatus_proto_rawDescData
}

var file_RogueStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueStatus_proto_goTypes = []interface{}{
	(RogueStatus)(0), // 0: RogueStatus
}
var file_RogueStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueStatus_proto_init() }
func file_RogueStatus_proto_init() {
	if File_RogueStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueStatus_proto_goTypes,
		DependencyIndexes: file_RogueStatus_proto_depIdxs,
		EnumInfos:         file_RogueStatus_proto_enumTypes,
	}.Build()
	File_RogueStatus_proto = out.File
	file_RogueStatus_proto_rawDesc = nil
	file_RogueStatus_proto_goTypes = nil
	file_RogueStatus_proto_depIdxs = nil
}