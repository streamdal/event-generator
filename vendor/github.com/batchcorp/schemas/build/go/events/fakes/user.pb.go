// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package fakes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string       `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string       `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	MaidenName           string       `protobuf:"bytes,4,opt,name=maiden_name,json=maidenName,proto3" json:"maiden_name,omitempty"`
	Age                  int32        `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Gender               string       `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Email                string       `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string       `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
	Username             string       `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	BirthDate            string       `protobuf:"bytes,10,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	ProfileImgUrl        string       `protobuf:"bytes,11,opt,name=profile_img_url,json=profileImgUrl,proto3" json:"profile_img_url,omitempty"`
	BloodGroup           string       `protobuf:"bytes,12,opt,name=blood_group,json=bloodGroup,proto3" json:"blood_group,omitempty"`
	EyeColor             string       `protobuf:"bytes,13,opt,name=eye_color,json=eyeColor,proto3" json:"eye_color,omitempty"`
	Hair                 *UserHair    `protobuf:"bytes,14,opt,name=hair,proto3" json:"hair,omitempty"`
	Domain               string       `protobuf:"bytes,15,opt,name=domain,proto3" json:"domain,omitempty"`
	IpAddress            string       `protobuf:"bytes,16,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Address              *UserAddress `protobuf:"bytes,17,opt,name=address,proto3" json:"address,omitempty"`
	MacAddress           string       `protobuf:"bytes,18,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	University           string       `protobuf:"bytes,19,opt,name=university,proto3" json:"university,omitempty"`
	Bank                 *UserBank    `protobuf:"bytes,20,opt,name=bank,proto3" json:"bank,omitempty"`
	Company              *UserCompany `protobuf:"bytes,21,opt,name=company,proto3" json:"company,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetMaidenName() string {
	if m != nil {
		return m.MaidenName
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

func (m *User) GetProfileImgUrl() string {
	if m != nil {
		return m.ProfileImgUrl
	}
	return ""
}

func (m *User) GetBloodGroup() string {
	if m != nil {
		return m.BloodGroup
	}
	return ""
}

func (m *User) GetEyeColor() string {
	if m != nil {
		return m.EyeColor
	}
	return ""
}

func (m *User) GetHair() *UserHair {
	if m != nil {
		return m.Hair
	}
	return nil
}

func (m *User) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *User) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *User) GetAddress() *UserAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *User) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *User) GetUniversity() string {
	if m != nil {
		return m.University
	}
	return ""
}

func (m *User) GetBank() *UserBank {
	if m != nil {
		return m.Bank
	}
	return nil
}

func (m *User) GetCompany() *UserCompany {
	if m != nil {
		return m.Company
	}
	return nil
}

type UserHair struct {
	Color                string   `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Style                string   `protobuf:"bytes,2,opt,name=style,proto3" json:"style,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserHair) Reset()         { *m = UserHair{} }
func (m *UserHair) String() string { return proto.CompactTextString(m) }
func (*UserHair) ProtoMessage()    {}
func (*UserHair) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserHair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserHair.Unmarshal(m, b)
}
func (m *UserHair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserHair.Marshal(b, m, deterministic)
}
func (m *UserHair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserHair.Merge(m, src)
}
func (m *UserHair) XXX_Size() int {
	return xxx_messageInfo_UserHair.Size(m)
}
func (m *UserHair) XXX_DiscardUnknown() {
	xxx_messageInfo_UserHair.DiscardUnknown(m)
}

var xxx_messageInfo_UserHair proto.InternalMessageInfo

func (m *UserHair) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *UserHair) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

type UserAddress struct {
	Street               string           `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City                 string           `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string           `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zip                  string           `protobuf:"bytes,4,opt,name=zip,proto3" json:"zip,omitempty"`
	CountryCode          string           `protobuf:"bytes,5,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Coordinates          *UserCoordinates `protobuf:"bytes,6,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserAddress) Reset()         { *m = UserAddress{} }
