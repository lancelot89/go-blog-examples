package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pbecho "github.com/izayo/go-blog-examples/grpc-gateway-k6-observability/proto/echo"
)

const (
	port = ":50051"
)

type server struct {
	pbecho.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *pbecho.EchoRequest) (*pbecho.EchoResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pbecho.EchoResponse{Message: in.GetMessage()}, nil
}

func Start() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbecho.RegisterEchoServiceServer(s, &server{})
	fmt.Println("gRPC server listening on", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
