package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"go-microservices/initialize"
	"go-microservices/internal"
	"go-microservices/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

// Create an exported global variable to hold the database connection pool.
var db *sql.DB = initialize.Initialize()

type grpcServer struct {
	pb.UnimplementedCustomerServiceServer
}
func (s *grpcServer) GetCustomers(ctx context.Context, in *pb.GetCustomersRequest) (*pb.GetCustomersResponse, error) {
	customers := internal.GetPbCustomers(db)

	return &pb.GetCustomersResponse{Customers: customers}, nil
}

func main()  {
	go func() {
		lis, err := net.Listen("tcp", ":5555")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterCustomerServiceServer(s, &grpcServer{})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		time.Sleep(2 + time.Second)

		internal.GetGrpcCustomers()
	}()

	http.HandleFunc("/getCustomers", internal.ProduceGetCustomerHandler(db))
	http.HandleFunc("/getOrders", internal.ProduceOrderCustomerHandler(db))


	http.HandleFunc("/ws/getCustomers", internal.ProduceWsEndpointForCustomerApi(db))
	http.HandleFunc("/ws/getOrders", internal.ProduceWsEndpointForOrderApi(db))


	log.Fatal(http.ListenAndServe(":8080", nil))
}
