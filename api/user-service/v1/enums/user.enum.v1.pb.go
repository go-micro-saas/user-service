// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.6
// source: api/user-service/v1/enums/user.enum.v1.proto

package enumv1

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

type UserInitEnum_UserInit int32

const (
	UserInitEnum_UNSPECIFIED UserInitEnum_UserInit = 0 // 未指定
)

// Enum value maps for UserInitEnum_UserInit.
var (
	UserInitEnum_UserInit_name = map[int32]string{
		0: "UNSPECIFIED",
	}
	UserInitEnum_UserInit_value = map[string]int32{
		"UNSPECIFIED": 0,
	}
)

func (x UserInitEnum_UserInit) Enum() *UserInitEnum_UserInit {
	p := new(UserInitEnum_UserInit)
	*p = x
	return p
}

func (x UserInitEnum_UserInit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInitEnum_UserInit) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[0].Descriptor()
}

func (UserInitEnum_UserInit) Type() protoreflect.EnumType {
	return &file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[0]
}

func (x UserInitEnum_UserInit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInitEnum_UserInit.Descriptor instead.
func (UserInitEnum_UserInit) EnumDescriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{0, 0}
}

type UserStatusEnum_UserStatus int32

const (
	UserStatusEnum_INITIAL   UserStatusEnum_UserStatus = 0 // 初始状态
	UserStatusEnum_ENABLE    UserStatusEnum_UserStatus = 1 // 启用
	UserStatusEnum_DISABLE   UserStatusEnum_UserStatus = 2 // 禁用
	UserStatusEnum_WHITELIST UserStatusEnum_UserStatus = 3 // 白名单
	UserStatusEnum_BLACKLIST UserStatusEnum_UserStatus = 4 // 黑名单
	UserStatusEnum_DELETED   UserStatusEnum_UserStatus = 5 // 已删除
)

// Enum value maps for UserStatusEnum_UserStatus.
var (
	UserStatusEnum_UserStatus_name = map[int32]string{
		0: "INITIAL",
		1: "ENABLE",
		2: "DISABLE",
		3: "WHITELIST",
		4: "BLACKLIST",
		5: "DELETED",
	}
	UserStatusEnum_UserStatus_value = map[string]int32{
		"INITIAL":   0,
		"ENABLE":    1,
		"DISABLE":   2,
		"WHITELIST": 3,
		"BLACKLIST": 4,
		"DELETED":   5,
	}
)

func (x UserStatusEnum_UserStatus) Enum() *UserStatusEnum_UserStatus {
	p := new(UserStatusEnum_UserStatus)
	*p = x
	return p
}

func (x UserStatusEnum_UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatusEnum_UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[1].Descriptor()
}

func (UserStatusEnum_UserStatus) Type() protoreflect.EnumType {
	return &file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[1]
}

func (x UserStatusEnum_UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatusEnum_UserStatus.Descriptor instead.
func (UserStatusEnum_UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{1, 0}
}

type UserGenderEnum_UserGender int32

const (
	UserGenderEnum_UNSPECIFIED UserGenderEnum_UserGender = 0 // 未指定
	UserGenderEnum_MALE        UserGenderEnum_UserGender = 1 // 男
	UserGenderEnum_FEMALE      UserGenderEnum_UserGender = 2 // 女
	UserGenderEnum_SECRET      UserGenderEnum_UserGender = 3 // 秘密
)

// Enum value maps for UserGenderEnum_UserGender.
var (
	UserGenderEnum_UserGender_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "MALE",
		2: "FEMALE",
		3: "SECRET",
	}
	UserGenderEnum_UserGender_value = map[string]int32{
		"UNSPECIFIED": 0,
		"MALE":        1,
		"FEMALE":      2,
		"SECRET":      3,
	}
)

func (x UserGenderEnum_UserGender) Enum() *UserGenderEnum_UserGender {
	p := new(UserGenderEnum_UserGender)
	*p = x
	return p
}

func (x UserGenderEnum_UserGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserGenderEnum_UserGender) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[2].Descriptor()
}

func (UserGenderEnum_UserGender) Type() protoreflect.EnumType {
	return &file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[2]
}

