package main

import (
	"flag"

	"github.com/usa4ev/parser_sample/internal/gateway"
	"github.com/usa4ev/parser_sample/internal/grpcsrv"
	
)

var (
	// command-line options:
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	httpServerEndpoint = flag.String("http-server-endpoint", "localhost:8080", "HTTP server endpoint")
)

func run()error{
	srv := grpcsrv.New(*grpcServerEndpoint)

	go srv.ListenAndServe()

	return gateway.Run(*grpcServerEndpoint, *httpServerEndpoint)
}

func main() {
	flag.Parse()

	run()
}
