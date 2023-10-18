package main

import (
	"github.com/example/grpc_sample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Print("Start server.")

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc_sample.RegisterSampleServiceServer(grpcServer, &Sample{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}

	log.Print("Finish")
}

type Sample struct {
	name string
}

func (s *Sample) GetData(
	ctx context.Context,
	message *grpc_sample.Message,
) (*grpc_sample.Message, error) {
	log.Print(message.Body)
	return &grpc_sample.Message{Body: "It's a server response."}, nil
}


