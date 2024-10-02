package main

import (
	"Servidor-Bate_Papo/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Crio a struct do server
type Server struct {
	proto.UnimplementedChatServer
}

// Função sendMessage arquivo .proto
func (service *Server) SendMessage(ctx context.Context, message *proto.Message) (*proto.Void, error) {
	return &proto.Void{}, nil
}

func main() {
	println("Servidor gRPC em funcionamento")

	// Listener para monitorar se o client requisitou conexão
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	// Criação Servidor gRPC
	grpcServer := grpc.NewServer()

	// Importo a struct server para o servidor gRPC
	proto.RegisterChatServer(grpcServer, &Server{})

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
