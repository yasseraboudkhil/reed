package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	service "github.com/yasseraboudkhil/reed/proto/v1"
)

//RunServer function registers ToDo service and starts gRPC server.
func RunServer(ctx context.Context, reedSvcServer service.ReedServiceServer, port string) error {

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil
	}

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, os.Interrupt)

	// register the Reed service
	grpcServer := grpc.NewServer()
	service.RegisterReedServiceServer(grpcServer, reedSvcServer)

	go func() {

		<-gracefulStop

		log.Println("Received an OS interrupt, so shutting down the gRPC server")
		grpcServer.GracefulStop()

	}()

	log.Println("Starting the gRPC server")
	// start the gRPC server
	return grpcServer.Serve(listen)
}
