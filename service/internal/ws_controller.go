package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ProduceWsEndpointForCustomerApi(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Cannot establish ws connection")
			panic(err)
		}
		defer connection.Close()
		for {
			messageType, _, err := connection.ReadMessage()
			if err != nil {
				fmt.Println("Read message error")
				panic(err)
			}

			customers, err := json.Marshal(GetCustomers(db))

			if err != nil {
				log.Println(err)
				connection.WriteMessage(messageType, []byte("Error during getting customers"))
				return
			}
			connection.WriteMessage(messageType, customers)
		}
	}
}

func ProduceWsEndpointForOrderApi(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Cannot establish ws connection")
			panic(err)
		}
		defer connection.Close()
		for {
			messageType, customerIdText, err := connection.ReadMessage()
			if err != nil {
				fmt.Println("Read message error")
				panic(err)
			}

			orders, err := json.Marshal(GetOrders(db, string(customerIdText)))

			if err != nil {
				log.Println(err)
				connection.WriteMessage(messageType, []byte("Error during getting customers"))
				return
			}
			connection.WriteMessage(messageType, orders)
		}
	}
}