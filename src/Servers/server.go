package main

import (
	"context"
	"fmt"
	"log"
	"net"

	billingpb "Project/src/ProtoBuffers"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Server is running....")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Server failed to connect %v", err)
	}

	s := grpc.NewServer()
	billingpb.RegisterBillingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server is not serving %v", err)
	}
}

func (*server) SendInvoice(ctx context.Context, req *billingpb.SendInvoiceRequest) (*billingpb.SendInvoiceResponse, error) {
	firstName := req.Biller.GetFirstName()
	result := "Hello " + firstName
	resp := billingpb.SendInvoiceResponse{
		Result: result,
	}
	return &resp, nil
}
