package gateway

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/swaggo/http-swagger/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/usa4ev/parser_sample/internal/grpcsrv/protoparser"
)

func Run(grpcSrvEndpoint, httpSrvEndpointstring string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	gRPCmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterParserHandlerFromEndpoint(ctx, gRPCmux, grpcSrvEndpoint, opts)

	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	//mux.Handle("/swaggerui/", v5.NewHandler("My API", "/openapi.json", "/swaggerui/"))

	mux.Handle("/swaggerui/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/openapi.json")))


	mux.Handle("/", gRPCmux)

	mux.Handle("/openapi.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f, err := os.Open("../internal/grpcsrv/protoparser/openapiv2/server.swagger.json")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(os.Getwd())

		openAPIJSON, err := io.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}

		w.Write(openAPIJSON)
	}))

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(httpSrvEndpointstring, mux)
}
