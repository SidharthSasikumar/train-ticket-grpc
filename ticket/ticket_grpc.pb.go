// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: ticket.proto

package ticket

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
	TicketingService_PurchaseTicket_FullMethodName = "/TicketingService/PurchaseTicket"
	TicketingService_GetReceipt_FullMethodName     = "/TicketingService/GetReceipt"
	TicketingService_ViewUsers_FullMethodName      = "/TicketingService/ViewUsers"
	TicketingService_RemoveUser_FullMethodName     = "/TicketingService/RemoveUser"
	TicketingService_ModifySeat_FullMethodName     = "/TicketingService/ModifySeat"
)

// TicketingServiceClient is the client API for TicketingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketingServiceClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
	GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error)
	ViewUsers(ctx context.Context, in *ViewUsersRequest, opts ...grpc.CallOption) (*ViewUsersResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error)
}

type ticketingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketingServiceClient(cc grpc.ClientConnInterface) TicketingServiceClient {
	return &ticketingServiceClient{cc}
}

func (c *ticketingServiceClient) PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, TicketingService_PurchaseTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingServiceClient) GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error) {
	out := new(GetReceiptResponse)
	err := c.cc.Invoke(ctx, TicketingService_GetReceipt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingServiceClient) ViewUsers(ctx context.Context, in *ViewUsersRequest, opts ...grpc.CallOption) (*ViewUsersResponse, error) {
	out := new(ViewUsersResponse)
	err := c.cc.Invoke(ctx, TicketingService_ViewUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, TicketingService_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error) {
	out := new(ModifySeatResponse)
	err := c.cc.Invoke(ctx, TicketingService_ModifySeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketingServiceServer is the server API for TicketingService service.
// All implementations must embed UnimplementedTicketingServiceServer
// for forward compatibility
type TicketingServiceServer interface {
	PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
	GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error)
	ViewUsers(context.Context, *ViewUsersRequest) (*ViewUsersResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error)
	mustEmbedUnimplementedTicketingServiceServer()
}

// UnimplementedTicketingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketingServiceServer struct {
}

func (UnimplementedTicketingServiceServer) PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTicketingServiceServer) GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTicketingServiceServer) ViewUsers(context.Context, *ViewUsersRequest) (*ViewUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewUsers not implemented")
}
func (UnimplementedTicketingServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTicketingServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTicketingServiceServer) mustEmbedUnimplementedTicketingServiceServer() {}

// UnsafeTicketingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketingServiceServer will
// result in compilation errors.
type UnsafeTicketingServiceServer interface {
	mustEmbedUnimplementedTicketingServiceServer()
}

func RegisterTicketingServiceServer(s grpc.ServiceRegistrar, srv TicketingServiceServer) {
	s.RegisterService(&TicketingService_ServiceDesc, srv)
}

func _TicketingService_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServiceServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingService_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServiceServer).PurchaseTicket(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingService_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServiceServer).GetReceipt(ctx, req.(*GetReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingService_ViewUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServiceServer).ViewUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingService_ViewUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServiceServer).ViewUsers(ctx, req.(*ViewUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingService_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketingService_ServiceDesc is the grpc.ServiceDesc for TicketingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TicketingService",
	HandlerType: (*TicketingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TicketingService_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TicketingService_GetReceipt_Handler,
		},
		{
			MethodName: "ViewUsers",
			Handler:    _TicketingService_ViewUsers_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TicketingService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TicketingService_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket.proto",
}
