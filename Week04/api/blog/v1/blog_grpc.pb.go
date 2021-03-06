// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogClient interface {
	Publish(ctx context.Context, in *PublishReply, opts ...grpc.CallOption) (*PublishReply, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) Publish(ctx context.Context, in *PublishReply, opts ...grpc.CallOption) (*PublishReply, error) {
	out := new(PublishReply)
	err := c.cc.Invoke(ctx, "/Blog/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServer is the server API for Blog service.
// All implementations must embed UnimplementedBlogServer
// for forward compatibility
type BlogServer interface {
	Publish(context.Context, *PublishReply) (*PublishReply, error)
	mustEmbedUnimplementedBlogServer()
}

// UnimplementedBlogServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (UnimplementedBlogServer) Publish(context.Context, *PublishReply) (*PublishReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedBlogServer) mustEmbedUnimplementedBlogServer() {}

// UnsafeBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServer will
// result in compilation errors.
type UnsafeBlogServer interface {
	mustEmbedUnimplementedBlogServer()
}

func RegisterBlogServer(s grpc.ServiceRegistrar, srv BlogServer) {
	s.RegisterService(&_Blog_serviceDesc, srv)
}

func _Blog_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishReply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Blog/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).Publish(ctx, req.(*PublishReply))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Blog_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/blog/blog.proto",
}