func (m *UserAddress) String() string { return proto.CompactTextString(m) }
func (*UserAddress) ProtoMessage()    {}
func (*UserAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddress.Unmarshal(m, b)
}
func (m *UserAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddress.Marshal(b, m, deterministic)
}
func (m *UserAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddress.Merge(m, src)
}
func (m *UserAddress) XXX_Size() int {
	return xxx_messageInfo_UserAddress.Size(m)
}
func (m *UserAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddress.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddress proto.InternalMessageInfo

func (m *UserAddress) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *UserAddress) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserAddress) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *UserAddress) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *UserAddress) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *UserAddress) GetCoordinates() *UserCoordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

type UserCoordinates struct {
	Latitude             float32  `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCoordinates) Reset()         { *m = UserCoordinates{} }
func (m *UserCoordinates) String() string { return proto.CompactTextString(m) }
func (*UserCoordinates) ProtoMessage()    {}
func (*UserCoordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserCoordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCoordinates.Unmarshal(m, b)
}
func (m *UserCoordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCoordinates.Marshal(b, m, deterministic)
}
func (m *UserCoordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCoordinates.Merge(m, src)
}
func (m *UserCoordinates) XXX_Size() int {
	return xxx_messageInfo_UserCoordinates.Size(m)
}
func (m *UserCoordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCoordinates.DiscardUnknown(m)
}

var xxx_messageInfo_UserCoordinates proto.InternalMessageInfo

func (m *UserCoordinates) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *UserCoordinates) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type UserBank struct {
	CardExpireMonth      int32    `protobuf:"varint,1,opt,name=card_expire_month,json=cardExpireMonth,proto3" json:"card_expire_month,omitempty"`
	CardExpireYear       int32    `protobuf:"varint,2,opt,name=card_expire_year,json=cardExpireYear,proto3" json:"card_expire_year,omitempty"`
	CardNumber           int64    `protobuf:"varint,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardType             string   `protobuf:"bytes,4,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	Currency             string   `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Iban                 string   `protobuf:"bytes,6,opt,name=iban,proto3" json:"iban,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBank) Reset()         { *m = UserBank{} }
func (m *UserBank) String() string { return proto.CompactTextString(m) }
func (*UserBank) ProtoMessage()    {}
func (*UserBank) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserBank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBank.Unmarshal(m, b)
}
func (m *UserBank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBank.Marshal(b, m, deterministic)
}
func (m *UserBank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBank.Merge(m, src)
}
func (m *UserBank) XXX_Size() int {
	return xxx_messageInfo_UserBank.Size(m)
}
func (m *UserBank) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBank.DiscardUnknown(m)
}

var xxx_messageInfo_UserBank proto.InternalMessageInfo

func (m *UserBank) GetCardExpireMonth() int32 {
	if m != nil {
		return m.CardExpireMonth
	}
	return 0
}

func (m *UserBank) GetCardExpireYear() int32 {
	if m != nil {
		return m.CardExpireYear
	}
	return 0
}

func (m *UserBank) GetCardNumber() int64 {
	if m != nil {
		return m.CardNumber
	}
	return 0
}

func (m *UserBank) GetCardType() string {
	if m != nil {
		return m.CardType
	}
	return ""
}

func (m *UserBank) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *UserBank) GetIban() string {
	if m != nil {
		return m.Iban
	}
	return ""
}

