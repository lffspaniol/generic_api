package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/lffspaniol/generic_api/gen"
	"github.com/lffspaniol/generic_api/temperature"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func main() {
	log.Println("Initializing the server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Falied to create the listener server: %v", err)
	}
	s := grpc.NewServer()
	temperature.Register(s)

	log.Println(("Listen..."))
	reflection.Register(s)
	go gateway()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falied to serve : %v", err)
	}

}

func gateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gen.RegisterTemperatureServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	log.Println("Initializing the reverse gateway")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Println("erro in start reverse gateway")
		return err
	}
	return nil
}
