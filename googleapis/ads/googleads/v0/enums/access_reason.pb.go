// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/access_reason.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible access reasons.
type AccessReasonEnum_AccessReason int32

const (
	// Not specified.
	AccessReasonEnum_UNSPECIFIED AccessReasonEnum_AccessReason = 0
	// Used for return value only. Represents value unknown in this version.
	AccessReasonEnum_UNKNOWN AccessReasonEnum_AccessReason = 1
	// The entity is owned by the user.
	AccessReasonEnum_OWNED AccessReasonEnum_AccessReason = 2
	// The entity is shared to the user.
	AccessReasonEnum_SHARED AccessReasonEnum_AccessReason = 3
	// The entity is licensed to the user.
	AccessReasonEnum_LICENSED AccessReasonEnum_AccessReason = 4
	// The user subscribed to the entity.
	AccessReasonEnum_SUBSCRIBED AccessReasonEnum_AccessReason = 5
	// The entity is accessible to the user.
	AccessReasonEnum_AFFILIATED AccessReasonEnum_AccessReason = 6
)

var AccessReasonEnum_AccessReason_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OWNED",
	3: "SHARED",
	4: "LICENSED",
	5: "SUBSCRIBED",
	6: "AFFILIATED",
}
var AccessReasonEnum_AccessReason_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"OWNED":       2,
	"SHARED":      3,
	"LICENSED":    4,
	"SUBSCRIBED":  5,
	"AFFILIATED":  6,
}

func (x AccessReasonEnum_AccessReason) String() string {
	return proto.EnumName(AccessReasonEnum_AccessReason_name, int32(x))
}
func (AccessReasonEnum_AccessReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_access_reason_757eb5f4e3d691fd, []int{0, 0}
}

// Indicates the way the entity such as user list is related to a user.
type AccessReasonEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessReasonEnum) Reset()         { *m = AccessReasonEnum{} }
func (m *AccessReasonEnum) String() string { return proto.CompactTextString(m) }
func (*AccessReasonEnum) ProtoMessage()    {}
func (*AccessReasonEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_reason_757eb5f4e3d691fd, []int{0}
}
func (m *AccessReasonEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessReasonEnum.Unmarshal(m, b)
}
func (m *AccessReasonEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessReasonEnum.Marshal(b, m, deterministic)
}
func (dst *AccessReasonEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessReasonEnum.Merge(dst, src)
}
func (m *AccessReasonEnum) XXX_Size() int {
	return xxx_messageInfo_AccessReasonEnum.Size(m)
}
func (m *AccessReasonEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessReasonEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccessReasonEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccessReasonEnum)(nil), "google.ads.googleads.v0.enums.AccessReasonEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AccessReasonEnum_AccessReason", AccessReasonEnum_AccessReason_name, AccessReasonEnum_AccessReason_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/access_reason.proto", fileDescriptor_access_reason_757eb5f4e3d691fd)
}

var fileDescriptor_access_reason_757eb5f4e3d691fd = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0xc2, 0x30,
	0x1c, 0xc6, 0x1d, 0x08, 0xea, 0x1f, 0xa2, 0xb5, 0x77, 0x0e, 0xf2, 0x00, 0xdd, 0x8c, 0x47, 0x4f,
	0x1d, 0x2d, 0xd8, 0x48, 0x0a, 0xd9, 0x1c, 0x24, 0x66, 0x89, 0x99, 0xdb, 0xd2, 0x98, 0xb0, 0x15,
	0x57, 0xe1, 0x0d, 0x7c, 0x11, 0x8f, 0x3e, 0x86, 0x47, 0x9f, 0xca, 0xac, 0x13, 0xc2, 0x45, 0x2f,
	0xcd, 0xd7, 0x7e, 0xfd, 0xb5, 0xff, 0xef, 0x83, 0x6b, 0xa5, 0xb5, 0x5a, 0xe5, 0x6e, 0x92, 0x19,
	0xb7, 0x91, 0xb5, 0xda, 0x7a, 0x6e, 0x5e, 0x6e, 0x0a, 0xe3, 0x26, 0x69, 0x9a, 0x1b, 0xf3, 0x54,
	0xe5, 0x89, 0xd1, 0x25, 0x59, 0x57, 0xfa, 0x4d, 0xe3, 0x41, 0x73, 0x8f, 0x24, 0x99, 0x21, 0x7b,
	0x84, 0x6c, 0x3d, 0x62, 0x91, 0xe1, 0xbb, 0x03, 0x88, 0x5a, 0x2c, 0xb0, 0x14, 0x2f, 0x37, 0xc5,
	0xf0, 0x15, 0xfa, 0x87, 0x67, 0xf8, 0x02, 0x7a, 0x91, 0x0c, 0xe7, 0x7c, 0x24, 0xc6, 0x82, 0x33,
	0x74, 0x84, 0x7b, 0x70, 0x12, 0xc9, 0x7b, 0x39, 0x5b, 0x4a, 0xe4, 0xe0, 0x33, 0xe8, 0xcc, 0x96,
	0x92, 0x33, 0xd4, 0xc2, 0x00, 0xdd, 0xf0, 0x8e, 0x06, 0x9c, 0xa1, 0x36, 0xee, 0xc3, 0xe9, 0x54,
	0x8c, 0xb8, 0x0c, 0x39, 0x43, 0xc7, 0xf8, 0x1c, 0x20, 0x8c, 0xfc, 0x70, 0x14, 0x08, 0x9f, 0x33,
	0xd4, 0xa9, 0xf7, 0x74, 0x3c, 0x16, 0x53, 0x41, 0x1f, 0x38, 0x43, 0x5d, 0xff, 0xcb, 0x81, 0xab,
	0x54, 0x17, 0xe4, 0xdf, 0x69, 0xfd, 0xcb, 0xc3, 0xb1, 0xe6, 0x75, 0xbe, 0xb9, 0xf3, 0xe8, 0xff,
	0x32, 0x4a, 0xaf, 0x92, 0x52, 0x11, 0x5d, 0x29, 0x57, 0xe5, 0xa5, 0x4d, 0xbf, 0x2b, 0x69, 0xfd,
	0x62, 0xfe, 0xe8, 0xec, 0xd6, 0xae, 0x1f, 0xad, 0xf6, 0x84, 0xd2, 0xcf, 0xd6, 0x60, 0xd2, 0x3c,
	0x45, 0x33, 0x43, 0x1a, 0x59, 0xab, 0x85, 0x47, 0xea, 0x5a, 0xcc, 0xf7, 0xce, 0x8f, 0x69, 0x66,
	0xe2, 0xbd, 0x1f, 0x2f, 0xbc, 0xd8, 0xfa, 0xcf, 0x5d, 0xfb, 0xe9, 0xcd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0x24, 0x60, 0xa5, 0xa7, 0x01, 0x00, 0x00,
}
