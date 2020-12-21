// protoc --go_out . api/blog.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/blog/blog.proto

package v1

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

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_blog_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PublishReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleID int64 `protobuf:"varint,1,opt,name=articleID,proto3" json:"articleID,omitempty"`
}

func (x *PublishReply) Reset() {
	*x = PublishReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReply) ProtoMessage() {}

func (x *PublishReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReply.ProtoReflect.Descriptor instead.
func (*PublishReply) Descriptor() ([]byte, []int) {
	return file_api_blog_blog_proto_rawDescGZIP(), []int{1}
}

func (x *PublishReply) GetArticleID() int64 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

var File_api_blog_blog_proto protoreflect.FileDescriptor

var file_api_blog_blog_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x44, 0x32, 0x31, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x29, 0x0a,
	0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x0d, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x0d, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x47, 0x6f, 0x2d, 0x30,
	0x30, 0x30, 0x2f, 0x57, 0x65, 0x65, 0x6b, 0x30, 0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_blog_blog_proto_rawDescOnce sync.Once
	file_api_blog_blog_proto_rawDescData = file_api_blog_blog_proto_rawDesc
)

func file_api_blog_blog_proto_rawDescGZIP() []byte {
	file_api_blog_blog_proto_rawDescOnce.Do(func() {
		file_api_blog_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_blog_blog_proto_rawDescData)
	})
	return file_api_blog_blog_proto_rawDescData
}

var file_api_blog_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_blog_blog_proto_goTypes = []interface{}{
	(*PublishRequest)(nil), // 0: PublishRequest
	(*PublishReply)(nil),   // 1: PublishReply
}
var file_api_blog_blog_proto_depIdxs = []int32{
	1, // 0: Blog.Publish:input_type -> PublishReply
	1, // 1: Blog.Publish:output_type -> PublishReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_blog_blog_proto_init() }
func file_api_blog_blog_proto_init() {
	if File_api_blog_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_blog_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_api_blog_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReply); i {
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
			RawDescriptor: file_api_blog_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_blog_blog_proto_goTypes,
		DependencyIndexes: file_api_blog_blog_proto_depIdxs,
		MessageInfos:      file_api_blog_blog_proto_msgTypes,
	}.Build()
	File_api_blog_blog_proto = out.File
	file_api_blog_blog_proto_rawDesc = nil
	file_api_blog_blog_proto_goTypes = nil
	file_api_blog_blog_proto_depIdxs = nil
}
