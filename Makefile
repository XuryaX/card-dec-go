build:
	go mod download
	go run cmd/main.go build

run:
	go run cmd/main.go run

test:
	go test -v ./tests/...