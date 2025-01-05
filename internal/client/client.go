package client

import (
	"context"
	"grpccli_srv/pb"
	"time"

	logs "grpccli_srv/internal/logs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func RunClient() {
	// Estabelecer a conexão com o servidor gRPC sem TLS
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logs.Logger.Sugar().DPanicf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMeuServicoClient(conn)

	// Preparar o contexto com timeout para a solicitação
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Enviar a solicitação ao servidor
	r, err := client.MeuMetodo(ctx, &pb.MeuRequest{Mensagem: "Olá, servidor!"})
	if err != nil {
		logs.Logger.Sugar().DPanicf("could not send message: %v", err)
	}
	logs.Logger.Sugar().Debugf("Resposta do servidor: %s", r.GetResposta())
}
