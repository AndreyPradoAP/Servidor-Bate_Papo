package main

import (
	"Servidor-Bate_Papo/proto"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func userScreen(client proto.ChatClient) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\n---------- CHAT INICIADO ----------\n\n")

	// Entrada do nome
	fmt.Printf("Para inciar, digite um nome de usuário\n:> ")
	scanner.Scan()

	nome := scanner.Text()

	fmt.Printf("\n--- O Chat está conectando. Digite a mensagem requerida, e aperte ENTER pra enviar. Para sair, basta digitar exit() ---\n")

	inicarBroadcast(client)

	var texto string

	// Incio chat
	for {
		fmt.Printf("\t:> ")
		scanner.Scan()

		texto = scanner.Text()

		// Verifica se o usuário quer sair do modo chat
		if texto == "exit()" {
			fmt.Printf("\nChat Finalizado\n")
			break
		}

		// Pega o horário da mensagem
		dataHora := time.Now()

		// Pega os dados e coloca no padrão proto.Message
		req := &proto.Message{
			Name:    nome,
			Time:    dataHora.String(),
			Message: texto,
		}

		// Envio a mensagem para o servidor
		_, err := client.SendMessage(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func inicarBroadcast(client proto.ChatClient) {
	mensagensStream, err := client.ReceiveMessage(context.Background(), &proto.Void{})

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			msg, err := mensagensStream.Recv()

			if err != nil {
				log.Fatal(err)
				return
			}
			if msg.Message != "" {
				fmt.Printf("%s -> %s\n", msg.Name, msg.Message)
			}
		}
	}()
}
