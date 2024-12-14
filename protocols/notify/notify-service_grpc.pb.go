// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: notify-service.proto

package notify

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NotificationService_GetAllAppliedCandidatesByNoty_FullMethodName = "/service.NotificationService/GetAllAppliedCandidatesByNoty"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ResumeService is
type NotificationServiceClient interface {
	GetAllAppliedCandidatesByNoty(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*Notifications, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetAllAppliedCandidatesByNoty(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*Notifications, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Notifications)
	err := c.cc.Invoke(ctx, NotificationService_GetAllAppliedCandidatesByNoty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility.
//
// ResumeService is
type NotificationServiceServer interface {
	GetAllAppliedCandidatesByNoty(context.Context, *NotifyReq) (*Notifications, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotificationServiceServer struct{}

func (UnimplementedNotificationServiceServer) GetAllAppliedCandidatesByNoty(context.Context, *NotifyReq) (*Notifications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAppliedCandidatesByNoty not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}
func (UnimplementedNotificationServiceServer) testEmbeddedByValue()                             {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	// If the following call pancis, it indicates UnimplementedNotificationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_GetAllAppliedCandidatesByNoty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllAppliedCandidatesByNoty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetAllAppliedCandidatesByNoty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllAppliedCandidatesByNoty(ctx, req.(*NotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAppliedCandidatesByNoty",
			Handler:    _NotificationService_GetAllAppliedCandidatesByNoty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notify-service.proto",
}
