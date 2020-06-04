// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Otis/backend/dtypes/user/proto/user.proto

package user

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type User_Gender int32

const (
	User_MALE   User_Gender = 0
	User_FEMALE User_Gender = 1
	User_OTHER  User_Gender = 2
)

var User_Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
	2: "OTHER",
}

var User_Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
	"OTHER":  2,
}

func (x User_Gender) String() string {
	return proto.EnumName(User_Gender_name, int32(x))
}

func (User_Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_398f4ef115ea1f8f, []int{0, 0}
}

type User struct {
	Id                   int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Dob                  string      `protobuf:"bytes,2,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender               User_Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=user.User_Gender" json:"gender,omitempty"`
	GenderDescription    string      `protobuf:"bytes,4,opt,name=genderDescription,proto3" json:"genderDescription,omitempty"`
	UniversityId         string      `protobuf:"bytes,5,opt,name=university_id,json=universityId,proto3" json:"university_id,omitempty"`
	CreatedAt            string      `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	GraduationYear       string      `protobuf:"bytes,7,opt,name=graduationYear,proto3" json:"graduationYear,omitempty"`
	PhotoUrl             string      `protobuf:"bytes,12,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	FirstName            string      `protobuf:"bytes,13,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string      `protobuf:"bytes,14,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Alias                string      `protobuf:"bytes,15,opt,name=alias,proto3" json:"alias,omitempty"`
	Country              string      `protobuf:"bytes,16,opt,name=country,proto3" json:"country,omitempty"`
	AverageWeeklySpend   float64     `protobuf:"fixed64,17,opt,name=averageWeeklySpend,proto3" json:"averageWeeklySpend,omitempty"`
	ExpenseBands         string      `protobuf:"bytes,18,opt,name=expenseBands,proto3" json:"expenseBands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_398f4ef115ea1f8f, []int{0}
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

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetDob() string {
	if m != nil {
		return m.Dob
	}
	return ""
}

func (m *User) GetGender() User_Gender {
	if m != nil {
		return m.Gender
	}
	return User_MALE
}

func (m *User) GetGenderDescription() string {
	if m != nil {
		return m.GenderDescription
	}
	return ""
}

func (m *User) GetUniversityId() string {
	if m != nil {
		return m.UniversityId
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetGraduationYear() string {
	if m != nil {
		return m.GraduationYear
	}
	return ""
}

func (m *User) GetPhotoUrl() string {
	if m != nil {
		return m.PhotoUrl
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

func (m *User) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *User) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *User) GetAverageWeeklySpend() float64 {
	if m != nil {
		return m.AverageWeeklySpend
	}
	return 0
}

func (m *User) GetExpenseBands() string {
	if m != nil {
		return m.ExpenseBands
	}
	return ""
}

func init() {
	proto.RegisterEnum("user.User_Gender", User_Gender_name, User_Gender_value)
	proto.RegisterType((*User)(nil), "user.User")
}

func init() {
	proto.RegisterFile("Otis/backend/dtypes/user/proto/user.proto", fileDescriptor_398f4ef115ea1f8f)
}

var fileDescriptor_398f4ef115ea1f8f = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd4, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0x07, 0xf0, 0x4d, 0x48, 0x02, 0x19, 0x01, 0x0b, 0x73, 0x65, 0xed, 0xae, 0xd6, 0x91, 0xb5,
	0x12, 0x89, 0xc4, 0x26, 0x61, 0x09, 0x5a, 0x9a, 0xf4, 0x26, 0x21, 0x86, 0x46, 0x22, 0xa1, 0x72,
	0x9d, 0x7e, 0xa9, 0x52, 0x3a, 0xf1, 0x1c, 0xcc, 0x80, 0xe3, 0x71, 0xc7, 0x63, 0x44, 0x1e, 0xac,
	0xaf, 0xe3, 0x8b, 0x3e, 0x82, 0x9f, 0xa0, 0xf2, 0x84, 0x10, 0x28, 0x25, 0x42, 0xea, 0xdd, 0xb1,
	0x3c, 0xe7, 0xe7, 0xff, 0x8c, 0x8f, 0x8d, 0x2a, 0x67, 0x92, 0x85, 0xb5, 0x31, 0x71, 0xae, 0xc0,
	0xa7, 0x35, 0x2a, 0xa7, 0x01, 0x84, 0xb5, 0x28, 0x04, 0x51, 0x0b, 0x04, 0x97, 0x5c, 0x95, 0x55,
	0x55, 0xe2, 0x5c, 0x5a, 0xff, 0xd1, 0x70, 0x79, 0x2d, 0x14, 0x4e, 0xcd, 0x65, 0xf2, 0x22, 0x1a,
	0x57, 0x1d, 0x3e, 0xa9, 0xb9, 0xdc, 0xe5, 0xb3, 0xe5, 0xe3, 0xe8, 0x5c, 0x5d, 0xcd, 0x7a, 0xd3,
	0x6a, 0xd6, 0x6b, 0x7c, 0x2d, 0xa2, 0xdc, 0x30, 0x04, 0x81, 0x2f, 0x51, 0x96, 0x51, 0x2d, 0x53,
	0xca, 0x94, 0xf3, 0x9d, 0x8f, 0xdf, 0x62, 0x3d, 0xdb, 0xeb, 0x26, 0xb1, 0xfe, 0x3a, 0xfc, 0xe2,
	0x35, 0x8d, 0x40, 0xb0, 0x09, 0x11, 0xd3, 0xd1, 0x15, 0x4c, 0x5b, 0x0e, 0xf7, 0xa2, 0x89, 0xdf,
	0x64, 0xb4, 0x95, 0x46, 0x6a, 0xf6, 0x06, 0xb6, 0x79, 0x62, 0x5a, 0xad, 0xf6, 0xd0, 0x3e, 0x1b,
	0xf5, 0x06, 0x47, 0x96, 0xd9, 0x37, 0x07, 0xb6, 0x51, 0xba, 0x0c, 0xb9, 0xdf, 0x34, 0x7c, 0x32,
	0x81, 0x26, 0xa3, 0xbb, 0x7c, 0xc2, 0x24, 0x4c, 0x02, 0x39, 0x35, 0xac, 0x2c, 0xa3, 0xb8, 0x8a,
	0x56, 0x28, 0x1f, 0x6b, 0xd9, 0x52, 0xa6, 0x5c, 0xec, 0xfc, 0x95, 0xc4, 0xba, 0xa6, 0x1e, 0xa3,
	0x40, 0x4a, 0x24, 0xb4, 0x7c, 0x2e, 0x4b, 0x7e, 0xe4, 0x79, 0x86, 0x95, 0x2e, 0xc4, 0x7d, 0x54,
	0x70, 0xc1, 0xa7, 0x20, 0xb4, 0x95, 0x52, 0xa6, 0xbc, 0xf9, 0xdf, 0x76, 0x55, 0xed, 0x3e, 0xcd,
	0x5d, 0x3d, 0x51, 0x37, 0x3a, 0x46, 0x12, 0xeb, 0x7f, 0x2f, 0x94, 0x6b, 0x22, 0x9c, 0x0b, 0x22,
	0xca, 0x87, 0x95, 0x7b, 0xd6, 0x2d, 0x82, 0xc7, 0x68, 0x7b, 0x56, 0x75, 0x21, 0x74, 0x04, 0x0b,
	0x24, 0xe3, 0xbe, 0x96, 0x53, 0x61, 0x1a, 0x49, 0xac, 0xd7, 0x67, 0xd9, 0x17, 0x99, 0x4b, 0x8f,
	0xdd, 0xbd, 0x7a, 0xbd, 0xd2, 0xea, 0x9a, 0xc7, 0xed, 0xe1, 0xa9, 0xdd, 0xb4, 0x1e, 0x73, 0xf8,
	0x2d, 0xda, 0x88, 0x7c, 0x76, 0x0d, 0x22, 0x64, 0x72, 0x3a, 0x62, 0x54, 0xcb, 0x2b, 0x7f, 0x2f,
	0x89, 0xf5, 0x7f, 0x97, 0xf8, 0xcc, 0x97, 0xe0, 0x82, 0xb8, 0xa3, 0x0d, 0xc3, 0x5a, 0x5f, 0x38,
	0x3d, 0x8a, 0xdf, 0x23, 0xe4, 0x08, 0x20, 0x12, 0xe8, 0x88, 0x48, 0xad, 0xa0, 0xd0, 0x17, 0x49,
	0xac, 0x1f, 0x2c, 0x41, 0xd5, 0x91, 0xce, 0xc5, 0xa3, 0xa1, 0x65, 0x99, 0x03, 0x7b, 0xd4, 0x6d,
	0xdb, 0xa6, 0x61, 0x15, 0x6f, 0xb1, 0xb6, 0xc4, 0x27, 0x68, 0xd3, 0x15, 0x84, 0x46, 0x24, 0xcd,
	0xff, 0x01, 0x88, 0xd0, 0x56, 0x95, 0xae, 0x27, 0xb1, 0xfe, 0xe7, 0xe3, 0x84, 0x8b, 0x63, 0xfd,
	0xa1, 0x0d, 0xb7, 0x51, 0x31, 0xb8, 0xe0, 0x92, 0x8f, 0x22, 0xe1, 0x69, 0xeb, 0xca, 0xf8, 0x27,
	0x89, 0xf5, 0xa7, 0x4e, 0x91, 0xc2, 0x39, 0x89, 0x3c, 0xd9, 0x34, 0xac, 0x35, 0xd5, 0x36, 0x14,
	0x1e, 0xb6, 0x11, 0x3a, 0x67, 0x22, 0x94, 0xa3, 0x74, 0x80, 0xb4, 0x0d, 0x65, 0x1c, 0x24, 0xb1,
	0xbe, 0xf7, 0xdc, 0x57, 0xb3, 0x48, 0x57, 0x54, 0xd0, 0x80, 0x4c, 0x00, 0x5b, 0xa8, 0xe8, 0x91,
	0x39, 0xba, 0xf9, 0x2b, 0xe8, 0x5a, 0xea, 0x28, 0xb3, 0x8f, 0xf2, 0xc4, 0x63, 0x24, 0xd4, 0x7e,
	0x57, 0xde, 0xff, 0x49, 0xac, 0xef, 0x3f, 0xd7, 0x9b, 0xef, 0x7c, 0x67, 0xc7, 0xb0, 0x66, 0x0a,
	0x7e, 0x89, 0x56, 0x1d, 0x1e, 0xf9, 0x52, 0x4c, 0xb5, 0x2d, 0x05, 0x3e, 0x31, 0xd7, 0x8d, 0xfb,
	0x69, 0xe6, 0x2d, 0xf8, 0x33, 0xc2, 0xe4, 0x1a, 0x04, 0x71, 0xe1, 0x1d, 0xc0, 0x95, 0x37, 0x7d,
	0x13, 0x80, 0x4f, 0xb5, 0xed, 0x52, 0xa6, 0x9c, 0xe9, 0xd4, 0x93, 0x58, 0xdf, 0x5d, 0x36, 0x24,
	0x3c, 0x1a, 0x7b, 0x70, 0x97, 0xa9, 0x6e, 0x58, 0x3f, 0xb1, 0xf0, 0x27, 0xb4, 0x0e, 0x37, 0x01,
	0xf8, 0x21, 0x74, 0x88, 0x4f, 0x43, 0x0d, 0xab, 0x90, 0x87, 0x49, 0xac, 0x37, 0x96, 0xd8, 0x12,
	0x6e, 0xe4, 0x5d, 0xde, 0x07, 0xdb, 0x7e, 0xa0, 0x19, 0x15, 0x54, 0x98, 0x7d, 0xce, 0x78, 0x0d,
	0xe5, 0xfa, 0xed, 0x53, 0x73, 0xeb, 0x37, 0x8c, 0x50, 0xe1, 0xd8, 0x54, 0x75, 0x06, 0x17, 0x51,
	0xfe, 0xcc, 0x7e, 0x65, 0x5a, 0x5b, 0xd9, 0x71, 0x41, 0xfd, 0xbe, 0xf6, 0xbf, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x9f, 0x08, 0x55, 0xcf, 0x27, 0x05, 0x00, 0x00,
}