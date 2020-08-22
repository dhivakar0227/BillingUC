package main

import (
	"context"
	"fmt"
	billingpb "github/billing/src/ProtoBuffers"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started ...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	c := billingpb.NewBillingServiceClient(cc)

	//doUnary(c)

	//doServerStream(c)

	doClientStream(c)
}

func doUnary(c billingpb.BillingServiceClient) {

	s := &billingpb.SendInvoiceRequest{
		Biller: &billingpb.Bill{
			FirstName:   "Dhivakar",
			LastName:    "Jeganathan",
			InvoiceDate: "8/21/2022",
			InvoiceAmt:  2000,
		},
	}

	sir, err := c.SendInvoice(context.Background(), s)
	if err != nil {
		log.Fatalf("Error happened")
	} else {
		fmt.Println(sir)
	}

}

func doServerStream(c billingpb.BillingServiceClient) {
	s := &billingpb.ReceiveStreamInvoiceRequest{
		Biller: &billingpb.Bill{
			FirstName:   "Dhivakar",
			LastName:    "Jeganathan",
			InvoiceDate: "8/21/2022",
			InvoiceAmt:  2000,
		},
	}

	cresp, err := c.ReceiveStreamInvoice(context.Background(), s)
	if err != nil {
		log.Fatalf("error happened")
	}
	for {
		resp, err2 := cresp.Recv()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatalf("error happened")
		}

		fmt.Println(resp.Result)
	}

}

func doClientStream(c billingpb.BillingServiceClient) {

	stream, err := c.SendStreamInvoice(context.Background())
	if err != nil {
		log.Fatalf("error happened when starting the connection %v", err)
	}

	stream.Send(&billingpb.SendStreamInvoiceRequest{
		Biller: &billingpb.Bill{
			FirstName:   "John",
			LastName:    "Abraham",
			InvoiceDate: "08/22/2020",
			InvoiceAmt:  4000,
		},
	})

	time.Sleep(1000 * time.Millisecond)

	stream.Send(&billingpb.SendStreamInvoiceRequest{
		Biller: &billingpb.Bill{
			FirstName:   "Jack",
			LastName:    "Abraham",
			InvoiceDate: "08/22/2020",
			InvoiceAmt:  4000,
		},
	})

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error happened during receiving")
	}

	fmt.Println(resp.Result)
}
