// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: zitadel/user/v2beta/password.proto

package user

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type NotificationType int32

const (
	NotificationType_NOTIFICATION_TYPE_Unspecified NotificationType = 0
	NotificationType_NOTIFICATION_TYPE_Email       NotificationType = 1
	NotificationType_NOTIFICATION_TYPE_SMS         NotificationType = 2
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "NOTIFICATION_TYPE_Unspecified",
		1: "NOTIFICATION_TYPE_Email",
		2: "NOTIFICATION_TYPE_SMS",
	}
	NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_Unspecified": 0,
		"NOTIFICATION_TYPE_Email":       1,
		"NOTIFICATION_TYPE_SMS":         2,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_zitadel_user_v2beta_password_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_zitadel_user_v2beta_password_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_password_proto_rawDescGZIP(), []int{0}
}

type Password struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password       string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	ChangeRequired bool   `protobuf:"varint,2,opt,name=change_required,json=changeRequired,proto3" json:"change_required,omitempty"`
}

func (x *Password) Reset() {
	*x = Password{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Password) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Password) ProtoMessage() {}

func (x *Password) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Password.ProtoReflect.Descriptor instead.
func (*Password) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_password_proto_rawDescGZIP(), []int{0}
}

func (x *Password) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Password) GetChangeRequired() bool {
	if x != nil {
		return x.ChangeRequired
	}
	return false
}

type HashedPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash           string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ChangeRequired bool   `protobuf:"varint,2,opt,name=change_required,json=changeRequired,proto3" json:"change_required,omitempty"`
}

func (x *HashedPassword) Reset() {
	*x = HashedPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashedPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashedPassword) ProtoMessage() {}

func (x *HashedPassword) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashedPassword.ProtoReflect.Descriptor instead.
func (*HashedPassword) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_password_proto_rawDescGZIP(), []int{1}
}

func (x *HashedPassword) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *HashedPassword) GetChangeRequired() bool {
	if x != nil {
		return x.ChangeRequired
	}
	return false
}

type SendPasswordResetLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationType NotificationType `protobuf:"varint,1,opt,name=notification_type,json=notificationType,proto3,enum=zitadel.user.v2beta.NotificationType" json:"notification_type,omitempty"`
	UrlTemplate      *string          `protobuf:"bytes,2,opt,name=url_template,json=urlTemplate,proto3,oneof" json:"url_template,omitempty"`
}

func (x *SendPasswordResetLink) Reset() {
	*x = SendPasswordResetLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_password_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPasswordResetLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPasswordResetLink) ProtoMessage() {}

func (x *SendPasswordResetLink) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_password_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPasswordResetLink.ProtoReflect.Descriptor instead.
func (*SendPasswordResetLink) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_password_proto_rawDescGZIP(), []int{2}
}

func (x *SendPasswordResetLink) GetNotificationType() NotificationType {
	if x != nil {
		return x.NotificationType
	}
	return NotificationType_NOTIFICATION_TYPE_Unspecified
}

func (x *SendPasswordResetLink) GetUrlTemplate() string {
	if x != nil && x.UrlTemplate != nil {
		return *x.UrlTemplate
	}
	return ""
}

type ReturnPasswordResetCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReturnPasswordResetCode) Reset() {
	*x = ReturnPasswordResetCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_password_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnPasswordResetCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnPasswordResetCode) ProtoMessage() {}

func (x *ReturnPasswordResetCode) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_password_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnPasswordResetCode.ProtoReflect.Descriptor instead.
func (*ReturnPasswordResetCode) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_password_proto_rawDescGZIP(), []int{3}
}

var File_zitadel_user_v2beta_password_proto protoreflect.FileDescriptor