func (x UserGenderEnum_UserGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserGenderEnum_UserGender.Descriptor instead.
func (UserGenderEnum_UserGender) EnumDescriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{2, 0}
}

type UserRegisterTypeEnum_UserRegisterType int32

const (
	UserRegisterTypeEnum_UNSPECIFIED    UserRegisterTypeEnum_UserRegisterType = 0    // 未指定
	UserRegisterTypeEnum_EMAIL          UserRegisterTypeEnum_UserRegisterType = 1    // 邮箱
	UserRegisterTypeEnum_PHONE          UserRegisterTypeEnum_UserRegisterType = 2    // 手机
	UserRegisterTypeEnum_GOOGLE_OAUTH   UserRegisterTypeEnum_UserRegisterType = 101  // GOOGLE_OAUTH
	UserRegisterTypeEnum_GITHUB_OAUTH   UserRegisterTypeEnum_UserRegisterType = 102  // GITHUB_OAUTH
	UserRegisterTypeEnum_JULIANG_OAUTH  UserRegisterTypeEnum_UserRegisterType = 103  // 巨量OAUTH
	UserRegisterTypeEnum_DINGTALK_OAUTH UserRegisterTypeEnum_UserRegisterType = 104  // 钉钉OAUTH
	UserRegisterTypeEnum_THIRD_PARTY    UserRegisterTypeEnum_UserRegisterType = 2001 // 第三方
)

// Enum value maps for UserRegisterTypeEnum_UserRegisterType.
var (
	UserRegisterTypeEnum_UserRegisterType_name = map[int32]string{
		0:    "UNSPECIFIED",
		1:    "EMAIL",
		2:    "PHONE",
		101:  "GOOGLE_OAUTH",
		102:  "GITHUB_OAUTH",
		103:  "JULIANG_OAUTH",
		104:  "DINGTALK_OAUTH",
		2001: "THIRD_PARTY",
	}
	UserRegisterTypeEnum_UserRegisterType_value = map[string]int32{
		"UNSPECIFIED":    0,
		"EMAIL":          1,
		"PHONE":          2,
		"GOOGLE_OAUTH":   101,
		"GITHUB_OAUTH":   102,
		"JULIANG_OAUTH":  103,
		"DINGTALK_OAUTH": 104,
		"THIRD_PARTY":    2001,
	}
)

func (x UserRegisterTypeEnum_UserRegisterType) Enum() *UserRegisterTypeEnum_UserRegisterType {
	p := new(UserRegisterTypeEnum_UserRegisterType)
	*p = x
	return p
}

func (x UserRegisterTypeEnum_UserRegisterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRegisterTypeEnum_UserRegisterType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[3].Descriptor()
}

func (UserRegisterTypeEnum_UserRegisterType) Type() protoreflect.EnumType {
	return &file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes[3]
}

func (x UserRegisterTypeEnum_UserRegisterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRegisterTypeEnum_UserRegisterType.Descriptor instead.
func (UserRegisterTypeEnum_UserRegisterType) EnumDescriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{3, 0}
}

// UserInitEnum UserInitEnum enum
type UserInitEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserInitEnum) Reset() {
	*x = UserInitEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInitEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInitEnum) ProtoMessage() {}

func (x *UserInitEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInitEnum.ProtoReflect.Descriptor instead.
func (*UserInitEnum) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{0}
}

// UserStatusEnum 用户状态
type UserStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserStatusEnum) Reset() {
	*x = UserStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatusEnum) ProtoMessage() {}

func (x *UserStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatusEnum.ProtoReflect.Descriptor instead.
func (*UserStatusEnum) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{1}
}

// UserGenderEnum 用户性别
type UserGenderEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserGenderEnum) Reset() {
	*x = UserGenderEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGenderEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGenderEnum) ProtoMessage() {}

func (x *UserGenderEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGenderEnum.ProtoReflect.Descriptor instead.
func (*UserGenderEnum) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{2}
}

// UserRegisterTypeEnum ...
type UserRegisterTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserRegisterTypeEnum) Reset() {
	*x = UserRegisterTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterTypeEnum) ProtoMessage() {}

func (x *UserRegisterTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterTypeEnum.ProtoReflect.Descriptor instead.
func (*UserRegisterTypeEnum) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP(), []int{3}
}

