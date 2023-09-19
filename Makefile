build-linux:
    GOOS=linux GOARCH=amd64 go build -o ./build/iot-linux-amd64 ./cmd/iot-depin-protocold/main.go
    GOOS=linux GOARCH=arm64 go build -o ./build/iot-linux-arm64 ./cmd/iot-depin-protocold/main.go

build-darwin:
    GOOS=darwin GOARCH=amd64 go build -o ./build/iot-darwin-amd64 ./cmd/iot-depin-protocold/main.go
    GOOS=darwin GOARCH=arm64 go build -o ./build/iot-darwin-arm64 ./cmd/iot-depin-protocold/main.go

build-linux-amd64:
    GOOS=linux GOARCH=amd64 go build -o ./build/iot-linux-amd64 ./cmd/iot-depin-protocold/main.go

build-darwin-arm64:
    GOOS=darwin GOARCH=arm64 go build -o ./build/iot-darwin-arm64 ./cmd/iot-depin-protocold/main.go