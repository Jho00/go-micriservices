package main

import (
	"bufio"
	"client/sdk"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Actions list:")
	fmt.Println("1. Get all customers")
	fmt.Println("2. Get all orders for customer")
	fmt.Print("Enter action number: ")
	for scanner.Scan() {
		text := scanner.Text()
		action, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Entered text must be a number!")
			continue
		}
		switch action {
		case 1:
			fmt.Println(sdk.GetCustomers())
			break
		case 2:
			fmt.Print("Enter customer id: ")
			scanner.Scan()
			text = scanner.Text()
			customerId, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Entered text must be a number!")
				break
			}
			orders := sdk.GetOrders(customerId)
			if orders == "[]" {
				fmt.Println("There are no orders for the target customer")
				break
			}
			fmt.Println(orders)
			break
		default:
			fmt.Println("Unknown action :(")
			break
		}
		fmt.Print("Enter action number: ")
	}
}