// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc2
// source: api/slacker/v1/record.proto

package v1

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
	Record_BeginRecord_FullMethodName = "/api.slacker.v1.Record/BeginRecord"
	Record_EndRecord_FullMethodName   = "/api.slacker.v1.Record/EndRecord"
)

// RecordClient is the client API for Record service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordClient interface {
	BeginRecord(ctx context.Context, in *BeginRecordRequest, opts ...grpc.CallOption) (*BeginRecordReply, error)
	EndRecord(ctx context.Context, in *EndRecordRequest, opts ...grpc.CallOption) (*EndRecordReply, error)
}

type recordClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordClient(cc grpc.ClientConnInterface) RecordClient {
	return &recordClient{cc}
}

func (c *recordClient) BeginRecord(ctx context.Context, in *BeginRecordRequest, opts ...grpc.CallOption) (*BeginRecordReply, error) {
	out := new(BeginRecordReply)
	err := c.cc.Invoke(ctx, Record_BeginRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) EndRecord(ctx context.Context, in *EndRecordRequest, opts ...grpc.CallOption) (*EndRecordReply, error) {
	out := new(EndRecordReply)
	err := c.cc.Invoke(ctx, Record_EndRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServer is the server API for Record service.
// All implementations must embed UnimplementedRecordServer
// for forward compatibility
type RecordServer interface {
	BeginRecord(context.Context, *BeginRecordRequest) (*BeginRecordReply, error)
	EndRecord(context.Context, *EndRecordRequest) (*EndRecordReply, error)
	mustEmbedUnimplementedRecordServer()
}

// UnimplementedRecordServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServer struct {
}

func (UnimplementedRecordServer) BeginRecord(context.Context, *BeginRecordRequest) (*BeginRecordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginRecord not implemented")
}
func (UnimplementedRecordServer) EndRecord(context.Context, *EndRecordRequest) (*EndRecordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndRecord not implemented")
}
func (UnimplementedRecordServer) mustEmbedUnimplementedRecordServer() {}

// UnsafeRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServer will
// result in compilation errors.
type UnsafeRecordServer interface {
	mustEmbedUnimplementedRecordServer()
}

func RegisterRecordServer(s grpc.ServiceRegistrar, srv RecordServer) {
	s.RegisterService(&Record_ServiceDesc, srv)
}

func _Record_BeginRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).BeginRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Record_BeginRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).BeginRecord(ctx, req.(*BeginRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_EndRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).EndRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Record_EndRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).EndRecord(ctx, req.(*EndRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Record_ServiceDesc is the grpc.ServiceDesc for Record service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Record_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.slacker.v1.Record",
	HandlerType: (*RecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeginRecord",
			Handler:    _Record_BeginRecord_Handler,
		},
		{
			MethodName: "EndRecord",
			Handler:    _Record_EndRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/slacker/v1/record.proto",
}