// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: controller/api/resources/accounts/v1/account.proto

package accounts

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	scopes "github.com/hashicorp/boundary/internal/gen/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/internal/gen/controller/protooptions"
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

// Account contains all fields related to an Account resource
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The ID of the Account.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Scope information for the Account.
	Scope *scopes.ScopeInfo `protobuf:"bytes,20,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional name for identification purposes.
	Name *wrappers.StringValue `protobuf:"bytes,30,opt,name=name,proto3" json:"name,omitempty"`
	// Optional user-set description for identification purposes.
	Description *wrappers.StringValue `protobuf:"bytes,40,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The time this resource was created.
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,50,opt,name=created_time,proto3" json:"created_time,omitempty"`
	// Output only. The time this resource was last updated.
	UpdatedTime *timestamp.Timestamp `protobuf:"bytes,60,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
	// Version is used in mutation requests, after the initial creation, to ensure this resource has not changed.
	// The mutation will fail if the version does not match the latest known good version.
	Version uint32 `protobuf:"varint,70,opt,name=version,proto3" json:"version,omitempty"`
	// The type of this Account.
	Type string `protobuf:"bytes,80,opt,name=type,proto3" json:"type,omitempty"`
	// The ID of the Auth Method that is associated with this Account.
	AuthMethodId string `protobuf:"bytes,90,opt,name=auth_method_id,proto3" json:"auth_method_id,omitempty"`
	// The attributes that are applicable for the specific Account type.
	Attributes *_struct.Struct `protobuf:"bytes,100,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_accounts_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_accounts_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_accounts_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Account) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Account) GetDescription() *wrappers.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Account) GetCreatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Account) GetUpdatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Account) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Account) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Account) GetAuthMethodId() string {
	if x != nil {
		return x.AuthMethodId
	}
	return ""
}

func (x *Account) GetAttributes() *_struct.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Account) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

// Attributes associated only with Accounts with type "password".
type PasswordAccountAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The login name of this Account. This is unique per Auth Method.
	LoginName string `protobuf:"bytes,10,opt,name=login_name,proto3" json:"login_name,omitempty"`
	// The password for this Account.
	Password *wrappers.StringValue `protobuf:"bytes,20,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PasswordAccountAttributes) Reset() {
	*x = PasswordAccountAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_accounts_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordAccountAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordAccountAttributes) ProtoMessage() {}

func (x *PasswordAccountAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_accounts_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordAccountAttributes.ProtoReflect.Descriptor instead.
func (*PasswordAccountAttributes) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_accounts_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *PasswordAccountAttributes) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *PasswordAccountAttributes) GetPassword() *wrappers.StringValue {
	if x != nil {
		return x.Password
	}
	return nil
}

// Attributes associated only with Accounts with type "oidc".
type OidcAccountAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. issuer is a case sensitive URL that maps to the OIDC iss claim
	Issuer string `protobuf:"bytes,80,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// Output only. subject is a case sensitive string that maps to the OIDC sub claim.
	Subject string `protobuf:"bytes,90,opt,name=subject,proto3" json:"subject,omitempty"`
	// Output only. full_name is a string that maps to the OIDC name claim
	FullName string `protobuf:"bytes,100,opt,name=full_name,proto3" json:"full_name,omitempty"`
	// Output only. email is a string that maps to the OIDC email claim.
	Email string `protobuf:"bytes,110,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *OidcAccountAttributes) Reset() {
	*x = OidcAccountAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_accounts_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OidcAccountAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OidcAccountAttributes) ProtoMessage() {}

func (x *OidcAccountAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_accounts_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OidcAccountAttributes.ProtoReflect.Descriptor instead.
func (*OidcAccountAttributes) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_accounts_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *OidcAccountAttributes) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *OidcAccountAttributes) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *OidcAccountAttributes) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *OidcAccountAttributes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_controller_api_resources_accounts_v1_account_proto protoreflect.FileDescriptor

var file_controller_api_resources_accounts_v1_account_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x04, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x14, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x62,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x22, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x46, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x04, 0xa0, 0xda, 0x29, 0x01, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xac, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x19, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xa0, 0xda, 0x29, 0x01,
	0xc2, 0xdd, 0x29, 0x22, 0x0a, 0x15, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x04, 0xa0, 0xda, 0x29, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x4f, 0x69, 0x64, 0x63, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x5a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xa0, 0xda, 0x29, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x6e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resources_accounts_v1_account_proto_rawDescOnce sync.Once
	file_controller_api_resources_accounts_v1_account_proto_rawDescData = file_controller_api_resources_accounts_v1_account_proto_rawDesc
)

func file_controller_api_resources_accounts_v1_account_proto_rawDescGZIP() []byte {
	file_controller_api_resources_accounts_v1_account_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_accounts_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_accounts_v1_account_proto_rawDescData)
	})
	return file_controller_api_resources_accounts_v1_account_proto_rawDescData
}

var file_controller_api_resources_accounts_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controller_api_resources_accounts_v1_account_proto_goTypes = []interface{}{
	(*Account)(nil),                   // 0: controller.api.resources.accounts.v1.Account
	(*PasswordAccountAttributes)(nil), // 1: controller.api.resources.accounts.v1.PasswordAccountAttributes
	(*OidcAccountAttributes)(nil),     // 2: controller.api.resources.accounts.v1.OidcAccountAttributes
	(*scopes.ScopeInfo)(nil),          // 3: controller.api.resources.scopes.v1.ScopeInfo
	(*wrappers.StringValue)(nil),      // 4: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),       // 5: google.protobuf.Timestamp
	(*_struct.Struct)(nil),            // 6: google.protobuf.Struct
}
var file_controller_api_resources_accounts_v1_account_proto_depIdxs = []int32{
	3, // 0: controller.api.resources.accounts.v1.Account.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	4, // 1: controller.api.resources.accounts.v1.Account.name:type_name -> google.protobuf.StringValue
	4, // 2: controller.api.resources.accounts.v1.Account.description:type_name -> google.protobuf.StringValue
	5, // 3: controller.api.resources.accounts.v1.Account.created_time:type_name -> google.protobuf.Timestamp
	5, // 4: controller.api.resources.accounts.v1.Account.updated_time:type_name -> google.protobuf.Timestamp
	6, // 5: controller.api.resources.accounts.v1.Account.attributes:type_name -> google.protobuf.Struct
	4, // 6: controller.api.resources.accounts.v1.PasswordAccountAttributes.password:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_controller_api_resources_accounts_v1_account_proto_init() }
func file_controller_api_resources_accounts_v1_account_proto_init() {
	if File_controller_api_resources_accounts_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_accounts_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_controller_api_resources_accounts_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordAccountAttributes); i {
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
		file_controller_api_resources_accounts_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OidcAccountAttributes); i {
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
			RawDescriptor: file_controller_api_resources_accounts_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_accounts_v1_account_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_accounts_v1_account_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_accounts_v1_account_proto_msgTypes,
	}.Build()
	File_controller_api_resources_accounts_v1_account_proto = out.File
	file_controller_api_resources_accounts_v1_account_proto_rawDesc = nil
	file_controller_api_resources_accounts_v1_account_proto_goTypes = nil
	file_controller_api_resources_accounts_v1_account_proto_depIdxs = nil
}
