package main

import (
	"fmt"
	"sync"

	"github.com/izayo/go-blog-examples/grpc-gateway-k6-observability/gateway"
	"github.com/izayo/go-blog-examples/grpc-gateway-k6-observability/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		server.Start()
	}()

	go func() {
		defer wg.Done()
		gateway.Start()
	}()

	fmt.Println("Successfully started server and gateway")
	wg.Wait()
}
