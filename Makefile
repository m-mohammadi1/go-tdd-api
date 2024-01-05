run:
	go run ./cmd/api

test:
	go test -v -cover -race ./...


.PHONY: run test