# Projeto de Mensagens Unary com gRPC em Go

## Objetivo
Este projeto tem como objetivo demonstrar o uso de Golang e gRPC para criar uma estrutura cliente/servidor. Vamos criar um sistema de mensagens unária onde o cliente envia uma mensagem ao servidor, e este responde.

## Requisitos
- Go 1.16 ou superior
- gRPC
- Protocolo Buffers
- Certificados SSL (server.crt e server.key)
- Git

## Estrutura do Projeto
```
go_gRPC/
┣ certs/
┃ ┣ ca.crt
┃ ┣ ca.key
┃ ┣ ca.srl
┃ ┣ client.crt
┃ ┣ client.csr
┃ ┗ client.key
┣ cmd/
┃ ┗ main.go
┣ data/
┣ internal/
┃ ┣ client/
┃ ┃ ┗ client.go
┃ ┣ configs/
┃ ┃ ┗ configs.go
┃ ┣ logs/
┃ ┃ ┗ logs.go
┃ ┣ servidor/
┃ ┃ ┗ servidor.go
┃ ┗ utils/
┣ proto/
┣ scripts/
┃ ┗ generate_certs.ps1
┣ .gitignore
┣ LICENSE
┗ README.md
```

## Como Executar
Passo 1: Clonar o Repositório
```
git clone https://github.com/usuario/meu_projeto.git
cd meu_projeto
```

Passo 2: Instalar Dependências
```
go mod tidy
```

Passo 3: Gerar Código a partir dos Arquivos .proto
```
protoc --go_out=plugins=grpc:./proto ./proto/meu_servico.proto
```

Passo 4: Executar o Servidor
```
go run cmd/main.go
```