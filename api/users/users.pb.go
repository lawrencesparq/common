// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: api/users/users.proto

package users

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

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName     string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber   string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Province      string `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	City          string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Address       string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Password      string `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	IsTnCsChecked bool   `protobuf:"varint,9,opt,name=IsTnCsChecked,proto3" json:"IsTnCsChecked,omitempty"`
	ID            string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_api_users_users_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegistrationRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegistrationRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegistrationRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *RegistrationRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *RegistrationRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegistrationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegistrationRequest) GetIsTnCsChecked() bool {
	if x != nil {
		return x.IsTnCsChecked
	}
	return false
}

func (x *RegistrationRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type RegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegistrationResponse) Reset() {
	*x = RegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationResponse) ProtoMessage() {}

func (x *RegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationResponse.ProtoReflect.Descriptor instead.
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return file_api_users_users_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RegistrationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LogInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ID       string `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	Otp      string `protobuf:"bytes,4,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *LogInRequest) Reset() {
	*x = LogInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInRequest) ProtoMessage() {}

func (x *LogInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInRequest.ProtoReflect.Descriptor instead.
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return file_api_users_users_proto_rawDescGZIP(), []int{2}
}

func (x *LogInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LogInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LogInRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LogInRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type UserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName          string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName           string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email              string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Province           string `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	City               string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Address            string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Password           string `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	ChangePassword     bool   `protobuf:"varint,9,opt,name=changePassword,proto3" json:"changePassword,omitempty"`
	IsActive           bool   `protobuf:"varint,10,opt,name=isActive,proto3" json:"isActive,omitempty"`
	IsTnCsChecked      bool   `protobuf:"varint,11,opt,name=IsTnCsChecked,proto3" json:"IsTnCsChecked,omitempty"`
	LoginAttempts      int64  `protobuf:"varint,12,opt,name=loginAttempts,proto3" json:"loginAttempts,omitempty"`
	IsLocked           bool   `protobuf:"varint,13,opt,name=isLocked,proto3" json:"isLocked,omitempty"`
	OtpExpiryDate      string `protobuf:"bytes,14,opt,name=otp_expiry_date,json=otpExpiryDate,proto3" json:"otp_expiry_date,omitempty"`
	LastLoginDate      string `protobuf:"bytes,15,opt,name=last_login_date,json=lastLoginDate,proto3" json:"last_login_date,omitempty"`
	LastPasswordChange string `protobuf:"bytes,16,opt,name=last_password_change,json=lastPasswordChange,proto3" json:"last_password_change,omitempty"`
	SendNewsLetter     bool   `protobuf:"varint,17,opt,name=sendNewsLetter,proto3" json:"sendNewsLetter,omitempty"`
	PhoneNumber        string `protobuf:"bytes,18,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	ProvinceID         int64  `protobuf:"varint,19,opt,name=provinceID,proto3" json:"provinceID,omitempty"`
	CityID             int64  `protobuf:"varint,20,opt,name=cityID,proto3" json:"cityID,omitempty"`
	Otp                string `protobuf:"bytes,21,opt,name=otp,proto3" json:"otp,omitempty"`
	CreatedAt          string `protobuf:"bytes,22,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          string `protobuf:"bytes,23,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserDetails) Reset() {
	*x = UserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetails) ProtoMessage() {}

func (x *UserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetails.ProtoReflect.Descriptor instead.
func (*UserDetails) Descriptor() ([]byte, []int) {
	return file_api_users_users_proto_rawDescGZIP(), []int{3}
}

func (x *UserDetails) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserDetails) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserDetails) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDetails) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *UserDetails) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserDetails) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserDetails) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserDetails) GetChangePassword() bool {
	if x != nil {
		return x.ChangePassword
	}
	return false
}

func (x *UserDetails) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *UserDetails) GetIsTnCsChecked() bool {
	if x != nil {
		return x.IsTnCsChecked
	}
	return false
}

func (x *UserDetails) GetLoginAttempts() int64 {
	if x != nil {
		return x.LoginAttempts
	}
	return 0
}

func (x *UserDetails) GetIsLocked() bool {
	if x != nil {
		return x.IsLocked
	}
	return false
}

func (x *UserDetails) GetOtpExpiryDate() string {
	if x != nil {
		return x.OtpExpiryDate
	}
	return ""
}

func (x *UserDetails) GetLastLoginDate() string {
	if x != nil {
		return x.LastLoginDate
	}
	return ""
}

func (x *UserDetails) GetLastPasswordChange() string {
	if x != nil {
		return x.LastPasswordChange
	}
	return ""
}

func (x *UserDetails) GetSendNewsLetter() bool {
	if x != nil {
		return x.SendNewsLetter
	}
	return false
}

func (x *UserDetails) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserDetails) GetProvinceID() int64 {
	if x != nil {
		return x.ProvinceID
	}
	return 0
}

func (x *UserDetails) GetCityID() int64 {
	if x != nil {
		return x.CityID
	}
	return 0
}

