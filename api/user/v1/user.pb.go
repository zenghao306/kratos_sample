// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.26.0
// source: api/user/v1/user.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserGroupRequest) Reset() {
	*x = GetUserGroupRequest{}
	mi := &file_api_user_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserGroupRequest) ProtoMessage() {}

func (x *GetUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserGroupRequest.ProtoReflect.Descriptor instead.
func (*GetUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserGroupRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GroupIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupId       int64                  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupIdResponse) Reset() {
	*x = GroupIdResponse{}
	mi := &file_api_user_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupIdResponse) ProtoMessage() {}

func (x *GroupIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupIdResponse.ProtoReflect.Descriptor instead.
func (*GroupIdResponse) Descriptor() ([]byte, []int) {
	return file_api_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GroupIdResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_api_user_v1_user_proto protoreflect.FileDescriptor

const file_api_user_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x16api/user/v1/user.proto\x12\vapi.user.v1\".\n" +
	"\x13GetUserGroupRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\",\n" +
	"\x0fGroupIdResponse\x12\x19\n" +
	"\bgroup_id\x18\x01 \x01(\x03R\agroupId2X\n" +
	"\x04User\x12P\n" +
	"\x0eGetUserGroupId\x12 .api.user.v1.GetUserGroupRequest\x1a\x1c.api.user.v1.GroupIdResponseB-\n" +
	"\vapi.user.v1P\x01Z\x1ckratos_sample/api/user/v1;v1b\x06proto3"

var (
	file_api_user_v1_user_proto_rawDescOnce sync.Once
	file_api_user_v1_user_proto_rawDescData []byte
)

func file_api_user_v1_user_proto_rawDescGZIP() []byte {
	file_api_user_v1_user_proto_rawDescOnce.Do(func() {
		file_api_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_user_v1_user_proto_rawDesc), len(file_api_user_v1_user_proto_rawDesc)))
	})
	return file_api_user_v1_user_proto_rawDescData
}

var file_api_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_user_v1_user_proto_goTypes = []any{
	(*GetUserGroupRequest)(nil), // 0: api.user.v1.GetUserGroupRequest
	(*GroupIdResponse)(nil),     // 1: api.user.v1.GroupIdResponse
}
var file_api_user_v1_user_proto_depIdxs = []int32{
	0, // 0: api.user.v1.User.GetUserGroupId:input_type -> api.user.v1.GetUserGroupRequest
	1, // 1: api.user.v1.User.GetUserGroupId:output_type -> api.user.v1.GroupIdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_user_proto_init() }
func file_api_user_v1_user_proto_init() {
	if File_api_user_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_user_v1_user_proto_rawDesc), len(file_api_user_v1_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_v1_user_proto_goTypes,
		DependencyIndexes: file_api_user_v1_user_proto_depIdxs,
		MessageInfos:      file_api_user_v1_user_proto_msgTypes,
	}.Build()
	File_api_user_v1_user_proto = out.File
	file_api_user_v1_user_proto_goTypes = nil
	file_api_user_v1_user_proto_depIdxs = nil
}