var file_zitadel_user_v2beta_password_proto_rawDesc = []byte{
	0x0a, 0x22, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x45, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x29, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0xe0, 0x41, 0x02,
	0x92, 0x41, 0x19, 0x4a, 0x11, 0x22, 0x53, 0x65, 0x63, 0x72, 0x33, 0x74, 0x50, 0x34, 0x73, 0x73,
	0x77, 0x30, 0x72, 0x64, 0x21, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22,
	0xb8, 0x01, 0x0a, 0x0e, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x7d, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x69, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0xe0, 0x41, 0x02, 0x92,
	0x41, 0x59, 0x32, 0x11, 0x22, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x20, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x4a, 0x3e, 0x22, 0x24, 0x32, 0x61, 0x24, 0x31, 0x32, 0x24, 0x6c,
	0x4a, 0x30, 0x38, 0x66, 0x71, 0x56, 0x72, 0x38, 0x62, 0x46, 0x4a, 0x69, 0x6c, 0x52, 0x56, 0x6e,
	0x44, 0x54, 0x39, 0x51, 0x65, 0x55, 0x4c, 0x49, 0x37, 0x59, 0x57, 0x2e, 0x6e, 0x54, 0x33, 0x69,
	0x77, 0x55, 0x76, 0x36, 0x64, 0x79, 0x67, 0x30, 0x61, 0x43, 0x72, 0x66, 0x6d, 0x33, 0x55, 0x59,
	0x38, 0x58, 0x52, 0x32, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0xe5, 0x03, 0x0a, 0x15, 0x53,
	0x65, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x52, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0xe6, 0x02, 0x0a, 0x0c, 0x75, 0x72, 0x6c,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0xbd, 0x02, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x92, 0x41, 0xaf, 0x02,
	0x32, 0xcb, 0x01, 0x22, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x20, 0x73,
	0x65, 0x74, 0x20, 0x61, 0x20, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x72, 0x65, 0x73, 0x65, 0x74, 0x20, 0x6d, 0x61, 0x69, 0x6c,
	0x20, 0x73, 0x65, 0x6e, 0x74, 0x20, 0x62, 0x79, 0x20, 0x5a, 0x49, 0x54, 0x41, 0x44, 0x45, 0x4c,
	0x20, 0x74, 0x6f, 0x20, 0x67, 0x75, 0x69, 0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x20, 0x49, 0x66, 0x20, 0x6e, 0x6f, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20,
	0x69, 0x73, 0x20, 0x73, 0x65, 0x74, 0x2c, 0x20, 0x74, 0x68, 0x65, 0x20, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x20, 0x5a, 0x49, 0x54, 0x41, 0x44, 0x45, 0x4c, 0x20, 0x75, 0x72, 0x6c, 0x20,
	0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x2e, 0x22, 0x4a, 0x59,
	0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x79, 0x3f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x3d, 0x7b, 0x7b, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x7d, 0x26, 0x63, 0x6f, 0x64, 0x65, 0x3d, 0x7b, 0x7b,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x7d, 0x7d, 0x26, 0x6f, 0x72, 0x67, 0x49, 0x44, 0x3d, 0x7b, 0x7b,
	0x2e, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x7d, 0x7d, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01, 0x48,
	0x00, 0x52, 0x0b, 0x75, 0x72, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x6d, 0x0a,
	0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x21, 0x0a, 0x1d, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4d, 0x53, 0x10, 0x02, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_user_v2beta_password_proto_rawDescOnce sync.Once
	file_zitadel_user_v2beta_password_proto_rawDescData = file_zitadel_user_v2beta_password_proto_rawDesc
)

func file_zitadel_user_v2beta_password_proto_rawDescGZIP() []byte {
	file_zitadel_user_v2beta_password_proto_rawDescOnce.Do(func() {
		file_zitadel_user_v2beta_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_user_v2beta_password_proto_rawDescData)
	})
	return file_zitadel_user_v2beta_password_proto_rawDescData
}

var file_zitadel_user_v2beta_password_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zitadel_user_v2beta_password_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_zitadel_user_v2beta_password_proto_goTypes = []interface{}{
	(NotificationType)(0),           // 0: zitadel.user.v2beta.NotificationType
	(*Password)(nil),                // 1: zitadel.user.v2beta.Password
	(*HashedPassword)(nil),          // 2: zitadel.user.v2beta.HashedPassword
	(*SendPasswordResetLink)(nil),   // 3: zitadel.user.v2beta.SendPasswordResetLink
	(*ReturnPasswordResetCode)(nil), // 4: zitadel.user.v2beta.ReturnPasswordResetCode
}
var file_zitadel_user_v2beta_password_proto_depIdxs = []int32{
	0, // 0: zitadel.user.v2beta.SendPasswordResetLink.notification_type:type_name -> zitadel.user.v2beta.NotificationType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_zitadel_user_v2beta_password_proto_init() }
func file_zitadel_user_v2beta_password_proto_init() {
	if File_zitadel_user_v2beta_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_user_v2beta_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Password); i {
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
		file_zitadel_user_v2beta_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashedPassword); i {
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
		file_zitadel_user_v2beta_password_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPasswordResetLink); i {
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
		file_zitadel_user_v2beta_password_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnPasswordResetCode); i {
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
	file_zitadel_user_v2beta_password_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_user_v2beta_password_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_user_v2beta_password_proto_goTypes,
		DependencyIndexes: file_zitadel_user_v2beta_password_proto_depIdxs,
		EnumInfos:         file_zitadel_user_v2beta_password_proto_enumTypes,
		MessageInfos:      file_zitadel_user_v2beta_password_proto_msgTypes,
	}.Build()
	File_zitadel_user_v2beta_password_proto = out.File
	file_zitadel_user_v2beta_password_proto_rawDesc = nil
	file_zitadel_user_v2beta_password_proto_goTypes = nil
	file_zitadel_user_v2beta_password_proto_depIdxs = nil
}
