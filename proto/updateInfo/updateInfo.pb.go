// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.15.5
// source: updateInfo.proto

package updateInfo

import (
	proto "github.com/golang/protobuf/proto"
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

type UpdateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInfo) Reset() {
	*x = UpdateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updateInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfo) ProtoMessage() {}

func (x *UpdateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_updateInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfo.ProtoReflect.Descriptor instead.
func (*UpdateInfo) Descriptor() ([]byte, []int) {
	return file_updateInfo_proto_rawDescGZIP(), []int{0}
}

var File_updateInfo_proto protoreflect.FileDescriptor

var file_updateInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x0d,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_updateInfo_proto_rawDescOnce sync.Once
	file_updateInfo_proto_rawDescData = file_updateInfo_proto_rawDesc
)

func file_updateInfo_proto_rawDescGZIP() []byte {
	file_updateInfo_proto_rawDescOnce.Do(func() {
		file_updateInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_updateInfo_proto_rawDescData)
	})
	return file_updateInfo_proto_rawDescData
}

var file_updateInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_updateInfo_proto_goTypes = []interface{}{
	(*UpdateInfo)(nil), // 0: updateInfo.update_info
}
var file_updateInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_updateInfo_proto_init() }
func file_updateInfo_proto_init() {
	if File_updateInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_updateInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfo); i {
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
			RawDescriptor: file_updateInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_updateInfo_proto_goTypes,
		DependencyIndexes: file_updateInfo_proto_depIdxs,
		MessageInfos:      file_updateInfo_proto_msgTypes,
	}.Build()
	File_updateInfo_proto = out.File
	file_updateInfo_proto_rawDesc = nil
	file_updateInfo_proto_goTypes = nil
	file_updateInfo_proto_depIdxs = nil
}
