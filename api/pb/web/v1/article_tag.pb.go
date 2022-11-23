// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: web/v1/article_tag.proto

package web

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

// 文章标签列表接口请求参数
type ArticleTagListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArticleTagListRequest) Reset() {
	*x = ArticleTagListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_article_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTagListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTagListRequest) ProtoMessage() {}

func (x *ArticleTagListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_article_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTagListRequest.ProtoReflect.Descriptor instead.
func (*ArticleTagListRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_article_tag_proto_rawDescGZIP(), []int{0}
}

// 文章标签列表接口响应参数
type ArticleTagListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*ArticleTagListResponse_Item `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ArticleTagListResponse) Reset() {
	*x = ArticleTagListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_article_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTagListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTagListResponse) ProtoMessage() {}

func (x *ArticleTagListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_article_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTagListResponse.ProtoReflect.Descriptor instead.
func (*ArticleTagListResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_article_tag_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleTagListResponse) GetTags() []*ArticleTagListResponse_Item {
	if x != nil {
		return x.Tags
	}
	return nil
}

// 文章标签编辑接口请求参数
type ArticleTagEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId   int32  `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	TagName string `protobuf:"bytes,2,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty" binding:"required"`
}

func (x *ArticleTagEditRequest) Reset() {
	*x = ArticleTagEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_article_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTagEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTagEditRequest) ProtoMessage() {}

func (x *ArticleTagEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_article_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTagEditRequest.ProtoReflect.Descriptor instead.
func (*ArticleTagEditRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_article_tag_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleTagEditRequest) GetTagId() int32 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *ArticleTagEditRequest) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

// 文章标签编辑接口请求参数
type ArticleTagEditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ArticleTagEditResponse) Reset() {
	*x = ArticleTagEditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_article_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTagEditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTagEditResponse) ProtoMessage() {}

func (x *ArticleTagEditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_article_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTagEditResponse.ProtoReflect.Descriptor instead.
func (*ArticleTagEditResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_article_tag_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleTagEditResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 文章标签删除接口请求参数
type ArticleTagDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId int32 `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty" binding:"required"`
}

func (x *ArticleTagDeleteRequest) Reset() {
	*x = ArticleTagDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_article_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTagDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTagDeleteRequest) ProtoMessage() {}

func (x *ArticleTagDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_article_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTagDeleteRequest.ProtoReflect.Descriptor instead.
func (*ArticleTagDeleteRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_article_tag_proto_rawDescGZIP(), []int{4}
}

func (x *ArticleTagDeleteRequest) GetTagId() int32 {
	if x != nil {
		return x.TagId
	}
	return 0
}

// 文章标签删除接口响应参数
type ArticleTagDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArticleTagDeleteResponse) Reset() {
	*x = ArticleTagDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_article_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTagDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTagDeleteResponse) ProtoMessage() {}

func (x *ArticleTagDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_article_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTagDeleteResponse.ProtoReflect.Descriptor instead.
func (*ArticleTagDeleteResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_article_tag_proto_rawDescGZIP(), []int{5}
}

type ArticleTagListResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TagName string `protobuf:"bytes,2,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	Count   int32  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ArticleTagListResponse_Item) Reset() {
	*x = ArticleTagListResponse_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_article_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleTagListResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleTagListResponse_Item) ProtoMessage() {}

func (x *ArticleTagListResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_article_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleTagListResponse_Item.ProtoReflect.Descriptor instead.
func (*ArticleTagListResponse_Item) Descriptor() ([]byte, []int) {
	return file_web_v1_article_tag_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ArticleTagListResponse_Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleTagListResponse_Item) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *ArticleTagListResponse_Item) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_web_v1_article_tag_proto protoreflect.FileDescriptor

var file_web_v1_article_tag_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x77, 0x65, 0x62, 0x1a,
	0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54,
	0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x97, 0x01,
	0x0a, 0x16, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x47,
	0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x15, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x54, 0x61, 0x67, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x22, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x17, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x54, 0x61, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64,
	0x22, 0x1a, 0x0a, 0x18, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a,
	0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_web_v1_article_tag_proto_rawDescOnce sync.Once
	file_web_v1_article_tag_proto_rawDescData = file_web_v1_article_tag_proto_rawDesc
)

func file_web_v1_article_tag_proto_rawDescGZIP() []byte {
	file_web_v1_article_tag_proto_rawDescOnce.Do(func() {
		file_web_v1_article_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_v1_article_tag_proto_rawDescData)
	})
	return file_web_v1_article_tag_proto_rawDescData
}

var file_web_v1_article_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_web_v1_article_tag_proto_goTypes = []interface{}{
	(*ArticleTagListRequest)(nil),       // 0: web.ArticleTagListRequest
	(*ArticleTagListResponse)(nil),      // 1: web.ArticleTagListResponse
	(*ArticleTagEditRequest)(nil),       // 2: web.ArticleTagEditRequest
	(*ArticleTagEditResponse)(nil),      // 3: web.ArticleTagEditResponse
	(*ArticleTagDeleteRequest)(nil),     // 4: web.ArticleTagDeleteRequest
	(*ArticleTagDeleteResponse)(nil),    // 5: web.ArticleTagDeleteResponse
	(*ArticleTagListResponse_Item)(nil), // 6: web.ArticleTagListResponse.Item
}
var file_web_v1_article_tag_proto_depIdxs = []int32{
	6, // 0: web.ArticleTagListResponse.tags:type_name -> web.ArticleTagListResponse.Item
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_v1_article_tag_proto_init() }
func file_web_v1_article_tag_proto_init() {
	if File_web_v1_article_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_v1_article_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTagListRequest); i {
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
		file_web_v1_article_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTagListResponse); i {
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
		file_web_v1_article_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTagEditRequest); i {
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
		file_web_v1_article_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTagEditResponse); i {
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
		file_web_v1_article_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTagDeleteRequest); i {
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
		file_web_v1_article_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTagDeleteResponse); i {
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
		file_web_v1_article_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleTagListResponse_Item); i {
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
			RawDescriptor: file_web_v1_article_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_v1_article_tag_proto_goTypes,
		DependencyIndexes: file_web_v1_article_tag_proto_depIdxs,
		MessageInfos:      file_web_v1_article_tag_proto_msgTypes,
	}.Build()
	File_web_v1_article_tag_proto = out.File
	file_web_v1_article_tag_proto_rawDesc = nil
	file_web_v1_article_tag_proto_goTypes = nil
	file_web_v1_article_tag_proto_depIdxs = nil
}
