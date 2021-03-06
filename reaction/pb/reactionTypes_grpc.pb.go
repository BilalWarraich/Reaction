// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// ReactionServiceClient is the client API for ReactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReactionServiceClient interface {
	PostReactionTypes(ctx context.Context, in *PostReactionTypesRequest, opts ...grpc.CallOption) (*PostReactionTypesResponse, error)
	GetReactionTypeID(ctx context.Context, in *GetReactionTypeIDRequest, opts ...grpc.CallOption) (*GetReactionTypeIDResponse, error)
	GetReactionTypes(ctx context.Context, in *GetReactionTypesRequest, opts ...grpc.CallOption) (*GetReactionTypesResponse, error)
	PostReactions(ctx context.Context, in *PostReactionsRequest, opts ...grpc.CallOption) (*PostReactionsResponse, error)
	GetUserByPostID(ctx context.Context, in *GetUserByPostIDRequest, opts ...grpc.CallOption) (*GetUserByPostIDResponse, error)
	GetReactByUserID(ctx context.Context, in *GetReactByUserIDRequest, opts ...grpc.CallOption) (*GetReactByUserIDResponse, error)
	DeleteReactionByPostID(ctx context.Context, in *DeleteReactionByPostIDRequest, opts ...grpc.CallOption) (*DeleteReactionByPostIDResponse, error)
	DeleteReactionByUserPostID(ctx context.Context, in *DeleteReactionByUserPostIDRequest, opts ...grpc.CallOption) (*DeleteReactionByUserPostIDResponse, error)
	UpdateReactions(ctx context.Context, in *UpdateReactionsRequest, opts ...grpc.CallOption) (*UpdateReactionsResponse, error)
	TotalReactionCount(ctx context.Context, in *TotalReactionCountRequest, opts ...grpc.CallOption) (*TotalReactionCountResponse, error)
	ReactionCountByType(ctx context.Context, in *ReactionCountByTypeRequest, opts ...grpc.CallOption) (*ReactionCountByTypeResponse, error)
}

type reactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReactionServiceClient(cc grpc.ClientConnInterface) ReactionServiceClient {
	return &reactionServiceClient{cc}
}

