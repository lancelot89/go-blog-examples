.PHONY: test tidy mock

test:
	go test ./... -v

tidy:
	go mod tidy

mock:
	mockgen -source=internal/domain/user.go -destination=mock/mock_user.go -package=mock
