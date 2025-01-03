# Defina o caminho do arquivo openssl.cnf
$opensslCnfPath = "C:\Program Files\OpenSSL-Win64\bin\openssl.cfg"

# Set environment variable for OpenSSL configuration file
$env:OPENSSL_CONF = $opensslCnfPath

# Gere a chave privada do CA
openssl genrsa -out ca.key 2048

# Gere o certificado do CA
openssl req -x509 -new -nodes -key ca.key -sha256 -days 1024 -out ca.crt -subj "/C=US/ST=State/L=City/O=Organization/OU=Org Unit/CN=example.com"

# Gere a chave privada do cliente
openssl genrsa -out client.key 2048

# Gere um CSR para o cliente
openssl req -new -key client.key -out client.csr -subj "/C=US/ST=State/L=City/O=Organization/OU=Org Unit/CN=client.example.com"

# Assine o CSR do cliente com o CA
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 500 -sha256

Write-Host "Certificados gerados com sucesso!"
