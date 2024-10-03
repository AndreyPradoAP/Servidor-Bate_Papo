package main

import (
	"Servidor-Bate_Papo/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*req = &proto.Message{
	Name:    "Andrey",
	Time:    "10h30",
	Message: "Comer tatu é bom",
}*/

var mensagem *proto.Message = &proto.Message{
	Name:    "",
	Time:    "",
	Message: "",
}

// Crio a struct do server
type Server struct {
	proto.UnimplementedChatServer
}

// Função sendMessage arquivo .proto
func (service *Server) SendMessage(ctx context.Context, message *proto.Message) (*proto.Void, error) {
	println(message.Name, " - ", message.Time, "\n\t", message.Message)

	mensagem = &proto.Message{
		Name:    message.Name,
		Time:    message.Time,
		Message: message.Message,
	}

	return &proto.Void{}, nil
}

// Função receiveMessage arquivo .proto
func (service *Server) ReceiveMessage(void *proto.Void, streamMessages grpc.ServerStreamingServer[proto.Message]) error {
	// Repetir até dar erro
	for {
		select {
		// Verifica se o cliente fechou a conexão
		case <-streamMessages.Context().Done():
			return status.Error(codes.Canceled, "Chat Fecahdo")
		// Envia a última mensagem recebida pelo servidor de qualquer cliente
		default:
			err := streamMessages.SendMsg(mensagem)

			if err != nil {
				return status.Error(codes.Canceled, "Chat Fecahdo")
			}

			mensagem = &proto.Message{
				Name:    "",
				Time:    "",
				Message: "",
			}
		}
	}
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
