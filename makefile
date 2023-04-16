SRV_SRC=./cmd/main.go
SRV_BINARY_NAME=ParserServer
BIN_PATH=./bin
GRPC_SRV_ENDPOINT=localhost:9090
HTTP_SRV_ENDPOINT=localhost:8080


 # Build server
build-srv-linux:
	GOARCH=amd64 GOOS=linux go build -o $(BIN_PATH)/${SRV_BINARY_NAME}-linux $(SRV_SRC)

build-srv-darwin:
	GOARCH=amd64 GOOS=darwin go build -o $(BIN_PATH)/${SRV_BINARY_NAME}-darwin $(SRV_SRC)

build-srv-windows:
	GOARCH=amd64 GOOS=windows go build -o $(BIN_PATH)/${SRV_BINARY_NAME}-windows $(SRV_SRC)

# Run server
run-srv-linux: build-srv-linux
	$(BIN_PATH)/${SRV_BINARY_NAME}-linux -grpc-server-endpoint ${GRPC_SRV_ENDPOINT} -http-server-endpoint ${HTTP_SRV_ENDPOINT}

run-srv-darwin: build-srv-darwin
	$(BIN_PATH)/${SRV_BINARY_NAME}-darwin -grpc-server-endpoint ${GRPC_SRV_ENDPOINT} -http-server-endpoint ${HTTP_SRV_ENDPOINT}

run-srv-windows: build-srv-windows
	$(BIN_PATH)/${SRV_BINARY_NAME}-windows -grpc-server-endpoint ${GRPC_SRV_ENDPOINT} -http-server-endpoint ${HTTP_SRV_ENDPOINT}

test:
	go test -v "./..."

