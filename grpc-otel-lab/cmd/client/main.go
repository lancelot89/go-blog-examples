package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	chatv1 "github.com/izayo/go-blog-examples/grpc-otel-lab/api/chat/v1"
	"github.com/izayo/go-blog-examples/grpc-otel-lab/otel"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// ----- OpenTelemetry 初期化 -----
	tp, _ := otel.InitTracer("chat-client")
	defer tp.Shutdown(ctx)

	// ----- gRPC クライアント構築（NewClient API） -----
	cli, err := grpc.NewClient(
		"server:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // SSL 無効 (ローカル)
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),       // OTEL トレース付与
	)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// DialContext の WithBlock 相当：接続確立まで待機
	cli.Connect()

	// ----- 双方向ストリーム開始 -----
	client := chatv1.NewChatClient(cli)
	stream, err := client.Chat(ctx)
	if err != nil {
		panic(err)
	}

	// サーバ → クライアント 受信ゴルーチン
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				fmt.Println("server closed:", err)
				return
			}
			fmt.Printf("%s> %s\n", in.User, in.Text)
		}
	}()

	// 標準入力をサーバへ送信
	user := "you"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		err := stream.Send(&chatv1.Message{
			User:       user,
			Text:       txt,
			SentAtUnix: time.Now().Unix(),
		})
		if err != nil {
			fmt.Println("send error:", err)
			return
		}
	}
}
