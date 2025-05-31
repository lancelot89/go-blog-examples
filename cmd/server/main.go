package main

import (
    "log"
    "net/http"

    "github.com/izayo/go-blog-examples/internal/handler"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", handler.HelloHandler)

    log.Println("start :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}
