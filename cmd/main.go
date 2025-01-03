package main

import (
	"context"
	"flag"
	"fmt"
	"grpccli_srv/pb"
	"os"

	logs "grpccli_srv/internal/logs"
	server "grpccli_srv/internal/servidor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Definindo as flags
	client := flag.Bool("c", false, "Executar como cliente")
	serverFlag := flag.Bool("s", false, "Executar como servidor")
	flag.Parse()

	// Inicializando o ponteiro
	logs.InitLogs()

	// Inicializando o log
	logs.Logger.Info("Iniciando aplicação...")

	if *client {
		// Lógica do cliente
		runClient()
	} else if *serverFlag {
		// Lógica do servidor
		s := server.Server{}
		s.Start()
	} else {
		fmt.Println("Por favor, use uma flag para especificar o modo: -c para cliente ou -s para servidor")
		os.Exit(1)
	}
}

func runClient() {
	logs.Logger.Info("Iniciando o cliente - localhost:50051")
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.Logger.Sugar().Fatalf("não foi possível se conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewMeuServicoClient(conn)
	r, err := c.MeuMetodo(context.Background(), &pb.MeuRequest{Mensagem: "Olá, servidor!"})
	if err != nil {
		logs.Logger.Sugar().Fatalf("erro ao chamar MeuMetodo: %v", err)
	}

	logs.Logger.Sugar().Debugf("Resposta do servidor: %s", r.Resposta)
}
