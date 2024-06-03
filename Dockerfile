# Usar uma imagem base do Golang para construir a aplicação
FROM golang:1.21-alpine as builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos go.mod e go.sum e baixa as dependências
COPY go.mod go.sum ./

# Copia o código fonte da aplicação
COPY . .
COPY ./cmd/server/.env .
RUN ls -la

ENV GOPROXY="https://goproxy.io"
RUN go mod tidy

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server/main.go

# Utiliza uma imagem alpine mais leve para rodar a aplicação
FROM alpine:latest

COPY --from=builder /app/server .
COPY ./cmd/server/.env .
# Instalar dependências adicionais necessárias
RUN apk add --no-cache libc6-compat

# Define o diretório de trabalho e a porta que será exposta
EXPOSE 8000

# Comando para rodar a aplicação
CMD ["./server"]
