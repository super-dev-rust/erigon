// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package remote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TXPOOLClient is the client API for TXPOOL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TXPOOLClient interface {
	Add(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*AddReply, error)
}

type tXPOOLClient struct {
	cc grpc.ClientConnInterface
}

func NewTXPOOLClient(cc grpc.ClientConnInterface) TXPOOLClient {
	return &tXPOOLClient{cc}
}

func (c *tXPOOLClient) Add(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/remote.TXPOOL/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TXPOOLServer is the server API for TXPOOL service.
// All implementations must embed UnimplementedTXPOOLServer
// for forward compatibility
type TXPOOLServer interface {
	Add(context.Context, *TxRequest) (*AddReply, error)
	mustEmbedUnimplementedTXPOOLServer()
}

// UnimplementedTXPOOLServer must be embedded to have forward compatible implementations.
type UnimplementedTXPOOLServer struct {
}

func (*UnimplementedTXPOOLServer) Add(context.Context, *TxRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTXPOOLServer) mustEmbedUnimplementedTXPOOLServer() {}

func RegisterTXPOOLServer(s *grpc.Server, srv TXPOOLServer) {
	s.RegisterService(&_TXPOOL_serviceDesc, srv)
}

func _TXPOOL_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TXPOOLServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.TXPOOL/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TXPOOLServer).Add(ctx, req.(*TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TXPOOL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.TXPOOL",
	HandlerType: (*TXPOOLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TXPOOL_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote/txpool.proto",
}