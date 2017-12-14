// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail.proto

/*
Package mail is a generated protocol buffer package.

It is generated from these files:
	mail.proto

It has these top-level messages:
	MailRequest
	MailResponse
*/
package mail

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MailRequest struct {
	Mail string `protobuf:"bytes,1,opt,name=mail" json:"mail,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *MailRequest) Reset()                    { *m = MailRequest{} }
func (m *MailRequest) String() string            { return proto.CompactTextString(m) }
func (*MailRequest) ProtoMessage()               {}
func (*MailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MailRequest) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *MailRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type MailResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *MailResponse) Reset()                    { *m = MailResponse{} }
func (m *MailResponse) String() string            { return proto.CompactTextString(m) }
func (*MailResponse) ProtoMessage()               {}
func (*MailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MailResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*MailRequest)(nil), "mail.MailRequest")
	proto.RegisterType((*MailResponse)(nil), "mail.MailResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Mail service

type MailClient interface {
	SendMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error)
}

type mailClient struct {
	cc *grpc.ClientConn
}

func NewMailClient(cc *grpc.ClientConn) MailClient {
	return &mailClient{cc}
}

func (c *mailClient) SendMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error) {
	out := new(MailResponse)
	err := grpc.Invoke(ctx, "/mail.Mail/SendMail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mail service

type MailServer interface {
	SendMail(context.Context, *MailRequest) (*MailResponse, error)
}

func RegisterMailServer(s *grpc.Server, srv MailServer) {
	s.RegisterService(&_Mail_serviceDesc, srv)
}

func _Mail_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mail.Mail/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).SendMail(ctx, req.(*MailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mail.Mail",
	HandlerType: (*MailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMail",
			Handler:    _Mail_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mail.proto",
}

func init() { proto.RegisterFile("mail.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0xcc, 0xcc,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x8c, 0xb9, 0xb8, 0x7d, 0x13,
	0x33, 0x73, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0xc0, 0xc2, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x69, 0x51, 0x8e, 0x04, 0x13,
	0x58, 0x08, 0xc4, 0x54, 0xd2, 0xe0, 0xe2, 0x81, 0x68, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0x06, 0x6b, 0xe4, 0x08, 0x82, 0x71,
	0x8d, 0xac, 0xb9, 0x58, 0x40, 0x2a, 0x85, 0x8c, 0xb9, 0x38, 0x82, 0x53, 0xf3, 0x52, 0xc0, 0x6c,
	0x41, 0x3d, 0xb0, 0x2b, 0x90, 0xac, 0x95, 0x12, 0x42, 0x16, 0x82, 0x18, 0xaa, 0xc4, 0x90, 0xc4,
	0x06, 0x76, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x27, 0x3e, 0xd7, 0xc3, 0xb6, 0x00, 0x00,
	0x00,
}
