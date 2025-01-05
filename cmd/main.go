package main

import (
	"flag"
	"fmt"
	"os"

	cli "grpccli_srv/internal/client"
	logs "grpccli_srv/internal/logs"
	server "grpccli_srv/internal/servidor"
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
		cli.RunClient()
	} else if *serverFlag {
		// Lógica do servidor
		s := server.Server{}
		s.Start()
	} else {
		fmt.Println("Por favor, use uma flag para especificar o modo: -c para cliente ou -s para servidor")
		os.Exit(1)
	}
	logs.Logger.Info("Finalizando o software")
}
