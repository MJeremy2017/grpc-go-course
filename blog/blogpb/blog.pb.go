// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blogpb/blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{1}
}

func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(m, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{2}
}

func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (m *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(m, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blog.createBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blog.createBlogResponse")
}

func init() { proto.RegisterFile("blog/blogpb/blog.proto", fileDescriptor_a4b0406114889fe6) }

var fileDescriptor_a4b0406114889fe6 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xca, 0xc9, 0x4f,
	0xd7, 0x07, 0x11, 0x05, 0x49, 0x60, 0x4a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x05, 0xc4,
	0x56, 0x4a, 0xe6, 0x62, 0x71, 0xca, 0xc9, 0x4f, 0x17, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe6, 0xe2, 0x4c, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0x8a, 0xcf, 0x4c, 0x91, 0x60, 0x02, 0x0b, 0x73, 0x40, 0x04, 0x3c, 0x53, 0x84, 0x44,
	0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x98, 0xc1, 0x12, 0x10, 0x8e, 0x90, 0x04, 0x17,
	0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0x0b, 0x58, 0x1c, 0xc6, 0x55, 0x32, 0xe6,
	0x12, 0x4c, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x05, 0x59, 0x15, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0x24, 0xc7, 0x05, 0x76, 0x01, 0xd8, 0x4e, 0x6e, 0x23, 0x2e, 0x3d, 0xb0, 0xd3, 0xc0, 0x0a,
	0x20, 0x2e, 0x33, 0xe1, 0x12, 0x42, 0xd6, 0x54, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x4a, 0x48, 0x97,
	0x51, 0x00, 0x17, 0x37, 0x88, 0x17, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc8, 0xc5,
	0xe5, 0x0c, 0x37, 0x44, 0x48, 0x1c, 0xa2, 0x1c, 0xc3, 0x2d, 0x52, 0x12, 0x98, 0x12, 0x10, 0xfb,
	0x94, 0x18, 0x9c, 0x38, 0xa2, 0xd8, 0x20, 0x81, 0x97, 0xc4, 0x06, 0x0e, 0x38, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xdb, 0xc5, 0xbe, 0x94, 0x52, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(ctx context.Context, req *CreateBlogRequest) (*CreateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/blogpb/blog.proto",
}