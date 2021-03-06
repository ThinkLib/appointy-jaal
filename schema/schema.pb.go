// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package schema

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type MethodOptions struct {
	// Types that are valid to be assigned to Type:
	//	*MethodOptions_Query
	//	*MethodOptions_Mutation
	//	*MethodOptions_Subscription
	Type                 isMethodOptions_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MethodOptions) Reset()         { *m = MethodOptions{} }
func (m *MethodOptions) String() string { return proto.CompactTextString(m) }
func (*MethodOptions) ProtoMessage()    {}
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *MethodOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodOptions.Unmarshal(m, b)
}
func (m *MethodOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodOptions.Marshal(b, m, deterministic)
}
func (m *MethodOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodOptions.Merge(m, src)
}
func (m *MethodOptions) XXX_Size() int {
	return xxx_messageInfo_MethodOptions.Size(m)
}
func (m *MethodOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MethodOptions proto.InternalMessageInfo

type isMethodOptions_Type interface {
	isMethodOptions_Type()
}

type MethodOptions_Query struct {
	Query string `protobuf:"bytes,1,opt,name=query,proto3,oneof"`
}

type MethodOptions_Mutation struct {
	Mutation string `protobuf:"bytes,2,opt,name=mutation,proto3,oneof"`
}

type MethodOptions_Subscription struct {
	Subscription string `protobuf:"bytes,3,opt,name=subscription,proto3,oneof"`
}

func (*MethodOptions_Query) isMethodOptions_Type() {}

func (*MethodOptions_Mutation) isMethodOptions_Type() {}

func (*MethodOptions_Subscription) isMethodOptions_Type() {}

func (m *MethodOptions) GetType() isMethodOptions_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *MethodOptions) GetQuery() string {
	if x, ok := m.GetType().(*MethodOptions_Query); ok {
		return x.Query
	}
	return ""
}

func (m *MethodOptions) GetMutation() string {
	if x, ok := m.GetType().(*MethodOptions_Mutation); ok {
		return x.Mutation
	}
	return ""
}

func (m *MethodOptions) GetSubscription() string {
	if x, ok := m.GetType().(*MethodOptions_Subscription); ok {
		return x.Subscription
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodOptions) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodOptions_Query)(nil),
		(*MethodOptions_Mutation)(nil),
		(*MethodOptions_Subscription)(nil),
	}
}

var E_Schema = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*MethodOptions)(nil),
	Field:         91111,
	Name:          "graphql.schema",
	Tag:           "bytes,91111,opt,name=schema",
	Filename:      "schema.proto",
}

var E_Skip = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         91112,
	Name:          "graphql.skip",
	Tag:           "varint,91112,opt,name=skip",
	Filename:      "schema.proto",
}

var E_FileSkip = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         91113,
	Name:          "graphql.file_skip",
	Tag:           "varint,91113,opt,name=file_skip",
	Filename:      "schema.proto",
}

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         91114,
	Name:          "graphql.name",
	Tag:           "bytes,91114,opt,name=name",
	Filename:      "schema.proto",
}

func init() {
	proto.RegisterType((*MethodOptions)(nil), "graphql.MethodOptions")
	proto.RegisterExtension(E_Schema)
	proto.RegisterExtension(E_Skip)
	proto.RegisterExtension(E_FileSkip)
	proto.RegisterExtension(E_Name)
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x09, 0xa4, 0x21, 0x31, 0x65, 0xf1, 0x50, 0x45, 0x50, 0x20, 0x42, 0x0c, 0x9d, 0x1c,
	0x09, 0xd4, 0x25, 0x6c, 0x1d, 0x10, 0x0b, 0x02, 0x85, 0x8d, 0x05, 0x39, 0xa9, 0x9b, 0x18, 0x92,
	0x5c, 0x37, 0x76, 0x86, 0x3c, 0x21, 0x8f, 0xc2, 0xcf, 0x53, 0xa0, 0xd8, 0x69, 0xa5, 0x08, 0xa4,
	0x8e, 0xbe, 0xf7, 0x7c, 0x9f, 0xcf, 0x45, 0x63, 0x99, 0xe6, 0xac, 0xa4, 0x44, 0xd4, 0xa0, 0x00,
	0x1f, 0x66, 0x35, 0x15, 0xf9, 0xba, 0x38, 0x09, 0x32, 0x80, 0xac, 0x60, 0xa1, 0x1e, 0x27, 0xcd,
	0x2a, 0x5c, 0x32, 0x99, 0xd6, 0x5c, 0x28, 0xa8, 0x4d, 0xf4, 0x52, 0xa2, 0xe3, 0x07, 0xa6, 0x72,
	0x58, 0x3e, 0x0a, 0xc5, 0xa1, 0x92, 0x78, 0x82, 0x46, 0xeb, 0x86, 0xd5, 0xad, 0x6f, 0x05, 0xd6,
	0xcc, 0xbb, 0xdf, 0x8b, 0xcd, 0x13, 0x4f, 0x91, 0x5b, 0x36, 0x8a, 0x76, 0x21, 0x7f, 0xbf, 0x5f,
	0x6d, 0x27, 0xf8, 0x0a, 0x8d, 0x65, 0x93, 0x18, 0x77, 0x97, 0x38, 0xe8, 0x13, 0x83, 0xe9, 0xc2,
	0x41, 0xb6, 0x6a, 0x05, 0x8b, 0x9e, 0x90, 0x63, 0xfa, 0xe2, 0x73, 0x62, 0x1a, 0x92, 0x4d, 0x43,
	0x32, 0x68, 0xe3, 0x7f, 0x7e, 0x8c, 0x02, 0x6b, 0x76, 0x74, 0x3d, 0x21, 0xfd, 0x49, 0xc3, 0x7d,
	0xdc, 0x7b, 0xa2, 0x39, 0xb2, 0xe5, 0x3b, 0x17, 0xf8, 0xe2, 0x1f, 0x9f, 0x94, 0x34, 0x63, 0x1b,
	0xe1, 0x97, 0x16, 0xba, 0xb1, 0x8e, 0x47, 0xb7, 0xc8, 0x5b, 0xf1, 0x82, 0xbd, 0x6a, 0x76, 0xfa,
	0x87, 0xbd, 0xe3, 0xc5, 0x16, 0xfc, 0xee, 0x41, 0xb7, 0x03, 0x9e, 0x3b, 0x78, 0x8e, 0xec, 0x8a,
	0x96, 0x6c, 0xf7, 0x9f, 0x3f, 0x1a, 0xf5, 0x62, 0x1d, 0x5f, 0x9c, 0xbd, 0x9c, 0x66, 0x40, 0xa8,
	0x10, 0xc0, 0x2b, 0xd5, 0x92, 0x14, 0xca, 0xf0, 0x8d, 0xd2, 0x22, 0x34, 0x97, 0x24, 0x8e, 0xb6,
	0xdc, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x73, 0x66, 0xd8, 0xd2, 0x01, 0x00, 0x00,
}
