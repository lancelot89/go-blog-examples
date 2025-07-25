FROM golang:1.24
WORKDIR /app
COPY . .
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/bufbuild/buf/cmd/buf@latest
CMD ["bash"]
