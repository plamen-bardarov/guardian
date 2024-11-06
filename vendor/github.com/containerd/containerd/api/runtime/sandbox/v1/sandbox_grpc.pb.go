//go:build !no_grpc

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/runtime/sandbox/v1/sandbox.proto

package sandbox

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

// SandboxClient is the client API for Sandbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SandboxClient interface {
	// CreateSandbox will be called right after sandbox shim instance launched.
	// It is a good place to initialize sandbox environment.
	CreateSandbox(ctx context.Context, in *CreateSandboxRequest, opts ...grpc.CallOption) (*CreateSandboxResponse, error)
	// StartSandbox will start a previously created sandbox.
	StartSandbox(ctx context.Context, in *StartSandboxRequest, opts ...grpc.CallOption) (*StartSandboxResponse, error)
	// Platform queries the platform the sandbox is going to run containers on.
	// containerd will use this to generate a proper OCI spec.
	Platform(ctx context.Context, in *PlatformRequest, opts ...grpc.CallOption) (*PlatformResponse, error)
	// StopSandbox will stop existing sandbox instance
	StopSandbox(ctx context.Context, in *StopSandboxRequest, opts ...grpc.CallOption) (*StopSandboxResponse, error)
	// WaitSandbox blocks until sandbox exits.
	WaitSandbox(ctx context.Context, in *WaitSandboxRequest, opts ...grpc.CallOption) (*WaitSandboxResponse, error)
	// SandboxStatus will return current status of the running sandbox instance
	SandboxStatus(ctx context.Context, in *SandboxStatusRequest, opts ...grpc.CallOption) (*SandboxStatusResponse, error)
	// PingSandbox is a lightweight API call to check whether sandbox alive.
	PingSandbox(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// ShutdownSandbox must shutdown shim instance.
	ShutdownSandbox(ctx context.Context, in *ShutdownSandboxRequest, opts ...grpc.CallOption) (*ShutdownSandboxResponse, error)
	// SandboxMetrics retrieves metrics about a sandbox instance.
	SandboxMetrics(ctx context.Context, in *SandboxMetricsRequest, opts ...grpc.CallOption) (*SandboxMetricsResponse, error)
}

type sandboxClient struct {
	cc grpc.ClientConnInterface
}

func NewSandboxClient(cc grpc.ClientConnInterface) SandboxClient {
	return &sandboxClient{cc}
}

