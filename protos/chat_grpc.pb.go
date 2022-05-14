// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/chat.proto

package protos

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	Writing(ctx context.Context, opts ...grpc.CallOption) (Chat_WritingClient, error)
	WhoIsWriting(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (Chat_WhoIsWritingClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error)
	ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (Chat_ReceiveMessageClient, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error)
	Disconnect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Writing(ctx context.Context, opts ...grpc.CallOption) (Chat_WritingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/protos.Chat/Writing", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatWritingClient{stream}
	return x, nil
}

type Chat_WritingClient interface {
	Send(*StatusRequest) error
	CloseAndRecv() (*StatusReply, error)
	grpc.ClientStream
}

type chatWritingClient struct {
	grpc.ClientStream
}

func (x *chatWritingClient) Send(m *StatusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatWritingClient) CloseAndRecv() (*StatusReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StatusReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) WhoIsWriting(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (Chat_WhoIsWritingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[1], "/protos.Chat/WhoIsWriting", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatWhoIsWritingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_WhoIsWritingClient interface {
	Recv() (*StatusReply, error)
	grpc.ClientStream
}

type chatWhoIsWritingClient struct {
	grpc.ClientStream
}

func (x *chatWhoIsWritingClient) Recv() (*StatusReply, error) {
	m := new(StatusReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[2], "/protos.Chat/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatSendMessageClient{stream}
	return x, nil
}

type Chat_SendMessageClient interface {
	Send(*MessageRequest) error
	CloseAndRecv() (*MessageReply, error)
	grpc.ClientStream
}

type chatSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatSendMessageClient) Send(m *MessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatSendMessageClient) CloseAndRecv() (*MessageReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) ReceiveMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (Chat_ReceiveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[3], "/protos.Chat/ReceiveMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatReceiveMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ReceiveMessageClient interface {
	Recv() (*MessageReply, error)
	grpc.ClientStream
}

type chatReceiveMessageClient struct {
	grpc.ClientStream
}

func (x *chatReceiveMessageClient) Recv() (*MessageReply, error) {
	m := new(MessageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error) {
	out := new(ConnectReply)
	err := c.cc.Invoke(ctx, "/protos.Chat/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Disconnect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error) {
	out := new(ConnectReply)
	err := c.cc.Invoke(ctx, "/protos.Chat/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	Writing(Chat_WritingServer) error
	WhoIsWriting(*StatusRequest, Chat_WhoIsWritingServer) error
	SendMessage(Chat_SendMessageServer) error
	ReceiveMessage(*MessageRequest, Chat_ReceiveMessageServer) error
	Connect(context.Context, *ConnectRequest) (*ConnectReply, error)
	Disconnect(context.Context, *ConnectRequest) (*ConnectReply, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) Writing(Chat_WritingServer) error {
	return status.Errorf(codes.Unimplemented, "method Writing not implemented")
}
func (UnimplementedChatServer) WhoIsWriting(*StatusRequest, Chat_WhoIsWritingServer) error {
	return status.Errorf(codes.Unimplemented, "method WhoIsWriting not implemented")
}
func (UnimplementedChatServer) SendMessage(Chat_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) ReceiveMessage(*MessageRequest, Chat_ReceiveMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (UnimplementedChatServer) Connect(context.Context, *ConnectRequest) (*ConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatServer) Disconnect(context.Context, *ConnectRequest) (*ConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_Writing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Writing(&chatWritingServer{stream})
}

type Chat_WritingServer interface {
	SendAndClose(*StatusReply) error
	Recv() (*StatusRequest, error)
	grpc.ServerStream
}

type chatWritingServer struct {
	grpc.ServerStream
}

func (x *chatWritingServer) SendAndClose(m *StatusReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatWritingServer) Recv() (*StatusRequest, error) {
	m := new(StatusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_WhoIsWriting_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).WhoIsWriting(m, &chatWhoIsWritingServer{stream})
}

type Chat_WhoIsWritingServer interface {
	Send(*StatusReply) error
	grpc.ServerStream
}

type chatWhoIsWritingServer struct {
	grpc.ServerStream
}

func (x *chatWhoIsWritingServer) Send(m *StatusReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).SendMessage(&chatSendMessageServer{stream})
}

type Chat_SendMessageServer interface {
	SendAndClose(*MessageReply) error
	Recv() (*MessageRequest, error)
	grpc.ServerStream
}

type chatSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatSendMessageServer) SendAndClose(m *MessageReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatSendMessageServer) Recv() (*MessageRequest, error) {
	m := new(MessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_ReceiveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).ReceiveMessage(m, &chatReceiveMessageServer{stream})
}

type Chat_ReceiveMessageServer interface {
	Send(*MessageReply) error
	grpc.ServerStream
}

type chatReceiveMessageServer struct {
	grpc.ServerStream
}

func (x *chatReceiveMessageServer) Send(m *MessageReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chat/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chat/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Disconnect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Chat_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Chat_Disconnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Writing",
			Handler:       _Chat_Writing_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "WhoIsWriting",
			Handler:       _Chat_WhoIsWriting_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _Chat_SendMessage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveMessage",
			Handler:       _Chat_ReceiveMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/chat.proto",
}