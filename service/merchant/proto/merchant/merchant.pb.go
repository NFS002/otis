// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/merchant/merchant.proto

package merchant

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Merchant struct {
	MerchantID           string      `protobuf:"bytes,1,opt,name=merchantID,proto3" json:"merchantID,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Locations            []*Location `protobuf:"bytes,3,rep,name=locations,proto3" json:"locations,omitempty"`
	Tags                 []*Tag      `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	ContactPhone         string      `protobuf:"bytes,5,opt,name=contactPhone,proto3" json:"contactPhone,omitempty"`
	ContactEmail         string      `protobuf:"bytes,6,opt,name=contactEmail,proto3" json:"contactEmail,omitempty"`
	Rate                 float32     `protobuf:"fixed32,7,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Merchant) Reset()         { *m = Merchant{} }
func (m *Merchant) String() string { return proto.CompactTextString(m) }
func (*Merchant) ProtoMessage()    {}
func (*Merchant) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{0}
}

func (m *Merchant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant.Unmarshal(m, b)
}
func (m *Merchant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant.Marshal(b, m, deterministic)
}
func (m *Merchant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant.Merge(m, src)
}
func (m *Merchant) XXX_Size() int {
	return xxx_messageInfo_Merchant.Size(m)
}
func (m *Merchant) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant proto.InternalMessageInfo

func (m *Merchant) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

func (m *Merchant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Merchant) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *Merchant) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Merchant) GetContactPhone() string {
	if m != nil {
		return m.ContactPhone
	}
	return ""
}

func (m *Merchant) GetContactEmail() string {
	if m != nil {
		return m.ContactEmail
	}
	return ""
}

func (m *Merchant) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

type Location struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type Tag struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{2}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type CreateResponse struct {
	Created              bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Merchant             *Merchant `protobuf:"bytes,2,opt,name=merchant,proto3" json:"merchant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *CreateResponse) GetMerchant() *Merchant {
	if m != nil {
		return m.Merchant
	}
	return nil
}

type GetRequest struct {
	MerchantID           string   `protobuf:"bytes,1,opt,name=merchantID,proto3" json:"merchantID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

type GetResponse struct {
	Merchants            []*Merchant `protobuf:"bytes,1,rep,name=merchants,proto3" json:"merchants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetMerchants() []*Merchant {
	if m != nil {
		return m.Merchants
	}
	return nil
}

type UpdateResponse struct {
	Updated              bool      `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	Merchant             *Merchant `protobuf:"bytes,2,opt,name=merchant,proto3" json:"merchant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetUpdated() bool {
	if m != nil {
		return m.Updated
	}
	return false
}

func (m *UpdateResponse) GetMerchant() *Merchant {
	if m != nil {
		return m.Merchant
	}
	return nil
}

