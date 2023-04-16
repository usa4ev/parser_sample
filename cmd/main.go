package main

import (
	"flag"
	"fmt"

	"github.com/usa4ev/parser_sample/internal/gateway"
	"github.com/usa4ev/parser_sample/internal/grpcsrv"
)

var (
	// command-line options:
	grpcServerEndpoint = flag.String("grpc-server-endpoint", ":9090", "gRPC server endpoint")
	httpServerEndpoint = flag.String("http-server-endpoint", ":8080", "HTTP server endpoint")
)

func run() error {
	srv := grpcsrv.New(*grpcServerEndpoint)

	fmt.Printf("starting grpc srerver on %v\n", *grpcServerEndpoint)

	go srv.ListenAndServe()

	fmt.Printf("starting gateway on %v\n", *httpServerEndpoint)

	return gateway.Run(*grpcServerEndpoint, *httpServerEndpoint)
}

func main() {
	flag.Parse()

	run()
}
