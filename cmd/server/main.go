package main

import (
	"Servidor-Bate_Papo/proto"
	"context"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChatServer
}

func (service Server*) SendMessage(ctx context.Context, message *proto.Message) (*proto.Messagek)

func main() {
	grpcServer := grpc.NewServer()

	proto.RegisterChatServer(grpcServer, $Server)
}
