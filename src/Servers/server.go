package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	billingpb "github/billing/src/ProtoBuffers"

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

func (*server) ReceiveStreamInvoice(req *billingpb.ReceiveStreamInvoiceRequest, stream billingpb.BillingService_ReceiveStreamInvoiceServer) error {
	firstName := req.Biller.GetFirstName()
	result := "Hello " + firstName
	for i := 0; i < 10; i++ {
		resp := billingpb.ReceiveStreamInvoiceResponse{
			Result: result,
		}
		stream.Send(&resp)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) SendStreamInvoice(stream billingpb.BillingService_SendStreamInvoiceServer) error {
	result := " "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return (stream.SendAndClose(&billingpb.SendStreamInvoiceResponse{
				Result: result,
			}))
		}
		if err != nil {
			log.Fatalf("Server is not serving %v", err)
		}

		result = result + " " + req.Biller.GetFirstName()
	}
}
