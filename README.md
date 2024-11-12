# pix-challenge
Convem pix challenge.

## Desafio da Convem para a 42 Rio.

Tive a honra de participar desse desafio proposto pela Convem, para realizar uma API que se integra com a API do ASAAS e realiza check-ins e check-outs.

### Infelizmente, não consegui fazer as seguintes tarefas:
- Integração com AWS.
- Docker
- Receber Webhooks

## Como rodar

### Dependências
- GO
- NodeJS

### Configure a chave API do ASAAS nas variáveis de ambiente do PC.
```sh
ASAASKEY=myKey
```

  
```sh
git clone https://github.com/marcosguimaraest/pix-challenge
```
1. Em outro terminal abra o back-end
```sh
cd server
go mod tidy
go get .
go run .
```
http://localhost:8080

2. Em um terminal abra o front-end
```sh
cd app/pix-challenge
npm --force i
npm run build
npm run start
```
http://localhost:3000

