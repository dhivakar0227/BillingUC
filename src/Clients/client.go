package main

import (
	"context"
	"fmt"
	billingpb "github/billing/src/ProtoBuffers"
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

	doUnary(c)
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