func (c *sandboxClient) CreateSandbox(ctx context.Context, in *CreateSandboxRequest, opts ...grpc.CallOption) (*CreateSandboxResponse, error) {
	out := new(CreateSandboxResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/CreateSandbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) StartSandbox(ctx context.Context, in *StartSandboxRequest, opts ...grpc.CallOption) (*StartSandboxResponse, error) {
	out := new(StartSandboxResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/StartSandbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) Platform(ctx context.Context, in *PlatformRequest, opts ...grpc.CallOption) (*PlatformResponse, error) {
	out := new(PlatformResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/Platform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) StopSandbox(ctx context.Context, in *StopSandboxRequest, opts ...grpc.CallOption) (*StopSandboxResponse, error) {
	out := new(StopSandboxResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/StopSandbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) WaitSandbox(ctx context.Context, in *WaitSandboxRequest, opts ...grpc.CallOption) (*WaitSandboxResponse, error) {
	out := new(WaitSandboxResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/WaitSandbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) SandboxStatus(ctx context.Context, in *SandboxStatusRequest, opts ...grpc.CallOption) (*SandboxStatusResponse, error) {
	out := new(SandboxStatusResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/SandboxStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) PingSandbox(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/PingSandbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) ShutdownSandbox(ctx context.Context, in *ShutdownSandboxRequest, opts ...grpc.CallOption) (*ShutdownSandboxResponse, error) {
	out := new(ShutdownSandboxResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/ShutdownSandbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sandboxClient) SandboxMetrics(ctx context.Context, in *SandboxMetricsRequest, opts ...grpc.CallOption) (*SandboxMetricsResponse, error) {
	out := new(SandboxMetricsResponse)
	err := c.cc.Invoke(ctx, "/containerd.runtime.sandbox.v1.Sandbox/SandboxMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SandboxServer is the server API for Sandbox service.
// All implementations must embed UnimplementedSandboxServer
// for forward compatibility
type SandboxServer interface {
	// CreateSandbox will be called right after sandbox shim instance launched.
	// It is a good place to initialize sandbox environment.
	CreateSandbox(context.Context, *CreateSandboxRequest) (*CreateSandboxResponse, error)
	// StartSandbox will start a previously created sandbox.
	StartSandbox(context.Context, *StartSandboxRequest) (*StartSandboxResponse, error)
	// Platform queries the platform the sandbox is going to run containers on.
	// containerd will use this to generate a proper OCI spec.
	Platform(context.Context, *PlatformRequest) (*PlatformResponse, error)
	// StopSandbox will stop existing sandbox instance
	StopSandbox(context.Context, *StopSandboxRequest) (*StopSandboxResponse, error)
	// WaitSandbox blocks until sandbox exits.
	WaitSandbox(context.Context, *WaitSandboxRequest) (*WaitSandboxResponse, error)
	// SandboxStatus will return current status of the running sandbox instance
	SandboxStatus(context.Context, *SandboxStatusRequest) (*SandboxStatusResponse, error)
	// PingSandbox is a lightweight API call to check whether sandbox alive.
	PingSandbox(context.Context, *PingRequest) (*PingResponse, error)
	// ShutdownSandbox must shutdown shim instance.
	ShutdownSandbox(context.Context, *ShutdownSandboxRequest) (*ShutdownSandboxResponse, error)
	// SandboxMetrics retrieves metrics about a sandbox instance.
	SandboxMetrics(context.Context, *SandboxMetricsRequest) (*SandboxMetricsResponse, error)
	mustEmbedUnimplementedSandboxServer()
}

// UnimplementedSandboxServer must be embedded to have forward compatible implementations.
type UnimplementedSandboxServer struct {
}

func (UnimplementedSandboxServer) CreateSandbox(context.Context, *CreateSandboxRequest) (*CreateSandboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSandbox not implemented")
}
func (UnimplementedSandboxServer) StartSandbox(context.Context, *StartSandboxRequest) (*StartSandboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSandbox not implemented")
}
func (UnimplementedSandboxServer) Platform(context.Context, *PlatformRequest) (*PlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Platform not implemented")
}
func (UnimplementedSandboxServer) StopSandbox(context.Context, *StopSandboxRequest) (*StopSandboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopSandbox not implemented")
}
func (UnimplementedSandboxServer) WaitSandbox(context.Context, *WaitSandboxRequest) (*WaitSandboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitSandbox not implemented")
}
func (UnimplementedSandboxServer) SandboxStatus(context.Context, *SandboxStatusRequest) (*SandboxStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SandboxStatus not implemented")
}
func (UnimplementedSandboxServer) PingSandbox(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingSandbox not implemented")
}
func (UnimplementedSandboxServer) ShutdownSandbox(context.Context, *ShutdownSandboxRequest) (*ShutdownSandboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutdownSandbox not implemented")
}
func (UnimplementedSandboxServer) SandboxMetrics(context.Context, *SandboxMetricsRequest) (*SandboxMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SandboxMetrics not implemented")
}
func (UnimplementedSandboxServer) mustEmbedUnimplementedSandboxServer() {}

// UnsafeSandboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SandboxServer will
// result in compilation errors.
type UnsafeSandboxServer interface {
	mustEmbedUnimplementedSandboxServer()
}

func RegisterSandboxServer(s grpc.ServiceRegistrar, srv SandboxServer) {
	s.RegisterService(&Sandbox_ServiceDesc, srv)
}

func _Sandbox_CreateSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSandboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).CreateSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/CreateSandbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).CreateSandbox(ctx, req.(*CreateSandboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_StartSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSandboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).StartSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/StartSandbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).StartSandbox(ctx, req.(*StartSandboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_Platform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).Platform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/Platform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).Platform(ctx, req.(*PlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_StopSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopSandboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).StopSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/StopSandbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).StopSandbox(ctx, req.(*StopSandboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_WaitSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitSandboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).WaitSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/WaitSandbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).WaitSandbox(ctx, req.(*WaitSandboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_SandboxStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).SandboxStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/SandboxStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).SandboxStatus(ctx, req.(*SandboxStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_PingSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).PingSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/PingSandbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).PingSandbox(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_ShutdownSandbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownSandboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).ShutdownSandbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/ShutdownSandbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).ShutdownSandbox(ctx, req.(*ShutdownSandboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sandbox_SandboxMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandboxMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SandboxServer).SandboxMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.runtime.sandbox.v1.Sandbox/SandboxMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SandboxServer).SandboxMetrics(ctx, req.(*SandboxMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sandbox_ServiceDesc is the grpc.ServiceDesc for Sandbox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sandbox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.runtime.sandbox.v1.Sandbox",
	HandlerType: (*SandboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSandbox",
			Handler:    _Sandbox_CreateSandbox_Handler,
		},
		{
			MethodName: "StartSandbox",
			Handler:    _Sandbox_StartSandbox_Handler,
		},
		{
			MethodName: "Platform",
			Handler:    _Sandbox_Platform_Handler,
		},
		{
			MethodName: "StopSandbox",
			Handler:    _Sandbox_StopSandbox_Handler,
		},
		{
			MethodName: "WaitSandbox",
			Handler:    _Sandbox_WaitSandbox_Handler,
		},
		{
			MethodName: "SandboxStatus",
			Handler:    _Sandbox_SandboxStatus_Handler,
		},
		{
			MethodName: "PingSandbox",
			Handler:    _Sandbox_PingSandbox_Handler,
		},
		{
			MethodName: "ShutdownSandbox",
			Handler:    _Sandbox_ShutdownSandbox_Handler,
		},
		{
			MethodName: "SandboxMetrics",
			Handler:    _Sandbox_SandboxMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/runtime/sandbox/v1/sandbox.proto",
}
