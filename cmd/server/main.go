package main

import (
	"Servidor-Bate_Papo/proto"
	"context"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChatServer
}

// Função pega do arquivo proto
func (service *Server) SendMessage(ctx context.Context, message *proto.Message) *proto.Message {
	println("Mensagem: ", req.Ge)
}

func main() {
	// Criação Servidor gRPC
	grpcServer := grpc.NewServer()

	// Registro do serviço do arquivo .proto
	proto.RegisterChatServer(grpcServer, &Server)
}
