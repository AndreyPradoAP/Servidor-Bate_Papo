package main

import (
	"Servidor-Bate_Papo/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Crio um novo client gRPC
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewChatClient(conn)

	userScreen(client)

	/*// Mensagem do client
	req := &proto.Message{
		Name:    "Andrey",
		Time:    "10h30",
		Message: "Comer tatu é bom",
	}

	// Execução função gRPC
	res, err := client.SendMessage(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)*/
}
