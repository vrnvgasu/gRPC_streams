// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.13.0
// source: server/api/api.proto

package pdfcompose

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PdfComposeService_Send_FullMethodName = "/pdfcompose.PdfComposeService/Send"
)

// PdfComposeServiceClient is the client API for PdfComposeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PdfComposeServiceClient interface {
	Send(ctx context.Context, in *FormData, opts ...grpc.CallOption) (*File, error)
}

type pdfComposeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPdfComposeServiceClient(cc grpc.ClientConnInterface) PdfComposeServiceClient {
	return &pdfComposeServiceClient{cc}
}

func (c *pdfComposeServiceClient) Send(ctx context.Context, in *FormData, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, PdfComposeService_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PdfComposeServiceServer is the server API for PdfComposeService service.
// All implementations must embed UnimplementedPdfComposeServiceServer
// for forward compatibility
type PdfComposeServiceServer interface {
	Send(context.Context, *FormData) (*File, error)
	mustEmbedUnimplementedPdfComposeServiceServer()
}

// UnimplementedPdfComposeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPdfComposeServiceServer struct {
}

func (UnimplementedPdfComposeServiceServer) Send(context.Context, *FormData) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedPdfComposeServiceServer) mustEmbedUnimplementedPdfComposeServiceServer() {}

// UnsafePdfComposeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PdfComposeServiceServer will
// result in compilation errors.
type UnsafePdfComposeServiceServer interface {
	mustEmbedUnimplementedPdfComposeServiceServer()
}

func RegisterPdfComposeServiceServer(s grpc.ServiceRegistrar, srv PdfComposeServiceServer) {
	s.RegisterService(&PdfComposeService_ServiceDesc, srv)
}

func _PdfComposeService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PdfComposeServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PdfComposeService_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PdfComposeServiceServer).Send(ctx, req.(*FormData))
	}
	return interceptor(ctx, in, info, handler)
}

// PdfComposeService_ServiceDesc is the grpc.ServiceDesc for PdfComposeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PdfComposeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdfcompose.PdfComposeService",
	HandlerType: (*PdfComposeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _PdfComposeService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/api/api.proto",
}
