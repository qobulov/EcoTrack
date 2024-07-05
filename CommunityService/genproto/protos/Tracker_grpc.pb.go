// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: protos/Tracker.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	HabitTrackerService_CreateHabit_FullMethodName         = "/habittracker.HabitTrackerService/CreateHabit"
	HabitTrackerService_GetHabits_FullMethodName           = "/habittracker.HabitTrackerService/GetHabits"
	HabitTrackerService_LogHabit_FullMethodName            = "/habittracker.HabitTrackerService/LogHabit"
	HabitTrackerService_GetHabitLogs_FullMethodName        = "/habittracker.HabitTrackerService/GetHabitLogs"
	HabitTrackerService_GetHabitSuggestions_FullMethodName = "/habittracker.HabitTrackerService/GetHabitSuggestions"
	HabitTrackerService_GetUserHabits_FullMethodName       = "/habittracker.HabitTrackerService/GetUserHabits"
	HabitTrackerService_UpdateHabit_FullMethodName         = "/habittracker.HabitTrackerService/UpdateHabit"
	HabitTrackerService_DeleteHabit_FullMethodName         = "/habittracker.HabitTrackerService/DeleteHabit"
)

// HabitTrackerServiceClient is the client API for HabitTrackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HabitTrackerServiceClient interface {
	CreateHabit(ctx context.Context, in *CreateHabitRequest, opts ...grpc.CallOption) (*CreateHabitResponse, error)
	GetHabits(ctx context.Context, in *GetHabitsRequest, opts ...grpc.CallOption) (*GetHabitsResponse, error)
	LogHabit(ctx context.Context, in *LogHabitRequest, opts ...grpc.CallOption) (*LogHabitResponse, error)
	GetHabitLogs(ctx context.Context, in *GetHabitLogsRequest, opts ...grpc.CallOption) (*GetHabitLogsResponse, error)
	GetHabitSuggestions(ctx context.Context, in *GetHabitSuggestionsRequest, opts ...grpc.CallOption) (*GetHabitSuggestionsResponse, error)
	GetUserHabits(ctx context.Context, in *GetUserHabitsRequest, opts ...grpc.CallOption) (*GetUserHabitsResponse, error)
	UpdateHabit(ctx context.Context, in *UpdateHabitRequest, opts ...grpc.CallOption) (*UpdateHabitResponse, error)
	DeleteHabit(ctx context.Context, in *DeleteHabitRequest, opts ...grpc.CallOption) (*DeleteHabitResponse, error)
}

type habitTrackerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHabitTrackerServiceClient(cc grpc.ClientConnInterface) HabitTrackerServiceClient {
	return &habitTrackerServiceClient{cc}
}

func (c *habitTrackerServiceClient) CreateHabit(ctx context.Context, in *CreateHabitRequest, opts ...grpc.CallOption) (*CreateHabitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateHabitResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_CreateHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerServiceClient) GetHabits(ctx context.Context, in *GetHabitsRequest, opts ...grpc.CallOption) (*GetHabitsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHabitsResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_GetHabits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerServiceClient) LogHabit(ctx context.Context, in *LogHabitRequest, opts ...grpc.CallOption) (*LogHabitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogHabitResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_LogHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerServiceClient) GetHabitLogs(ctx context.Context, in *GetHabitLogsRequest, opts ...grpc.CallOption) (*GetHabitLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHabitLogsResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_GetHabitLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerServiceClient) GetHabitSuggestions(ctx context.Context, in *GetHabitSuggestionsRequest, opts ...grpc.CallOption) (*GetHabitSuggestionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHabitSuggestionsResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_GetHabitSuggestions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerServiceClient) GetUserHabits(ctx context.Context, in *GetUserHabitsRequest, opts ...grpc.CallOption) (*GetUserHabitsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserHabitsResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_GetUserHabits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerServiceClient) UpdateHabit(ctx context.Context, in *UpdateHabitRequest, opts ...grpc.CallOption) (*UpdateHabitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateHabitResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_UpdateHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *habitTrackerServiceClient) DeleteHabit(ctx context.Context, in *DeleteHabitRequest, opts ...grpc.CallOption) (*DeleteHabitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteHabitResponse)
	err := c.cc.Invoke(ctx, HabitTrackerService_DeleteHabit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HabitTrackerServiceServer is the server API for HabitTrackerService service.
// All implementations must embed UnimplementedHabitTrackerServiceServer
// for forward compatibility
type HabitTrackerServiceServer interface {
	CreateHabit(context.Context, *CreateHabitRequest) (*CreateHabitResponse, error)
	GetHabits(context.Context, *GetHabitsRequest) (*GetHabitsResponse, error)
	LogHabit(context.Context, *LogHabitRequest) (*LogHabitResponse, error)
	GetHabitLogs(context.Context, *GetHabitLogsRequest) (*GetHabitLogsResponse, error)
	GetHabitSuggestions(context.Context, *GetHabitSuggestionsRequest) (*GetHabitSuggestionsResponse, error)
	GetUserHabits(context.Context, *GetUserHabitsRequest) (*GetUserHabitsResponse, error)
	UpdateHabit(context.Context, *UpdateHabitRequest) (*UpdateHabitResponse, error)
	DeleteHabit(context.Context, *DeleteHabitRequest) (*DeleteHabitResponse, error)
	mustEmbedUnimplementedHabitTrackerServiceServer()
}

// UnimplementedHabitTrackerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHabitTrackerServiceServer struct {
}

