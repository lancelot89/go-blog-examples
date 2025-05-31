package handler

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/hello?name=Go", nil)
    rr := httptest.NewRecorder()

    HelloHandler(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("unexpected status: %d", rr.Code)
    }
    want := `{"message":"Hello, Go"}`
    if strings.TrimSpace(rr.Body.String()) != want {
        t.Fatalf("body = %s, want %s", rr.Body.String(), want)
    }
}
