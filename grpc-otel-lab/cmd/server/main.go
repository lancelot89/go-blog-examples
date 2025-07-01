package main

import (
	"log"
	"net"

	chatv1 "github.com/izayo/go-blog-examples/grpc-otel-lab/api/chat/v1"
	"github.com/izayo/go-blog-examples/grpc-otel-lab/internal/service"
	"github.com/izayo/go-blog-examples/grpc-otel-lab/otel"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	tp, ctx := otel.InitTracer("chat-server")
	defer tp.Shutdown(ctx)

	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer(
		// v0.61+ : StatsHandler 1 本で OK
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	chatv1.RegisterChatServer(s, &service.ChatService{})
	log.Println("server listening :50051")
	s.Serve(lis)
}
