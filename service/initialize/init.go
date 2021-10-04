package initialize

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Initialize() *sql.DB  {
	log.Println("Start initialize")
	connectionInfo := fmt.Sprintf("user=%v password=%v host=localhost dbname=%v sslmode=disable",
		"postgres", "root", "postgres")
	db, err := sql.Open("postgres", connectionInfo)

	if err != nil {
		log.Println("Open fatal")
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Ping fatal")
		log.Fatal(err)
	} else {
		log.Println("Ping was successful")
	}

	if !isDatabaseExist("go", db) {
		_, err = db.Exec("CREATE DATABASE GO")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Db GO Already exist")
	}


	db.Close()

	connectionInfo = fmt.Sprintf("user=%v password=%v host=localhost dbname=%v sslmode=disable",
		"postgres", "root", "go")
	db, err = sql.Open("postgres", connectionInfo)

	if err != nil {
		log.Println("Open fatal")
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS CUSTOMERS (CustomerID INT NOT NULL, CustomerName VARCHAR(40) NOT NULL, PRIMARY KEY (CustomerID))")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ORDERS (OrderID INT NOT NULL, OrderName VARCHAR(64) NOT NULL, CustomerID INT NOT NULL, PRIMARY KEY (OrderID))")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("DELETE FROM CUSTOMERS WHERE CustomerID > 0")
	db.Exec("DELETE FROM ORDERS WHERE OrderID > 0")


	db.Exec("INSERT INTO CUSTOMERS VALUES (1, 'DAN')")
	db.Exec("INSERT INTO CUSTOMERS VALUES (2, 'ANTON')")

	db.Exec("INSERT INTO ORDERS VALUES (1, 'PIZZA', 1)")
	db.Exec("INSERT INTO ORDERS VALUES (2, 'BURGER', 2)")
	db.Exec("INSERT INTO ORDERS VALUES (3, 'BEER', 2)")

	log.Println("End initialize")
	return db
}

func isDatabaseExist(dbName string, db *sql.DB) bool {
	rows, _ := db.Query("SELECT 1 FROM pg_database WHERE datname = $1", dbName)
	if rows.Next() {
		return true
	}
	return false
}