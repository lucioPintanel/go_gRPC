# Projeto de Mensagens Unary com gRPC em Go

## Objetivo
Este projeto tem como objetivo demonstrar o uso de Golang e gRPC para criar uma estrutura cliente/servidor. Vamos criar um sistema de mensagens unária onde o cliente envia uma mensagem ao servidor, e este responde.

## Requisitos
- Go 1.16 ou superior
- gRPC
- Protocolo Buffers
- Git

## Estrutura do Projeto
```
go_gRPC/
┣ certs/
┣ cmd/
┃ ┗ main.go
┣ controllers/
┃ ┗ produto_controller.go
┣ data/
┃ ┣ config.yaml
┃ ┗ database.db
┣ internal/
┃ ┣ client/
┃ ┃ ┗ client.go
┃ ┣ configs/
┃ ┃ ┣ configs.go
┃ ┃ ┣ db.go
┃ ┃ ┗ produto.go
┃ ┣ logs/
┃ ┃ ┗ logs.go
┃ ┣ servidor/
┃ ┃ ┗ servidor.go
┃ ┗ utils/
┣ models/
┃ ┗ produto.go
┣ pb/
┃ ┗ service.pb.go
┣ proto/
┃ ┗ service.proto
┣ repositories/
┃ ┗ produto_repository.go
┣ scripts/
┃ ┣ generate_certs.ps1
┃ ┗ generate_pb.ps1
┣ .gitignore
┣ go.mod
┣ go.sum
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
protoc --go_out=plugins=grpc:.\ proto\service.proto
```

Passo 4: Executar o Servidor
```
go run cmd/main.go -s
```

Passo 5: Executar o Cliente
```
go run cmd/main.go -c
```

Passo 6: Gerar a documentação e Executar o browser
```
.\scripts\generate_doc.ps1
```
