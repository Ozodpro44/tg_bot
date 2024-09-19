package main

import (
	"context"
	"log"
	"net"

	pb "path/to/your/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type orderServiceServer struct {
    pb.UnimplementedOrderServiceServer
}

func (s *orderServiceServer) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    // Forward the request to the DB service to save it to PostgreSQL
    // Simulating by logging the request
    log.Printf("Received order from user %d: %s", req.UserId, req.OrderText)

    // You'd normally call the DB service here using another gRPC call

    return &pb.OrderResponse{Success: true}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterOrderServiceServer(grpcServer, &orderServiceServer{})
    reflection.Register(grpcServer)

    log.Println("Order Service is running on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
