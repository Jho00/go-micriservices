package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-microservices/initialize"
	"go-microservices/internal"
	"log"
	"net/http"
)

// Create an exported global variable to hold the database connection pool.
var db *sql.DB = initialize.Initialize()

func main()  {
	http.HandleFunc("/getCustomers", internal.ProduceGetCustomerHandler(db))
	http.HandleFunc("/getOrders", internal.ProduceOrderCustomerHandler(db))


	http.HandleFunc("/ws/getCustomers", internal.ProduceWsEndpointForCustomerApi(db))
	http.HandleFunc("/ws/getOrders", internal.ProduceWsEndpointForOrderApi(db))



	log.Fatal(http.ListenAndServe(":8080", nil))
}
