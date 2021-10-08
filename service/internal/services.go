package internal

import (
	"database/sql"
	"go-microservices/pb"
	"log"
)


func GetCustomers(db *sql.DB) []Customer  {
	rows, err := db.Query("SELECT * FROM CUSTOMERS")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	result := []Customer{}
	for rows.Next() {
		var r Customer
		err = rows.Scan(&r.Customerid, &r.Customername)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		result = append(result, r)
	}


	return result
}

func GetPbCustomers(db *sql.DB) []*pb.Customer  {
	rows, err := db.Query("SELECT * FROM CUSTOMERS")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	result := []*pb.Customer{}
	for rows.Next() {
		var r pb.Customer
		err = rows.Scan(&r.Customerid, &r.Customername)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		result = append(result, &r)
	}

	return result
}


func GetOrders(db *sql.DB, CustomerID string) []Order  {
	rows, err := db.Query("SELECT * FROM ORDERS WHERE Customerid = $1", CustomerID)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	result := []Order{}
	for rows.Next() {
		var r Order
		err = rows.Scan(&r.Orderid, &r.Ordername, &r.Customerid)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		result = append(result, r)
	}


	return result
}