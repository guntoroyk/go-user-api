build:
	go build -o bin/go-user-api main.go

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/go-user-api-darwin-amd64 main.go

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/go-user-api-darwin-arm64 main.go

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/go-user-api-linux-amd64 main.go

build-linux-arm:
	GOOS=linux GOARCH=arm go build -o bin/go-user-api-linux-arm main.go

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/go-user-api-linux-arm64 main.go

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/go-user-api-windows-amd64 main.go

build-windows-arm:
	GOOS=windows GOARCH=arm go build -o bin/go-user-api-windows-arm main.go

build-windows-arm64:
	GOOS=windows GOARCH=arm64 go build -o bin/go-user-api-windows-arm64 main.go

build-all:
	make build-darwin-amd64
	make build-darwin-arm64
	make build-linux-amd64
	make build-linux-arm
	make build-linux-arm64
	make build-windows-amd64
	make build-windows-arm
	make build-windows-arm64

run:
	go run main.go