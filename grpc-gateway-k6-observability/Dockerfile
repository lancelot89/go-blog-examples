FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./grpc-gateway-k6-observability/cmd/main.go

EXPOSE 50051 8080

CMD ["/app/main"]
