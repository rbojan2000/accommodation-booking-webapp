// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: accommodation_service.proto

package accommodation

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
	AccommodationService_GetAll_FullMethodName              = "/accommodation.AccommodationService/GetAll"
	AccommodationService_GetById_FullMethodName             = "/accommodation.AccommodationService/GetById"
	AccommodationService_Search_FullMethodName              = "/accommodation.AccommodationService/Search"
	AccommodationService_CreateAccommodation_FullMethodName = "/accommodation.AccommodationService/CreateAccommodation"
	AccommodationService_CreateAppointment_FullMethodName   = "/accommodation.AccommodationService/CreateAppointment"
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	GetAll(ctx context.Context, in *GetAllAccommodationsRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error)
	GetById(ctx context.Context, in *GetAccommodationByIdRequest, opts ...grpc.CallOption) (*GetAccommodationByIdResponse, error)
	Search(ctx context.Context, in *GetAccommodationsByParamsRequest, opts ...grpc.CallOption) (*GetAccommodationsByParamsResponse, error)
	CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error)
	CreateAppointment(ctx context.Context, in *CreateAppointmentRequest, opts ...grpc.CallOption) (*CreateAppointmentResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) GetAll(ctx context.Context, in *GetAllAccommodationsRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error) {
	out := new(GetAllAccommodationsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetById(ctx context.Context, in *GetAccommodationByIdRequest, opts ...grpc.CallOption) (*GetAccommodationByIdResponse, error) {
	out := new(GetAccommodationByIdResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Search(ctx context.Context, in *GetAccommodationsByParamsRequest, opts ...grpc.CallOption) (*GetAccommodationsByParamsResponse, error) {
	out := new(GetAccommodationsByParamsResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error) {
	out := new(CreateAccommodationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_CreateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAppointment(ctx context.Context, in *CreateAppointmentRequest, opts ...grpc.CallOption) (*CreateAppointmentResponse, error) {
	out := new(CreateAppointmentResponse)
	err := c.cc.Invoke(ctx, AccommodationService_CreateAppointment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	GetAll(context.Context, *GetAllAccommodationsRequest) (*GetAllAccommodationsResponse, error)
	GetById(context.Context, *GetAccommodationByIdRequest) (*GetAccommodationByIdResponse, error)
	Search(context.Context, *GetAccommodationsByParamsRequest) (*GetAccommodationsByParamsResponse, error)
	CreateAccommodation(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error)
	CreateAppointment(context.Context, *CreateAppointmentRequest) (*CreateAppointmentResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) GetAll(context.Context, *GetAllAccommodationsRequest) (*GetAllAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAccommodationServiceServer) GetById(context.Context, *GetAccommodationByIdRequest) (*GetAccommodationByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedAccommodationServiceServer) Search(context.Context, *GetAccommodationsByParamsRequest) (*GetAccommodationsByParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAccommodation(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAppointment(context.Context, *CreateAppointmentRequest) (*CreateAppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppointment not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAll(ctx, req.(*GetAllAccommodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetById(ctx, req.(*GetAccommodationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationsByParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Search(ctx, req.(*GetAccommodationsByParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, req.(*CreateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreateAppointment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAppointment(ctx, req.(*CreateAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _AccommodationService_GetAll_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _AccommodationService_GetById_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _AccommodationService_Search_Handler,
		},
		{
			MethodName: "CreateAccommodation",
			Handler:    _AccommodationService_CreateAccommodation_Handler,
		},
		{
			MethodName: "CreateAppointment",
			Handler:    _AccommodationService_CreateAppointment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service.proto",
}
