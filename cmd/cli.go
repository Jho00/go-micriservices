package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetCustomers() {
	resp, err := http.Get("http://localhost:8080/getCustomers")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}


func GetOrders(customerId int) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/getOrders?id=%v", customerId))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
