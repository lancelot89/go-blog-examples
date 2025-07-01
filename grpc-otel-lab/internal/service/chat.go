package service

import (
	chatv1 "github.com/izayo/go-blog-examples/grpc-otel-lab/api/chat/v1"
	"io"
	"log"
	"time"
)

type ChatService struct {
	chatv1.UnimplementedChatServer
}

func (s *ChatService) Chat(stream chatv1.Chat_ChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("%s: %s", in.User, in.Text)
		in.SentAtUnix = time.Now().Unix()
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}
