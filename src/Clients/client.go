package main

import (
	billingpb "Project/src/ProtoBuffers"
	"context"
	"fmt"
	"log"

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

	s := &billingpb.SendInvoiceRequest{
		Biller: &billingpb.Bill{
			FirstName:   "Dhivakar",
			LastName:    "Jeganathan",
			InvoiceDate: "8/21/2020",
			InvoiceAmt:  2000,
		},
	}
	sir, err := c.SendInvoice(context.Background(), s)
	fmt.Println(sir)
}