func (x *UserDetails) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *UserDetails) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserDetails) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UserDetails `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_api_users_users_proto_rawDescGZIP(), []int{4}
}

func (x *UserData) GetData() []*UserDetails {
	if x != nil {
		return x.Data
	}
	return nil
}

type LogInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogInResponse) Reset() {
	*x = LogInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInResponse) ProtoMessage() {}

func (x *LogInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInResponse.ProtoReflect.Descriptor instead.
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return file_api_users_users_proto_rawDescGZIP(), []int{5}
}

func (x *LogInResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LogInResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogInResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_api_users_users_proto protoreflect.FileDescriptor

var file_api_users_users_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0xa6, 0x02, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x49, 0x73, 0x54, 0x6e, 0x43, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x73, 0x54, 0x6e, 0x43, 0x73, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x62, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6f, 0x74, 0x70, 0x22, 0xd6, 0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x49, 0x73, 0x54, 0x6e, 0x43, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x73, 0x54, 0x6e, 0x43, 0x73, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x74, 0x70, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6f, 0x74, 0x70, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x6e,
	0x64, 0x4e, 0x65, 0x77, 0x73, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x4c, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x74, 0x70, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x30, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xce, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x45,
	0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x77, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x70, 0x61, 0x72, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_users_users_proto_rawDescOnce sync.Once
	file_api_users_users_proto_rawDescData = file_api_users_users_proto_rawDesc
)

func file_api_users_users_proto_rawDescGZIP() []byte {
	file_api_users_users_proto_rawDescOnce.Do(func() {
		file_api_users_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_users_users_proto_rawDescData)
	})
	return file_api_users_users_proto_rawDescData
}

var file_api_users_users_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_users_users_proto_goTypes = []any{
	(*RegistrationRequest)(nil),  // 0: api.RegistrationRequest
	(*RegistrationResponse)(nil), // 1: api.RegistrationResponse
	(*LogInRequest)(nil),         // 2: api.LogInRequest
	(*UserDetails)(nil),          // 3: api.UserDetails
	(*UserData)(nil),             // 4: api.UserData
	(*LogInResponse)(nil),        // 5: api.LogInResponse
}
var file_api_users_users_proto_depIdxs = []int32{
	3,  // 0: api.UserData.data:type_name -> api.UserDetails
	0,  // 1: api.UserService.RegisterUser:input_type -> api.RegistrationRequest
	2,  // 2: api.UserService.LogInUser:input_type -> api.LogInRequest
	2,  // 3: api.UserService.UserProfile:input_type -> api.LogInRequest
	2,  // 4: api.UserService.VerifyEmail:input_type -> api.LogInRequest
	2,  // 5: api.UserService.VerifyOTP:input_type -> api.LogInRequest
	2,  // 6: api.UserService.ResetUserPassword:input_type -> api.LogInRequest
	2,  // 7: api.UserService.CloseAccount:input_type -> api.LogInRequest
	2,  // 8: api.UserService.ActivateAccount:input_type -> api.LogInRequest
	0,  // 9: api.UserService.EditUserDetails:input_type -> api.RegistrationRequest
	2,  // 10: api.UserService.BackOfficeUsers:input_type -> api.LogInRequest
	1,  // 11: api.UserService.RegisterUser:output_type -> api.RegistrationResponse
	5,  // 12: api.UserService.LogInUser:output_type -> api.LogInResponse
	3,  // 13: api.UserService.UserProfile:output_type -> api.UserDetails
	5,  // 14: api.UserService.VerifyEmail:output_type -> api.LogInResponse
	5,  // 15: api.UserService.VerifyOTP:output_type -> api.LogInResponse
	5,  // 16: api.UserService.ResetUserPassword:output_type -> api.LogInResponse
	5,  // 17: api.UserService.CloseAccount:output_type -> api.LogInResponse
	5,  // 18: api.UserService.ActivateAccount:output_type -> api.LogInResponse
	1,  // 19: api.UserService.EditUserDetails:output_type -> api.RegistrationResponse
	4,  // 20: api.UserService.BackOfficeUsers:output_type -> api.UserData
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_users_users_proto_init() }
func file_api_users_users_proto_init() {
	if File_api_users_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_users_users_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegistrationRequest); i {
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
		file_api_users_users_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegistrationResponse); i {
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
		file_api_users_users_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LogInRequest); i {
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
		file_api_users_users_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserDetails); i {
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
		file_api_users_users_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserData); i {
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
		file_api_users_users_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LogInResponse); i {
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
			RawDescriptor: file_api_users_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_users_users_proto_goTypes,
		DependencyIndexes: file_api_users_users_proto_depIdxs,
		MessageInfos:      file_api_users_users_proto_msgTypes,
	}.Build()
	File_api_users_users_proto = out.File
	file_api_users_users_proto_rawDesc = nil
	file_api_users_users_proto_goTypes = nil
	file_api_users_users_proto_depIdxs = nil
}
