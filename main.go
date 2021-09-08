package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-microservices/cmd"
	"go-microservices/initialize"
	"go-microservices/internal"
	"log"
	"net/http"
	"time"
)

// Create an exported global variable to hold the database connection pool.
var db *sql.DB = initialize.Initialize()

func main()  {
	http.HandleFunc("/getCustomers", internal.ProduceGetCustomerHandler(db))
	http.HandleFunc("/getOrders", internal.ProduceOrderCustomerHandler(db))

	go func() {
		time.Sleep(5 * time.Second)
		cmd.GetCustomers()
		cmd.GetOrders(2)
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
