# Inicia o servidor de documentação godoc
Start-Process "godoc" -ArgumentList "-http=:6060"

# Abre o navegador para visualizar a documentação
Start-Process "http://localhost:6060/pkg/grpccli_srv"
