package internal

type Customer struct {
	Customerid   int    `json:"CustomerId"`
	Customername string `json:"CustomerName"`
}

type Order struct {
	Orderid    int    `json:"OrderId"`
	Ordername  string `json:"OrderName"`
	Customerid int    `json:"CustomerId"`
}