var File_api_user_service_v1_enums_user_enum_v1_proto protoreflect.FileDescriptor

var file_api_user_service_v1_enums_user_enum_v1_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x76, 0x31, 0x22, 0x2b, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0x1b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x22, 0x6f, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0x5d, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49,
	0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x48, 0x49, 0x54, 0x45,
	0x4c, 0x49, 0x53, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x4c,
	0x49, 0x53, 0x54, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x05, 0x22, 0x51, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0x3f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43,
	0x52, 0x45, 0x54, 0x10, 0x03, 0x22, 0xaf, 0x01, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x96,
	0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x4f,
	0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x10, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x5f, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x10, 0x66, 0x12, 0x11,
	0x0a, 0x0d, 0x4a, 0x55, 0x4c, 0x49, 0x41, 0x4e, 0x47, 0x5f, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x10,
	0x67, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x49, 0x4e, 0x47, 0x54, 0x41, 0x4c, 0x4b, 0x5f, 0x4f, 0x41,
	0x55, 0x54, 0x48, 0x10, 0x68, 0x12, 0x10, 0x0a, 0x0b, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x59, 0x10, 0xd1, 0x0f, 0x42, 0x73, 0x0a, 0x14, 0x73, 0x61, 0x61, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x42,
	0x11, 0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_service_v1_enums_user_enum_v1_proto_rawDescOnce sync.Once
	file_api_user_service_v1_enums_user_enum_v1_proto_rawDescData = file_api_user_service_v1_enums_user_enum_v1_proto_rawDesc
)

func file_api_user_service_v1_enums_user_enum_v1_proto_rawDescGZIP() []byte {
	file_api_user_service_v1_enums_user_enum_v1_proto_rawDescOnce.Do(func() {
		file_api_user_service_v1_enums_user_enum_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_service_v1_enums_user_enum_v1_proto_rawDescData)
	})
	return file_api_user_service_v1_enums_user_enum_v1_proto_rawDescData
}

var file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_user_service_v1_enums_user_enum_v1_proto_goTypes = []interface{}{
	(UserInitEnum_UserInit)(0),                 // 0: saas.api.user.enumv1.UserInitEnum.UserInit
	(UserStatusEnum_UserStatus)(0),             // 1: saas.api.user.enumv1.UserStatusEnum.UserStatus
	(UserGenderEnum_UserGender)(0),             // 2: saas.api.user.enumv1.UserGenderEnum.UserGender
	(UserRegisterTypeEnum_UserRegisterType)(0), // 3: saas.api.user.enumv1.UserRegisterTypeEnum.UserRegisterType
	(*UserInitEnum)(nil),                       // 4: saas.api.user.enumv1.UserInitEnum
	(*UserStatusEnum)(nil),                     // 5: saas.api.user.enumv1.UserStatusEnum
	(*UserGenderEnum)(nil),                     // 6: saas.api.user.enumv1.UserGenderEnum
	(*UserRegisterTypeEnum)(nil),               // 7: saas.api.user.enumv1.UserRegisterTypeEnum
}
var file_api_user_service_v1_enums_user_enum_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_service_v1_enums_user_enum_v1_proto_init() }
func file_api_user_service_v1_enums_user_enum_v1_proto_init() {
	if File_api_user_service_v1_enums_user_enum_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInitEnum); i {
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
		file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatusEnum); i {
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
		file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGenderEnum); i {
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
		file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterTypeEnum); i {
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
			RawDescriptor: file_api_user_service_v1_enums_user_enum_v1_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_service_v1_enums_user_enum_v1_proto_goTypes,
		DependencyIndexes: file_api_user_service_v1_enums_user_enum_v1_proto_depIdxs,
		EnumInfos:         file_api_user_service_v1_enums_user_enum_v1_proto_enumTypes,
		MessageInfos:      file_api_user_service_v1_enums_user_enum_v1_proto_msgTypes,
	}.Build()
	File_api_user_service_v1_enums_user_enum_v1_proto = out.File
	file_api_user_service_v1_enums_user_enum_v1_proto_rawDesc = nil
	file_api_user_service_v1_enums_user_enum_v1_proto_goTypes = nil
	file_api_user_service_v1_enums_user_enum_v1_proto_depIdxs = nil
}