type DeleteRequest struct {
	MerchantID           string   `protobuf:"bytes,1,opt,name=merchantID,proto3" json:"merchantID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

type DeleteResponse struct {
	Deleted              bool     `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db34436ab21601, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func init() {
	proto.RegisterType((*Merchant)(nil), "merchant.Merchant")
	proto.RegisterType((*Location)(nil), "merchant.Location")
	proto.RegisterType((*Tag)(nil), "merchant.Tag")
	proto.RegisterType((*CreateResponse)(nil), "merchant.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "merchant.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "merchant.GetResponse")
	proto.RegisterType((*UpdateResponse)(nil), "merchant.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "merchant.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "merchant.DeleteResponse")
}

func init() { proto.RegisterFile("proto/merchant/merchant.proto", fileDescriptor_e9db34436ab21601) }

var fileDescriptor_e9db34436ab21601 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x51, 0x6b, 0x1a, 0x41,
	0x10, 0xf6, 0x3c, 0xab, 0xe7, 0x58, 0x6d, 0x59, 0x5a, 0x5c, 0x84, 0x16, 0xbb, 0x0f, 0x45, 0x4a,
	0x51, 0xb1, 0xaf, 0x85, 0x08, 0x51, 0x42, 0x20, 0x81, 0xb0, 0x31, 0x2f, 0x79, 0xdb, 0x9c, 0xc3,
	0x29, 0xe8, 0x9d, 0xb9, 0x5b, 0xf3, 0x63, 0xf2, 0x27, 0xf3, 0x17, 0xc2, 0x0d, 0xb7, 0xbb, 0x9e,
	0x04, 0x62, 0xde, 0x66, 0xbe, 0xf9, 0x76, 0xe6, 0x9b, 0x6f, 0xee, 0xe0, 0xc7, 0x2e, 0x4d, 0x74,
	0x32, 0xda, 0x62, 0x1a, 0xae, 0x54, 0xac, 0x6d, 0x30, 0x24, 0x9c, 0x05, 0x26, 0x17, 0x2f, 0x1e,
	0x04, 0xd7, 0x45, 0xc2, 0x7e, 0x02, 0x98, 0xc2, 0xe5, 0x8c, 0x7b, 0x7d, 0x6f, 0xd0, 0x94, 0x07,
	0x08, 0x63, 0x50, 0x8b, 0xd5, 0x16, 0x79, 0x95, 0x2a, 0x14, 0xb3, 0x31, 0x34, 0x37, 0x49, 0xa8,
	0xf4, 0x3a, 0x89, 0x33, 0xee, 0xf7, 0xfd, 0x41, 0x6b, 0xc2, 0x86, 0x76, 0xdc, 0x55, 0x51, 0x92,
	0x8e, 0xc4, 0x7e, 0x41, 0x4d, 0xab, 0x28, 0xe3, 0x35, 0x22, 0xb7, 0x1d, 0x79, 0xa1, 0x22, 0x49,
	0x25, 0x26, 0xe0, 0x73, 0x98, 0xc4, 0x5a, 0x85, 0xfa, 0x66, 0x95, 0xc4, 0xc8, 0x3f, 0xd1, 0xc0,
	0x12, 0x76, 0xc0, 0x99, 0x6f, 0xd5, 0x7a, 0xc3, 0xeb, 0x25, 0x0e, 0x61, 0xb9, 0xe0, 0x54, 0x69,
	0xe4, 0x8d, 0xbe, 0x37, 0xa8, 0x4a, 0x8a, 0xc5, 0x6f, 0x08, 0x8c, 0x2a, 0xd6, 0x83, 0xc0, 0xe8,
	0x2a, 0xd6, 0xb5, 0xb9, 0xe8, 0x82, 0xbf, 0x50, 0x11, 0xfb, 0x0a, 0xbe, 0x56, 0x51, 0x51, 0xcd,
	0x43, 0x71, 0x0f, 0x9d, 0xf3, 0x14, 0x95, 0x46, 0x89, 0xd9, 0x2e, 0x89, 0x33, 0x64, 0x1c, 0x1a,
	0x21, 0x21, 0x4b, 0xe2, 0x05, 0xd2, 0xa4, 0x6c, 0x08, 0xd6, 0x6a, 0x72, 0xad, 0x64, 0x8e, 0xf1,
	0x5d, 0xba, 0x73, 0xfc, 0x05, 0xb8, 0x40, 0x2d, 0xf1, 0x71, 0x8f, 0xd9, 0xbb, 0xf7, 0x10, 0x67,
	0xd0, 0x22, 0x76, 0x21, 0x63, 0x0c, 0x4d, 0x53, 0xcc, 0xb8, 0x77, 0x7c, 0x0a, 0x3b, 0xcd, 0x91,
	0xf2, 0x55, 0xee, 0x76, 0xcb, 0xa3, 0x55, 0xf6, 0x84, 0xd8, 0x55, 0x8a, 0xf4, 0xc3, 0xab, 0x8c,
	0xa0, 0x3d, 0xc3, 0x0d, 0xe6, 0xbd, 0x4f, 0xdb, 0xe6, 0x0f, 0x74, 0xcc, 0x03, 0x27, 0x66, 0x49,
	0x88, 0x15, 0x53, 0xa4, 0x93, 0xe7, 0x2a, 0x7c, 0x31, 0x33, 0x6f, 0x31, 0x7d, 0x5a, 0x87, 0xc8,
	0xa6, 0xe6, 0x2e, 0xf6, 0x7b, 0x7e, 0x43, 0x60, 0x8f, 0x3b, 0xac, 0x7c, 0x45, 0x51, 0x61, 0xff,
	0xc9, 0x4f, 0xfb, 0xfc, 0x9b, 0xa3, 0xba, 0xa3, 0xf4, 0xbe, 0x1f, 0xa1, 0xf6, 0xf5, 0xd4, 0x98,
	0x79, 0xea, 0xfc, 0xb2, 0xf5, 0xa2, 0xc2, 0xe6, 0xc6, 0x01, 0xdb, 0xa1, 0xeb, 0xd8, 0x25, 0x33,
	0x0f, 0xdb, 0x94, 0x4d, 0x13, 0x95, 0x87, 0x3a, 0xfd, 0xe4, 0xff, 0x5e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x71, 0x9a, 0x8f, 0xb5, 0x05, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MerchantService service

type MerchantServiceClient interface {
	CreateMerchant(ctx context.Context, in *Merchant, opts ...client.CallOption) (*CreateResponse, error)
	GetMerchant(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	UpdateMerchant(ctx context.Context, in *Merchant, opts ...client.CallOption) (*UpdateResponse, error)
	DeleteMerchant(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type merchantServiceClient struct {
	c           client.Client
	serviceName string
}

func NewMerchantServiceClient(serviceName string, c client.Client) MerchantServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "merchant"
	}
	return &merchantServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *merchantServiceClient) CreateMerchant(ctx context.Context, in *Merchant, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "MerchantService.CreateMerchant", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetMerchant(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "MerchantService.GetMerchant", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) UpdateMerchant(ctx context.Context, in *Merchant, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "MerchantService.UpdateMerchant", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) DeleteMerchant(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "MerchantService.DeleteMerchant", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MerchantService service

type MerchantServiceHandler interface {
	CreateMerchant(context.Context, *Merchant, *CreateResponse) error
	GetMerchant(context.Context, *GetRequest, *GetResponse) error
	UpdateMerchant(context.Context, *Merchant, *UpdateResponse) error
	DeleteMerchant(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterMerchantServiceHandler(s server.Server, hdlr MerchantServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&MerchantService{hdlr}, opts...))
}

type MerchantService struct {
	MerchantServiceHandler
}

func (h *MerchantService) CreateMerchant(ctx context.Context, in *Merchant, out *CreateResponse) error {
	return h.MerchantServiceHandler.CreateMerchant(ctx, in, out)
}

func (h *MerchantService) GetMerchant(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.MerchantServiceHandler.GetMerchant(ctx, in, out)
}

func (h *MerchantService) UpdateMerchant(ctx context.Context, in *Merchant, out *UpdateResponse) error {
	return h.MerchantServiceHandler.UpdateMerchant(ctx, in, out)
}

func (h *MerchantService) DeleteMerchant(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.MerchantServiceHandler.DeleteMerchant(ctx, in, out)
}
