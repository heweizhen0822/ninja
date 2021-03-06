// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

/*
Package blog is a generated protocol buffer package.

It is generated from these files:
	article.proto

It has these top-level messages:
	HelloReq
	HelloResp
*/
package blog

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

type HelloReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloReq) Reset()                    { *m = HelloReq{} }
func (m *HelloReq) String() string            { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()               {}
func (*HelloReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResp struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *HelloResp) Reset()                    { *m = HelloResp{} }
func (m *HelloResp) String() string            { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()               {}
func (*HelloResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "blog.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "blog.HelloResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Article service

type ArticleClient interface {
	Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type articleClient struct {
	cc *grpc.ClientConn
}

func NewArticleClient(cc *grpc.ClientConn) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := grpc.Invoke(ctx, "/blog.Article/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Article service

type ArticleServer interface {
	Hello(context.Context, *HelloReq) (*HelloResp, error)
}

func RegisterArticleServer(s *grpc.Server, srv ArticleServer) {
	s.RegisterService(&_Article_serviceDesc, srv)
}

func _Article_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.Article/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Hello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Article_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Article_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}

func init() { proto.RegisterFile("article.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0x2a, 0xc9,
	0x4c, 0xce, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xca, 0xc9, 0x4f, 0x57,
	0x92, 0xe3, 0xe2, 0xf0, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0xc9,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x64, 0xb9, 0x38,
	0xa1, 0xf2, 0xc5, 0x05, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x50, 0x79, 0x10, 0xd3, 0xc8,
	0x98, 0x8b, 0xdd, 0x11, 0x62, 0xaa, 0x90, 0x06, 0x17, 0x2b, 0x58, 0xa5, 0x10, 0x9f, 0x1e, 0xc8,
	0x64, 0x3d, 0x98, 0xb1, 0x52, 0xfc, 0x28, 0xfc, 0xe2, 0x82, 0x24, 0x36, 0xb0, 0x03, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xe3, 0xd5, 0xfe, 0x91, 0x00, 0x00, 0x00,
}