func (c *reactionServiceClient) PostReactionTypes(ctx context.Context, in *PostReactionTypesRequest, opts ...grpc.CallOption) (*PostReactionTypesResponse, error) {
	out := new(PostReactionTypesResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/PostReactionTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) GetReactionTypeID(ctx context.Context, in *GetReactionTypeIDRequest, opts ...grpc.CallOption) (*GetReactionTypeIDResponse, error) {
	out := new(GetReactionTypeIDResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/GetReactionTypeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) GetReactionTypes(ctx context.Context, in *GetReactionTypesRequest, opts ...grpc.CallOption) (*GetReactionTypesResponse, error) {
	out := new(GetReactionTypesResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/GetReactionTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) PostReactions(ctx context.Context, in *PostReactionsRequest, opts ...grpc.CallOption) (*PostReactionsResponse, error) {
	out := new(PostReactionsResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/PostReactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) GetUserByPostID(ctx context.Context, in *GetUserByPostIDRequest, opts ...grpc.CallOption) (*GetUserByPostIDResponse, error) {
	out := new(GetUserByPostIDResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/GetUserByPostID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) GetReactByUserID(ctx context.Context, in *GetReactByUserIDRequest, opts ...grpc.CallOption) (*GetReactByUserIDResponse, error) {
	out := new(GetReactByUserIDResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/GetReactByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) DeleteReactionByPostID(ctx context.Context, in *DeleteReactionByPostIDRequest, opts ...grpc.CallOption) (*DeleteReactionByPostIDResponse, error) {
	out := new(DeleteReactionByPostIDResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/DeleteReactionByPostID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) DeleteReactionByUserPostID(ctx context.Context, in *DeleteReactionByUserPostIDRequest, opts ...grpc.CallOption) (*DeleteReactionByUserPostIDResponse, error) {
	out := new(DeleteReactionByUserPostIDResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/DeleteReactionByUserPostID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) UpdateReactions(ctx context.Context, in *UpdateReactionsRequest, opts ...grpc.CallOption) (*UpdateReactionsResponse, error) {
	out := new(UpdateReactionsResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/UpdateReactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) TotalReactionCount(ctx context.Context, in *TotalReactionCountRequest, opts ...grpc.CallOption) (*TotalReactionCountResponse, error) {
	out := new(TotalReactionCountResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/TotalReactionCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactionServiceClient) ReactionCountByType(ctx context.Context, in *ReactionCountByTypeRequest, opts ...grpc.CallOption) (*ReactionCountByTypeResponse, error) {
	out := new(ReactionCountByTypeResponse)
	err := c.cc.Invoke(ctx, "/pb.ReactionService/ReactionCountByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReactionServiceServer is the server API for ReactionService service.
// All implementations must embed UnimplementedReactionServiceServer
// for forward compatibility
type ReactionServiceServer interface {
	PostReactionTypes(context.Context, *PostReactionTypesRequest) (*PostReactionTypesResponse, error)
	GetReactionTypeID(context.Context, *GetReactionTypeIDRequest) (*GetReactionTypeIDResponse, error)
	GetReactionTypes(context.Context, *GetReactionTypesRequest) (*GetReactionTypesResponse, error)
	PostReactions(context.Context, *PostReactionsRequest) (*PostReactionsResponse, error)
	GetUserByPostID(context.Context, *GetUserByPostIDRequest) (*GetUserByPostIDResponse, error)
	GetReactByUserID(context.Context, *GetReactByUserIDRequest) (*GetReactByUserIDResponse, error)
	DeleteReactionByPostID(context.Context, *DeleteReactionByPostIDRequest) (*DeleteReactionByPostIDResponse, error)
	DeleteReactionByUserPostID(context.Context, *DeleteReactionByUserPostIDRequest) (*DeleteReactionByUserPostIDResponse, error)
	UpdateReactions(context.Context, *UpdateReactionsRequest) (*UpdateReactionsResponse, error)
	TotalReactionCount(context.Context, *TotalReactionCountRequest) (*TotalReactionCountResponse, error)
	ReactionCountByType(context.Context, *ReactionCountByTypeRequest) (*ReactionCountByTypeResponse, error)
	mustEmbedUnimplementedReactionServiceServer()
}

// UnimplementedReactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReactionServiceServer struct {
}

func (UnimplementedReactionServiceServer) PostReactionTypes(context.Context, *PostReactionTypesRequest) (*PostReactionTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostReactionTypes not implemented")
}
func (UnimplementedReactionServiceServer) GetReactionTypeID(context.Context, *GetReactionTypeIDRequest) (*GetReactionTypeIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReactionTypeID not implemented")
}
func (UnimplementedReactionServiceServer) GetReactionTypes(context.Context, *GetReactionTypesRequest) (*GetReactionTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReactionTypes not implemented")
}
func (UnimplementedReactionServiceServer) PostReactions(context.Context, *PostReactionsRequest) (*PostReactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostReactions not implemented")
}
func (UnimplementedReactionServiceServer) GetUserByPostID(context.Context, *GetUserByPostIDRequest) (*GetUserByPostIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByPostID not implemented")
}
func (UnimplementedReactionServiceServer) GetReactByUserID(context.Context, *GetReactByUserIDRequest) (*GetReactByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReactByUserID not implemented")
}
func (UnimplementedReactionServiceServer) DeleteReactionByPostID(context.Context, *DeleteReactionByPostIDRequest) (*DeleteReactionByPostIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReactionByPostID not implemented")
}
func (UnimplementedReactionServiceServer) DeleteReactionByUserPostID(context.Context, *DeleteReactionByUserPostIDRequest) (*DeleteReactionByUserPostIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReactionByUserPostID not implemented")
}
func (UnimplementedReactionServiceServer) UpdateReactions(context.Context, *UpdateReactionsRequest) (*UpdateReactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReactions not implemented")
}
func (UnimplementedReactionServiceServer) TotalReactionCount(context.Context, *TotalReactionCountRequest) (*TotalReactionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalReactionCount not implemented")
}
func (UnimplementedReactionServiceServer) ReactionCountByType(context.Context, *ReactionCountByTypeRequest) (*ReactionCountByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReactionCountByType not implemented")
}
func (UnimplementedReactionServiceServer) mustEmbedUnimplementedReactionServiceServer() {}

// UnsafeReactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReactionServiceServer will
// result in compilation errors.
type UnsafeReactionServiceServer interface {
	mustEmbedUnimplementedReactionServiceServer()
}

func RegisterReactionServiceServer(s grpc.ServiceRegistrar, srv ReactionServiceServer) {
	s.RegisterService(&ReactionService_ServiceDesc, srv)
}

func _ReactionService_PostReactionTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReactionTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).PostReactionTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/PostReactionTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).PostReactionTypes(ctx, req.(*PostReactionTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_GetReactionTypeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReactionTypeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).GetReactionTypeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/GetReactionTypeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).GetReactionTypeID(ctx, req.(*GetReactionTypeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_GetReactionTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReactionTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).GetReactionTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/GetReactionTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).GetReactionTypes(ctx, req.(*GetReactionTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_PostReactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).PostReactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/PostReactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).PostReactions(ctx, req.(*PostReactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_GetUserByPostID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByPostIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).GetUserByPostID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/GetUserByPostID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).GetUserByPostID(ctx, req.(*GetUserByPostIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_GetReactByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReactByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).GetReactByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/GetReactByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).GetReactByUserID(ctx, req.(*GetReactByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_DeleteReactionByPostID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReactionByPostIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).DeleteReactionByPostID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/DeleteReactionByPostID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).DeleteReactionByPostID(ctx, req.(*DeleteReactionByPostIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_DeleteReactionByUserPostID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReactionByUserPostIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).DeleteReactionByUserPostID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/DeleteReactionByUserPostID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).DeleteReactionByUserPostID(ctx, req.(*DeleteReactionByUserPostIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_UpdateReactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).UpdateReactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/UpdateReactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).UpdateReactions(ctx, req.(*UpdateReactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_TotalReactionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TotalReactionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).TotalReactionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/TotalReactionCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).TotalReactionCount(ctx, req.(*TotalReactionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactionService_ReactionCountByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionCountByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactionServiceServer).ReactionCountByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReactionService/ReactionCountByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactionServiceServer).ReactionCountByType(ctx, req.(*ReactionCountByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReactionService_ServiceDesc is the grpc.ServiceDesc for ReactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReactionService",
	HandlerType: (*ReactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostReactionTypes",
			Handler:    _ReactionService_PostReactionTypes_Handler,
		},
		{
			MethodName: "GetReactionTypeID",
			Handler:    _ReactionService_GetReactionTypeID_Handler,
		},
		{
			MethodName: "GetReactionTypes",
			Handler:    _ReactionService_GetReactionTypes_Handler,
		},
		{
			MethodName: "PostReactions",
			Handler:    _ReactionService_PostReactions_Handler,
		},
		{
			MethodName: "GetUserByPostID",
			Handler:    _ReactionService_GetUserByPostID_Handler,
		},
		{
			MethodName: "GetReactByUserID",
			Handler:    _ReactionService_GetReactByUserID_Handler,
		},
		{
			MethodName: "DeleteReactionByPostID",
			Handler:    _ReactionService_DeleteReactionByPostID_Handler,
		},
		{
			MethodName: "DeleteReactionByUserPostID",
			Handler:    _ReactionService_DeleteReactionByUserPostID_Handler,
		},
		{
			MethodName: "UpdateReactions",
			Handler:    _ReactionService_UpdateReactions_Handler,
		},
		{
			MethodName: "TotalReactionCount",
			Handler:    _ReactionService_TotalReactionCount_Handler,
		},
		{
			MethodName: "ReactionCountByType",
			Handler:    _ReactionService_ReactionCountByType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reactionTypes.proto",
}
