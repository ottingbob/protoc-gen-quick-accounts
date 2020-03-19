package server

import (
	"context"
	"fmt"
	bb "grpc-gw/buffb"
	pb "grpc-gw/example"
	"mime"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func RunREST() error {
	fmt.Println("rest called")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		gwmux              = runtime.NewServeMux()
		grpcServerEndpoint = "localhost:5001"
		opts               = []grpc.DialOption{grpc.WithInsecure()}
		err                = pb.RegisterYourServiceHandlerFromEndpoint(ctx, gwmux, grpcServerEndpoint, opts)
	)

	var (
		gw2mux        = runtime.NewServeMux()
		grpcEndpoint2 = "localhost:5002"
		err2          = bb.RegisterQuickServiceHandlerFromEndpoint(ctx, gw2mux, grpcEndpoint2, opts)
	)

	if err != nil {
		return err
	}

	if err2 != nil {
		return err
	}

	// fs := http.FileServer(http.Dir("./swaggerui"))
	// http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
	// mux.serveSwagger()
	// mux.HandleFunc("/swagger.json",)

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.Handle("/v1/", gw2mux)
	serveSwagger(mux)
	return http.ListenAndServe(":8080", mux)
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swaggeer-ui/ on <host>/swagger-ui
	// fileServer := http.FileServer(&assetfs.AssetFS{
	// 	Asset:    swagger.Asset,
	// 	AssetDir: swagger.AssetDir,
	// 	Prefix: "third_party/swagger-ui",
	// })

	fileServer := http.FileServer(http.Dir("./swagger-ui"))
	prefix := "/swaggerui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
