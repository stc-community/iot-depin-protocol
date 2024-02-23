.PHONY: all build clean
all:
	if [ ! -d "./build/" ]; then \
	mkdir bin; \
	fi
	GOOS=linux GOARCH=amd64 go build -o ./build/iot-linux-amd64 ./cmd/iot-depin-protocold/main.go