func (UnimplementedHabitTrackerServiceServer) CreateHabit(context.Context, *CreateHabitRequest) (*CreateHabitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHabit not implemented")
}
func (UnimplementedHabitTrackerServiceServer) GetHabits(context.Context, *GetHabitsRequest) (*GetHabitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHabits not implemented")
}
func (UnimplementedHabitTrackerServiceServer) LogHabit(context.Context, *LogHabitRequest) (*LogHabitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogHabit not implemented")
}
func (UnimplementedHabitTrackerServiceServer) GetHabitLogs(context.Context, *GetHabitLogsRequest) (*GetHabitLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHabitLogs not implemented")
}
func (UnimplementedHabitTrackerServiceServer) GetHabitSuggestions(context.Context, *GetHabitSuggestionsRequest) (*GetHabitSuggestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHabitSuggestions not implemented")
}
func (UnimplementedHabitTrackerServiceServer) GetUserHabits(context.Context, *GetUserHabitsRequest) (*GetUserHabitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserHabits not implemented")
}
func (UnimplementedHabitTrackerServiceServer) UpdateHabit(context.Context, *UpdateHabitRequest) (*UpdateHabitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHabit not implemented")
}
func (UnimplementedHabitTrackerServiceServer) DeleteHabit(context.Context, *DeleteHabitRequest) (*DeleteHabitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHabit not implemented")
}
func (UnimplementedHabitTrackerServiceServer) mustEmbedUnimplementedHabitTrackerServiceServer() {}

// UnsafeHabitTrackerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HabitTrackerServiceServer will
// result in compilation errors.
type UnsafeHabitTrackerServiceServer interface {
	mustEmbedUnimplementedHabitTrackerServiceServer()
}

func RegisterHabitTrackerServiceServer(s grpc.ServiceRegistrar, srv HabitTrackerServiceServer) {
	s.RegisterService(&HabitTrackerService_ServiceDesc, srv)
}

func _HabitTrackerService_CreateHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHabitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).CreateHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_CreateHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).CreateHabit(ctx, req.(*CreateHabitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTrackerService_GetHabits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHabitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).GetHabits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_GetHabits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).GetHabits(ctx, req.(*GetHabitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTrackerService_LogHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogHabitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).LogHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_LogHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).LogHabit(ctx, req.(*LogHabitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTrackerService_GetHabitLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHabitLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).GetHabitLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_GetHabitLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).GetHabitLogs(ctx, req.(*GetHabitLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTrackerService_GetHabitSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHabitSuggestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).GetHabitSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_GetHabitSuggestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).GetHabitSuggestions(ctx, req.(*GetHabitSuggestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTrackerService_GetUserHabits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserHabitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).GetUserHabits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_GetUserHabits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).GetUserHabits(ctx, req.(*GetUserHabitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTrackerService_UpdateHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHabitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).UpdateHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_UpdateHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).UpdateHabit(ctx, req.(*UpdateHabitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HabitTrackerService_DeleteHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHabitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitTrackerServiceServer).DeleteHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HabitTrackerService_DeleteHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitTrackerServiceServer).DeleteHabit(ctx, req.(*DeleteHabitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HabitTrackerService_ServiceDesc is the grpc.ServiceDesc for HabitTrackerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HabitTrackerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "habittracker.HabitTrackerService",
	HandlerType: (*HabitTrackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHabit",
			Handler:    _HabitTrackerService_CreateHabit_Handler,
		},
		{
			MethodName: "GetHabits",
			Handler:    _HabitTrackerService_GetHabits_Handler,
		},
		{
			MethodName: "LogHabit",
			Handler:    _HabitTrackerService_LogHabit_Handler,
		},
		{
			MethodName: "GetHabitLogs",
			Handler:    _HabitTrackerService_GetHabitLogs_Handler,
		},
		{
			MethodName: "GetHabitSuggestions",
			Handler:    _HabitTrackerService_GetHabitSuggestions_Handler,
		},
		{
			MethodName: "GetUserHabits",
			Handler:    _HabitTrackerService_GetUserHabits_Handler,
		},
		{
			MethodName: "UpdateHabit",
			Handler:    _HabitTrackerService_UpdateHabit_Handler,
		},
		{
			MethodName: "DeleteHabit",
			Handler:    _HabitTrackerService_DeleteHabit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/Tracker.proto",
}