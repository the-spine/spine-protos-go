// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: hospital.proto

package hospital

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
	HospitalService_RegisterHospital_FullMethodName = "/hospital.HospitalService/RegisterHospital"
)

// HospitalServiceClient is the client API for HospitalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalServiceClient interface {
	RegisterHospital(ctx context.Context, in *RegisterHospitalRequest, opts ...grpc.CallOption) (*RegisterHospitalResponse, error)
}

type hospitalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalServiceClient(cc grpc.ClientConnInterface) HospitalServiceClient {
	return &hospitalServiceClient{cc}
}

func (c *hospitalServiceClient) RegisterHospital(ctx context.Context, in *RegisterHospitalRequest, opts ...grpc.CallOption) (*RegisterHospitalResponse, error) {
	out := new(RegisterHospitalResponse)
	err := c.cc.Invoke(ctx, HospitalService_RegisterHospital_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServiceServer is the server API for HospitalService service.
// All implementations must embed UnimplementedHospitalServiceServer
// for forward compatibility
type HospitalServiceServer interface {
	RegisterHospital(context.Context, *RegisterHospitalRequest) (*RegisterHospitalResponse, error)
	mustEmbedUnimplementedHospitalServiceServer()
}

// UnimplementedHospitalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHospitalServiceServer struct {
}

func (UnimplementedHospitalServiceServer) RegisterHospital(context.Context, *RegisterHospitalRequest) (*RegisterHospitalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterHospital not implemented")
}
func (UnimplementedHospitalServiceServer) mustEmbedUnimplementedHospitalServiceServer() {}

// UnsafeHospitalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HospitalServiceServer will
// result in compilation errors.
type UnsafeHospitalServiceServer interface {
	mustEmbedUnimplementedHospitalServiceServer()
}

func RegisterHospitalServiceServer(s grpc.ServiceRegistrar, srv HospitalServiceServer) {
	s.RegisterService(&HospitalService_ServiceDesc, srv)
}

func _HospitalService_RegisterHospital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterHospitalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).RegisterHospital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HospitalService_RegisterHospital_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).RegisterHospital(ctx, req.(*RegisterHospitalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HospitalService_ServiceDesc is the grpc.ServiceDesc for HospitalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HospitalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hospital.HospitalService",
	HandlerType: (*HospitalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterHospital",
			Handler:    _HospitalService_RegisterHospital_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hospital.proto",
}
