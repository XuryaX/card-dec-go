build:
	go mod download
	go build -o ./bin/app.bin ./cmd
	go run cmd/main.go migrate

run:
	go run cmd/main.go run

test:
	go test -v ./tests/...