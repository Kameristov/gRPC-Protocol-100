BINARY_NAME=gRPC-Protocol-100.exe

build:
	go build -o ${BINARY_NAME} ./cmd/app/main.go

run: build
	./${BINARY_NAME}

