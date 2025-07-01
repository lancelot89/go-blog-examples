generate:
	buf generate

run-server:
	go run grpc-otel-lab/cmd/server

run-client:
	go run grpc-otel-lab/cmd/client

deps:
	go get go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc@latest
	go mod tidy
