// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: auth.proto

package pb_auth

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

type PLATFORM_TYPE int32

const (
	PLATFORM_TYPE_UNKNOWN_PLATFORM_TYPE PLATFORM_TYPE = 0
	PLATFORM_TYPE_IOS                   PLATFORM_TYPE = 1
	PLATFORM_TYPE_ANDROID               PLATFORM_TYPE = 2
	PLATFORM_TYPE_MAC                   PLATFORM_TYPE = 3
	PLATFORM_TYPE_WINDOWS               PLATFORM_TYPE = 4
	PLATFORM_TYPE_WEB                   PLATFORM_TYPE = 5
)

// Enum value maps for PLATFORM_TYPE.
var (
	PLATFORM_TYPE_name = map[int32]string{
		0: "UNKNOWN_PLATFORM_TYPE",
		1: "IOS",
		2: "ANDROID",
		3: "MAC",
		4: "WINDOWS",
		5: "WEB",
	}
	PLATFORM_TYPE_value = map[string]int32{
		"UNKNOWN_PLATFORM_TYPE": 0,
		"IOS":                   1,
		"ANDROID":               2,
		"MAC":                   3,
		"WINDOWS":               4,
		"WEB":                   5,
	}
)

func (x PLATFORM_TYPE) Enum() *PLATFORM_TYPE {
	p := new(PLATFORM_TYPE)
	*p = x
	return p
}

func (x PLATFORM_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PLATFORM_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (PLATFORM_TYPE) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x PLATFORM_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PLATFORM_TYPE.Descriptor instead.
func (PLATFORM_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegPlatform PLATFORM_TYPE `protobuf:"varint,1,opt,name=reg_platform,json=regPlatform,proto3,enum=pb_auth.PLATFORM_TYPE" json:"reg_platform,omitempty"` // 注册平台
	NickName    string        `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`                                                      // 昵称
	Password    string        `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                                                      // 密码
	FirstName   string        `protobuf:"bytes,4,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string        `protobuf:"bytes,5,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Gender      int32         `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthTs     int64         `protobuf:"varint,7,opt,name=birth_ts,json=birthTs,proto3" json:"birth_ts,omitempty"` // 生日
	Email       string        `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Mobile      string        `protobuf:"bytes,9,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Avatar      string        `protobuf:"bytes,10,opt,name=avatar,proto3" json:"avatar,omitempty"`                //头像url
	CityId      int64         `protobuf:"varint,11,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"` // 城市ID
	Code        int64         `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`                   // 验证码
	Uuid        string        `protobuf:"bytes,13,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ServerId    int64         `protobuf:"varint,14,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"` // 服务器编号
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpReq) GetRegPlatform() PLATFORM_TYPE {
	if x != nil {
		return x.RegPlatform
	}
	return PLATFORM_TYPE_UNKNOWN_PLATFORM_TYPE
}

func (x *SignUpReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SignUpReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SignUpReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SignUpReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *SignUpReq) GetBirthTs() int64 {
	if x != nil {
		return x.BirthTs
	}
	return 0
}

func (x *SignUpReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SignUpReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SignUpReq) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *SignUpReq) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignUpReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SignUpReq) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type SignUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg          string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserInfo     *UserInfo `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	AccessToken  *Token    `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken *Token    `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *SignUpResp) Reset() {
	*x = SignUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResp) ProtoMessage() {}

func (x *SignUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResp.ProtoReflect.Descriptor instead.
func (*SignUpResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignUpResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SignUpResp) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *SignUpResp) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *SignUpResp) GetRefreshToken() *Token {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName  string `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Gender    int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthTs   int64  `protobuf:"varint,5,opt,name=birth_ts,json=birthTs,proto3" json:"birth_ts,omitempty"`
	Email     string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Mobile    string `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfo) GetBirthTs() int64 {
	if x != nil {
		return x.BirthTs
	}
	return 0
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag:json:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// @inject_tag:json:"expire"
	Expire int64 `protobuf:"varint,2,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x22, 0x8f, 0x03, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x33, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc1, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x62, 0x69, 0x72, 0x74, 0x68, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x35, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x2a,
	0x5f, 0x0a, 0x0d, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49,
	0x4f, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x49,
	0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x05,
	0x32, 0x39, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2f, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_proto_goTypes = []interface{}{
	(PLATFORM_TYPE)(0), // 0: pb_auth.PLATFORM_TYPE
	(*SignUpReq)(nil),  // 1: pb_auth.SignUpReq
	(*SignUpResp)(nil), // 2: pb_auth.SignUpResp
	(*UserInfo)(nil),   // 3: pb_auth.UserInfo
	(*Token)(nil),      // 4: pb_auth.Token
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: pb_auth.SignUpReq.reg_platform:type_name -> pb_auth.PLATFORM_TYPE
	3, // 1: pb_auth.SignUpResp.user_info:type_name -> pb_auth.UserInfo
	4, // 2: pb_auth.SignUpResp.access_token:type_name -> pb_auth.Token
	4, // 3: pb_auth.SignUpResp.refresh_token:type_name -> pb_auth.Token
	1, // 4: pb_auth.Auth.SignUp:input_type -> pb_auth.SignUpReq
	2, // 5: pb_auth.Auth.SignUp:output_type -> pb_auth.SignUpResp
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpReq); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResp); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
