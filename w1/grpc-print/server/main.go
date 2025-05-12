package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-print/internal/service"
	pb "grpc-print/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPrinterServer
	printerService *service.PrinterService
}

func (s *server) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	msg := req.GetMessage()
	log.Printf("New incoming gRPC request: %s", msg)

	if err := s.printerService.SaveMessage(msg); err != nil {
		log.Printf("Error saving message: %v", err)
		return nil, err
	}

	log.Println("Message saved successfully")
	return &pb.PrintResponse{Status: "Saved"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ps := service.NewPrinterService("messages.log")

	s := grpc.NewServer()
	pb.RegisterPrinterServer(s, &server{printerService: ps})

	fmt.Println("gRPC server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
