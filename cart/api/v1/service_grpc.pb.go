// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: api/v1/service.proto

package cart

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Cart_List_FullMethodName       = "/route256.cart.api.v1.Cart/List"
	Cart_Clear_FullMethodName      = "/route256.cart.api.v1.Cart/Clear"
	Cart_ItemAdd_FullMethodName    = "/route256.cart.api.v1.Cart/ItemAdd"
	Cart_ItemDelete_FullMethodName = "/route256.cart.api.v1.Cart/ItemDelete"
	Cart_Checkout_FullMethodName   = "/route256.cart.api.v1.Cart/Checkout"
)

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ItemAdd(ctx context.Context, in *ItemAddRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ItemDelete(ctx context.Context, in *ItemDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Cart_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Cart_Clear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) ItemAdd(ctx context.Context, in *ItemAddRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Cart_ItemAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) ItemDelete(ctx context.Context, in *ItemDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Cart_ItemDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error) {
	out := new(CheckoutResponse)
	err := c.cc.Invoke(ctx, Cart_Checkout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Clear(context.Context, *ClearRequest) (*empty.Empty, error)
	ItemAdd(context.Context, *ItemAddRequest) (*empty.Empty, error)
	ItemDelete(context.Context, *ItemDeleteRequest) (*empty.Empty, error)
	Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCartServer) Clear(context.Context, *ClearRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedCartServer) ItemAdd(context.Context, *ItemAddRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ItemAdd not implemented")
}
func (UnimplementedCartServer) ItemDelete(context.Context, *ItemDeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ItemDelete not implemented")
}
func (UnimplementedCartServer) Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_Clear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_ItemAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).ItemAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_ItemAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).ItemAdd(ctx, req.(*ItemAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_ItemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).ItemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_ItemDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).ItemDelete(ctx, req.(*ItemDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_Checkout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route256.cart.api.v1.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Cart_List_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _Cart_Clear_Handler,
		},
		{
			MethodName: "ItemAdd",
			Handler:    _Cart_ItemAdd_Handler,
		},
		{
			MethodName: "ItemDelete",
			Handler:    _Cart_ItemDelete_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _Cart_Checkout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/service.proto",
}
