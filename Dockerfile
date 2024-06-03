FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV GOPROXY=direct

RUN go build -o /app/main ./cmd/server/main.go


FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/main /app/main

WORKDIR /app
EXPOSE 8000

CMD ["/app/main"]
