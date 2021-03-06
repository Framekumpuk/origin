// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_client.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A link between the given customer and a client customer. CustomerClients only
// exist for manager customers. All direct and indirect client customers are
// included, as well as the manager itself.
type CustomerClient struct {
	// The resource name of the customer client.
	// CustomerClient resource names have the form:
	// `customers/{customer_id}/customerClients/{client_customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The resource name of the client-customer which is linked to
	// the given customer. Read only.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// Specifies whether this is a hidden account. Learn more about hidden
	// accounts <a
	// href="https://support.google.com/google-ads/answer/7519830">here</a>. Read
	// only.
	Hidden *wrappers.BoolValue `protobuf:"bytes,4,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// Distance between given customer and client. For self link, the level value
	// will be 0. Read only.
	Level                *wrappers.Int64Value `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CustomerClient) Reset()         { *m = CustomerClient{} }
func (m *CustomerClient) String() string { return proto.CompactTextString(m) }
func (*CustomerClient) ProtoMessage()    {}
func (*CustomerClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_client_bb55b7230747a822, []int{0}
}
func (m *CustomerClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClient.Unmarshal(m, b)
}
func (m *CustomerClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClient.Marshal(b, m, deterministic)
}
func (dst *CustomerClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClient.Merge(dst, src)
}
func (m *CustomerClient) XXX_Size() int {
	return xxx_messageInfo_CustomerClient.Size(m)
}
func (m *CustomerClient) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClient.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClient proto.InternalMessageInfo

func (m *CustomerClient) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerClient) GetClientCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ClientCustomer
	}
	return nil
}

func (m *CustomerClient) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

func (m *CustomerClient) GetLevel() *wrappers.Int64Value {
	if m != nil {
		return m.Level
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerClient)(nil), "google.ads.googleads.v1.resources.CustomerClient")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_client.proto", fileDescriptor_customer_client_bb55b7230747a822)
}

var fileDescriptor_customer_client_bb55b7230747a822 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x00, 0xc5, 0x69, 0xf7, 0xdf, 0xe0, 0x5f, 0x75, 0x42, 0xbd, 0x94, 0x39, 0x64, 0x53, 0x06, 0x3b,
	0xa5, 0x74, 0x8a, 0x42, 0x3c, 0x75, 0x43, 0x86, 0x1e, 0x64, 0x4c, 0xe8, 0x41, 0x0a, 0x23, 0x6b,
	0x63, 0x2d, 0xa4, 0x49, 0x49, 0xd2, 0x79, 0xf5, 0xb3, 0x78, 0xf4, 0xa3, 0xf8, 0x35, 0xbc, 0xf9,
	0x29, 0x64, 0x4d, 0x13, 0x90, 0x81, 0xde, 0x1e, 0xc9, 0xfb, 0xbd, 0xf7, 0x48, 0x9c, 0xab, 0x8c,
	0xb1, 0x8c, 0x60, 0x1f, 0xa5, 0xc2, 0x57, 0x72, 0xab, 0x36, 0x81, 0xcf, 0xb1, 0x60, 0x15, 0x4f,
	0xb0, 0xf0, 0x93, 0x4a, 0x48, 0x56, 0x60, 0xbe, 0x4a, 0x48, 0x8e, 0xa9, 0x04, 0x25, 0x67, 0x92,
	0xb9, 0x43, 0xe5, 0x06, 0x28, 0x15, 0xc0, 0x80, 0x60, 0x13, 0x00, 0x03, 0xf6, 0x4e, 0x9a, 0xec,
	0x1a, 0x58, 0x57, 0x4f, 0xfe, 0x0b, 0x47, 0x65, 0x89, 0xb9, 0x50, 0x11, 0xbd, 0xbe, 0xee, 0x2e,
	0x73, 0x1f, 0x51, 0xca, 0x24, 0x92, 0x39, 0xa3, 0xcd, 0xed, 0xe9, 0xa7, 0xe5, 0x74, 0x67, 0x4d,
	0xf5, 0xac, 0x6e, 0x76, 0xcf, 0x9c, 0x03, 0x9d, 0xbe, 0xa2, 0xa8, 0xc0, 0x9e, 0x35, 0xb0, 0xc6,
	0xff, 0x97, 0xfb, 0xfa, 0xf0, 0x1e, 0x15, 0xd8, 0xbd, 0x71, 0x0e, 0xd5, 0xd0, 0x95, 0x1e, 0xee,
	0xb5, 0x06, 0xd6, 0x78, 0x6f, 0xd2, 0x6f, 0x76, 0x02, 0xbd, 0x07, 0x3c, 0x48, 0x9e, 0xd3, 0x2c,
	0x42, 0xa4, 0xc2, 0xcb, 0xae, 0x82, 0x74, 0xa3, 0x3b, 0x71, 0x3a, 0xcf, 0x79, 0x9a, 0x62, 0xea,
	0xfd, 0xab, 0xe9, 0xde, 0x0e, 0x3d, 0x65, 0x8c, 0x28, 0xb6, 0x71, 0xba, 0x81, 0xd3, 0x26, 0x78,
	0x83, 0x89, 0xd7, 0xae, 0x91, 0xe3, 0x1d, 0xe4, 0x96, 0xca, 0xcb, 0x0b, 0xc5, 0x28, 0xe7, 0xf4,
	0xd5, 0x76, 0x46, 0x09, 0x2b, 0xc0, 0x9f, 0xaf, 0x39, 0x3d, 0xfa, 0xf9, 0x18, 0x8b, 0x6d, 0xe6,
	0xc2, 0x7a, 0xbc, 0x6b, 0xc8, 0x8c, 0x11, 0x44, 0x33, 0xc0, 0x78, 0xe6, 0x67, 0x98, 0xd6, 0x8d,
	0xfa, 0x43, 0xcb, 0x5c, 0xfc, 0xf2, 0xbf, 0xd7, 0x46, 0xbd, 0xd9, 0xad, 0x79, 0x18, 0xbe, 0xdb,
	0xc3, 0xb9, 0x8a, 0x0c, 0x53, 0x01, 0x94, 0xdc, 0xaa, 0x28, 0x00, 0x4b, 0xed, 0xfc, 0xd0, 0x9e,
	0x38, 0x4c, 0x45, 0x6c, 0x3c, 0x71, 0x14, 0xc4, 0xc6, 0xf3, 0x65, 0x8f, 0xd4, 0x05, 0x84, 0x61,
	0x2a, 0x20, 0x34, 0x2e, 0x08, 0xa3, 0x00, 0x42, 0xe3, 0x5b, 0x77, 0xea, 0xb1, 0xe7, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0x0b, 0x0a, 0xfa, 0x8b, 0x02, 0x00, 0x00,
}
