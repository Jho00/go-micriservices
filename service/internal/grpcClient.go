package internal;

import (
	"context"
	"go-microservices/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetGrpcCustomers()  {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("Connected")

	defer conn.Close()
	c := pb.NewCustomerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetCustomers(ctx, &pb.GetCustomersRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}


	log.Printf("Greeting: %s", r.GetCustomers())
}
