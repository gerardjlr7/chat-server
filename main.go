package main

import (
	"log"
	"net"

	"github.com/gerardjlr7/chat-server/protos"
	"github.com/gerardjlr7/chat-server/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protos.RegisterChatServer(s, server.NewChatServer())
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
