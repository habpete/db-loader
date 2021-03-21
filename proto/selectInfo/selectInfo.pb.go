// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.15.5
// source: selectInfo.proto

package selectInfo

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

type SelectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SelectInfo) Reset() {
	*x = SelectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selectInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectInfo) ProtoMessage() {}

func (x *SelectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_selectInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectInfo.ProtoReflect.Descriptor instead.
func (*SelectInfo) Descriptor() ([]byte, []int) {
	return file_selectInfo_proto_rawDescGZIP(), []int{0}
}

type SelectResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SelectResult) Reset() {
	*x = SelectResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selectInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectResult) ProtoMessage() {}

func (x *SelectResult) ProtoReflect() protoreflect.Message {
	mi := &file_selectInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectResult.ProtoReflect.Descriptor instead.
func (*SelectResult) Descriptor() ([]byte, []int) {
	return file_selectInfo_proto_rawDescGZIP(), []int{1}
}

var File_selectInfo_proto protoreflect.FileDescriptor

var file_selectInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x0d,
	0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_selectInfo_proto_rawDescOnce sync.Once
	file_selectInfo_proto_rawDescData = file_selectInfo_proto_rawDesc
)

func file_selectInfo_proto_rawDescGZIP() []byte {
	file_selectInfo_proto_rawDescOnce.Do(func() {
		file_selectInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_selectInfo_proto_rawDescData)
	})
	return file_selectInfo_proto_rawDescData
}

var file_selectInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_selectInfo_proto_goTypes = []interface{}{
	(*SelectInfo)(nil),   // 0: selectInfo.select_info
	(*SelectResult)(nil), // 1: selectInfo.select_result
}
var file_selectInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_selectInfo_proto_init() }
func file_selectInfo_proto_init() {
	if File_selectInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_selectInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectInfo); i {
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
		file_selectInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectResult); i {
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
			RawDescriptor: file_selectInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_selectInfo_proto_goTypes,
		DependencyIndexes: file_selectInfo_proto_depIdxs,
		MessageInfos:      file_selectInfo_proto_msgTypes,
	}.Build()
	File_selectInfo_proto = out.File
	file_selectInfo_proto_rawDesc = nil
	file_selectInfo_proto_goTypes = nil
	file_selectInfo_proto_depIdxs = nil
}
