// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: src/key.proto

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

// Key is the representation of a key in the keyboard to standarize the codes
type Key int32

const (
	Key_ENTER       Key = 0
	Key_ESC         Key = 1
	Key_F11         Key = 2
	Key_LEFT_ARROW  Key = 3
	Key_RIGHT_ARROW Key = 4
	Key_ALT_TAB     Key = 5
)

// Enum value maps for Key.
var (
	Key_name = map[int32]string{
		0: "ENTER",
		1: "ESC",
		2: "F11",
		3: "LEFT_ARROW",
		4: "RIGHT_ARROW",
		5: "ALT_TAB",
	}
	Key_value = map[string]int32{
		"ENTER":       0,
		"ESC":         1,
		"F11":         2,
		"LEFT_ARROW":  3,
		"RIGHT_ARROW": 4,
		"ALT_TAB":     5,
	}
)

func (x Key) Enum() *Key {
	p := new(Key)
	*p = x
	return p
}

func (x Key) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Key) Descriptor() protoreflect.EnumDescriptor {
	return file_src_key_proto_enumTypes[0].Descriptor()
}

func (Key) Type() protoreflect.EnumType {
	return &file_src_key_proto_enumTypes[0]
}

func (x Key) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Key.Descriptor instead.
func (Key) EnumDescriptor() ([]byte, []int) {
	return file_src_key_proto_rawDescGZIP(), []int{0}
}

var File_src_key_proto protoreflect.FileDescriptor

var file_src_key_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x2a, 0x50, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e,
	0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x53, 0x43, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x46, 0x31, 0x31, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x46, 0x54, 0x5f,
	0x41, 0x52, 0x52, 0x4f, 0x57, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x49, 0x47, 0x48, 0x54,
	0x5f, 0x41, 0x52, 0x52, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4c, 0x54, 0x5f,
	0x54, 0x41, 0x42, 0x10, 0x05, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_key_proto_rawDescOnce sync.Once
	file_src_key_proto_rawDescData = file_src_key_proto_rawDesc
)

func file_src_key_proto_rawDescGZIP() []byte {
	file_src_key_proto_rawDescOnce.Do(func() {
		file_src_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_key_proto_rawDescData)
	})
	return file_src_key_proto_rawDescData
}

var file_src_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_key_proto_goTypes = []interface{}{
	(Key)(0), // 0: pb.Key
}
var file_src_key_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_key_proto_init() }
func file_src_key_proto_init() {
	if File_src_key_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_key_proto_goTypes,
		DependencyIndexes: file_src_key_proto_depIdxs,
		EnumInfos:         file_src_key_proto_enumTypes,
	}.Build()
	File_src_key_proto = out.File
	file_src_key_proto_rawDesc = nil
	file_src_key_proto_goTypes = nil
	file_src_key_proto_depIdxs = nil
}