type UserCompany struct {
	Address                *UserAddress `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Department             string       `protobuf:"bytes,2,opt,name=department,proto3" json:"department,omitempty"`
	Name                   string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Title                  string       `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Ein                    string       `protobuf:"bytes,5,opt,name=ein,proto3" json:"ein,omitempty"`
	DateStartedUnixUtcNano int64        `protobuf:"varint,6,opt,name=date_started_unix_utc_nano,json=dateStartedUnixUtcNano,proto3" json:"date_started_unix_utc_nano,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}     `json:"-"`
	XXX_unrecognized       []byte       `json:"-"`
	XXX_sizecache          int32        `json:"-"`
}

func (m *UserCompany) Reset()         { *m = UserCompany{} }
func (m *UserCompany) String() string { return proto.CompactTextString(m) }
func (*UserCompany) ProtoMessage()    {}
func (*UserCompany) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserCompany) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCompany.Unmarshal(m, b)
}
func (m *UserCompany) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCompany.Marshal(b, m, deterministic)
}
func (m *UserCompany) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCompany.Merge(m, src)
}
func (m *UserCompany) XXX_Size() int {
	return xxx_messageInfo_UserCompany.Size(m)
}
func (m *UserCompany) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCompany.DiscardUnknown(m)
}

var xxx_messageInfo_UserCompany proto.InternalMessageInfo

func (m *UserCompany) GetAddress() *UserAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserCompany) GetDepartment() string {
	if m != nil {
		return m.Department
	}
	return ""
}

func (m *UserCompany) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserCompany) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UserCompany) GetEin() string {
	if m != nil {
		return m.Ein
	}
	return ""
}

func (m *UserCompany) GetDateStartedUnixUtcNano() int64 {
	if m != nil {
		return m.DateStartedUnixUtcNano
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "fakes.User")
	proto.RegisterType((*UserHair)(nil), "fakes.UserHair")
	proto.RegisterType((*UserAddress)(nil), "fakes.UserAddress")
	proto.RegisterType((*UserCoordinates)(nil), "fakes.UserCoordinates")
	proto.RegisterType((*UserBank)(nil), "fakes.UserBank")
	proto.RegisterType((*UserCompany)(nil), "fakes.UserCompany")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0x23, 0x45,
	0x10, 0x95, 0xed, 0x38, 0x6b, 0x97, 0x77, 0xe3, 0x6c, 0xb3, 0x44, 0xad, 0x05, 0x96, 0x60, 0x24,
	0x14, 0x21, 0x14, 0x4b, 0x01, 0x21, 0xc4, 0x8d, 0x0d, 0x08, 0x10, 0x22, 0x87, 0x81, 0x1c, 0xe0,
	0x32, 0xea, 0xe9, 0xae, 0x8c, 0x5b, 0x99, 0xe9, 0x1e, 0xf5, 0xf4, 0xac, 0x32, 0x7c, 0x01, 0x5f,
	0xc5, 0x27, 0x70, 0xe1, 0x87, 0x50, 0x55, 0xcf, 0x24, 0x86, 0x1c, 0xf6, 0xd6, 0xf5, 0xde, 0x73,
	0xb9, 0x5f, 0xd5, 0x9b, 0x06, 0xe8, 0x5a, 0x0c, 0xe7, 0x4d, 0xf0, 0xd1, 0x8b, 0xf9, 0x8d, 0xba,
	0xc5, 0x76, 0xf3, 0xe7, 0x1c, 0x0e, 0xae, 0x5b, 0x0c, 0xe2, 0x08, 0xa6, 0xd6, 0xc8, 0xc9, 0xe9,
	0xe4, 0x6c, 0x99, 0x4d, 0xad, 0x11, 0x1f, 0x00, 0xdc, 0xd8, 0xd0, 0xc6, 0xdc, 0xa9, 0x1a, 0xe5,
	0x94, 0xf1, 0x25, 0x23, 0x57, 0xaa, 0x46, 0xf1, 0x1e, 0x2c, 0x2b, 0x35, 0xb2, 0x33, 0x66, 0x17,
	0x04, 0x30, 0xf9, 0x21, 0xac, 0x6a, 0x65, 0x0d, 0xba, 0x44, 0x1f, 0x30, 0x0d, 0x09, 0x62, 0xc1,
	0x31, 0xcc, 0x54, 0x89, 0x72, 0x7e, 0x3a, 0x39, 0x9b, 0x67, 0x74, 0x14, 0x27, 0x70, 0x58, 0xa2,
	0x33, 0x18, 0xe4, 0x21, 0xab, 0x87, 0x4a, 0xbc, 0x80, 0x39, 0xd6, 0xca, 0x56, 0xf2, 0x09, 0xc3,
	0xa9, 0x20, 0xb4, 0xd9, 0x79, 0x87, 0x72, 0x91, 0x50, 0x2e, 0xc4, 0x4b, 0x58, 0x90, 0x41, 0xfe,
	0xcf, 0x65, 0xba, 0xd2, 0x58, 0x93, 0x9d, 0xc2, 0x86, 0xb8, 0xcb, 0x8d, 0x8a, 0x28, 0x21, 0xd9,
	0x61, 0xe4, 0x5b, 0x15, 0x51, 0x7c, 0x02, 0xeb, 0x26, 0xf8, 0x1b, 0x5b, 0x61, 0x6e, 0xeb, 0x32,
	0xef, 0x42, 0x25, 0x57, 0xac, 0x79, 0x36, 0xc0, 0x3f, 0xd6, 0xe5, 0x75, 0xa8, 0xc8, 0x59, 0x51,
	0x79, 0x6f, 0xf2, 0x32, 0xf8, 0xae, 0x91, 0x4f, 0x93, 0x33, 0x86, 0xbe, 0x27, 0x84, 0xe6, 0x82,
	0x3d, 0xe6, 0xda, 0x57, 0x3e, 0xc8, 0x67, 0xe9, 0x12, 0xd8, 0xe3, 0x25, 0xd5, 0xe2, 0x63, 0x38,
	0xd8, 0x29, 0x1b, 0xe4, 0xd1, 0xe9, 0xe4, 0x6c, 0x75, 0xb1, 0x3e, 0xe7, 0x15, 0x9c, 0xd3, 0xf8,
	0x7f, 0x50, 0x36, 0x64, 0x4c, 0xd2, 0x24, 0x8c, 0xaf, 0x95, 0x75, 0x72, 0x9d, 0x26, 0x91, 0x2a,
	0x72, 0x60, 0x9b, 0x5c, 0x19, 0x13, 0xb0, 0x6d, 0xe5, 0x71, 0x72, 0x60, 0x9b, 0x6f, 0x12, 0x20,
	0x3e, 0x83, 0x27, 0x23, 0xf7, 0x9c, 0xdb, 0x8b, 0xbd, 0xf6, 0x83, 0x28, 0x1b, 0x25, 0x69, 0x43,
	0xfa, 0xbe, 0x9b, 0x18, 0x37, 0xa4, 0xc7, 0x76, 0xaf, 0x00, 0x3a, 0x67, 0xdf, 0x60, 0x68, 0x6d,
	0xec, 0xe5, 0x3b, 0x89, 0x7f, 0x40, 0xc8, 0x4a, 0xa1, 0xdc, 0xad, 0x7c, 0xf1, 0xc8, 0xca, 0x6b,
	0xe5, 0x6e, 0x33, 0x26, 0xe9, 0x4e, 0xda, 0xd7, 0x8d, 0x72, 0xbd, 0x7c, 0xf7, 0xd1, 0x9d, 0x2e,
	0x13, 0x93, 0x8d, 0x92, 0xcd, 0x97, 0xb0, 0x18, 0x47, 0x41, 0x0b, 0x4e, 0x23, 0x4c, 0x81, 0x4c,
	0x05, 0xa1, 0x6d, 0xec, 0xab, 0x31, 0x8e, 0xa9, 0xd8, 0xfc, 0x35, 0x81, 0xd5, 0x9e, 0x49, 0x1a,
	0x60, 0x1b, 0x03, 0x62, 0x1c, 0x7e, 0x3c, 0x54, 0x42, 0xc0, 0x81, 0x26, 0x33, 0xe9, 0xc7, 0x7c,
	0x4e, 0x1d, 0x29, 0x11, 0xb3, 0xb1, 0x23, 0xa5, 0xe1, 0x18, 0x66, 0x7f, 0xd8, 0x66, 0xc8, 0x2d,
	0x1d, 0xc5, 0x47, 0xf0, 0x54, 0xfb, 0xce, 0xc5, 0xd0, 0xe7, 0xda, 0x9b, 0x94, 0xdc, 0x65, 0xb6,
	0x1a, 0xb0, 0x4b, 0x6f, 0x50, 0x7c, 0x05, 0x2b, 0xed, 0x7d, 0x30, 0xd6, 0xa9, 0x88, 0x2d, 0xc7,
	0x78, 0x75, 0x71, 0xf2, 0x1f, 0xc3, 0xf7, 0x6c, 0xb6, 0x2f, 0xdd, 0xfc, 0x04, 0xeb, 0xff, 0xf1,
	0x14, 0xe5, 0x4a, 0x45, 0x1b, 0x3b, 0x83, 0xec, 0x62, 0x9a, 0xdd, 0xd7, 0xe2, 0x7d, 0x58, 0x56,
	0xde, 0x95, 0x89, 0x9c, 0x32, 0xf9, 0x00, 0x6c, 0xfe, 0x9e, 0xa4, 0x31, 0xd2, 0x1a, 0xc4, 0xa7,
	0xf0, 0x5c, 0xab, 0x60, 0x72, 0xbc, 0x6b, 0x6c, 0xc0, 0xbc, 0xf6, 0x2e, 0xee, 0xb8, 0xdf, 0x3c,
	0x5b, 0x13, 0xf1, 0x1d, 0xe3, 0x3f, 0x13, 0x2c, 0xce, 0xe0, 0x78, 0x5f, 0xdb, 0xa3, 0x0a, 0xdc,
	0x7d, 0x9e, 0x1d, 0x3d, 0x48, 0x7f, 0x43, 0x15, 0x28, 0x3c, 0xac, 0x74, 0x5d, 0x5d, 0x60, 0xe0,
	0xd1, 0xcd, 0x32, 0x20, 0xe8, 0x8a, 0x11, 0xfa, 0x08, 0x58, 0x10, 0xfb, 0x66, 0xfc, 0xfa, 0x17,
	0x04, 0xfc, 0xda, 0x37, 0xfc, 0x95, 0xea, 0x2e, 0x04, 0x74, 0xba, 0x1f, 0xc6, 0x78, 0x5f, 0xd3,
	0x8a, 0x6c, 0xa1, 0xdc, 0xf0, 0x06, 0xf0, 0x79, 0xf3, 0xcf, 0xb0, 0xde, 0x21, 0x2f, 0xfb, 0x41,
	0x9f, 0xbc, 0x3d, 0xe8, 0xaf, 0x00, 0x0c, 0x36, 0x2a, 0xc4, 0x1a, 0x5d, 0x1c, 0x56, 0xbf, 0x87,
	0xd0, 0x3f, 0xee, 0x3d, 0x61, 0x7c, 0xa6, 0x50, 0x44, 0x1b, 0xab, 0xf1, 0xea, 0xa9, 0xa0, 0x50,
	0xa0, 0x75, 0xc3, 0x95, 0xe9, 0x28, 0xbe, 0x86, 0x97, 0xf4, 0x9a, 0xe4, 0x6d, 0x54, 0x21, 0xa2,
	0xc9, 0x3b, 0x67, 0xef, 0xf2, 0x2e, 0xea, 0xdc, 0x29, 0xe7, 0xd9, 0xc3, 0x2c, 0x3b, 0x21, 0xc5,
	0x2f, 0x49, 0x70, 0xed, 0xec, 0xdd, 0x75, 0xd4, 0x57, 0xca, 0xf9, 0xd7, 0x5f, 0xfc, 0x7e, 0x51,
	0xda, 0xb8, 0xeb, 0x8a, 0x73, 0xed, 0xeb, 0x6d, 0xa1, 0xa2, 0xde, 0x69, 0x1f, 0x9a, 0x6d, 0xab,
	0x77, 0x58, 0xab, 0x76, 0x5b, 0x74, 0xb6, 0x32, 0xdb, 0xd2, 0x6f, 0xf1, 0x0d, 0xba, 0xd8, 0x6e,
	0xd9, 0x62, 0x71, 0xc8, 0x6f, 0xf7, 0xe7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x20, 0xe6, 0x0c,
	0x0e, 0xc9, 0x05, 0x00, 0x00,
}
