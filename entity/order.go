package entity

import (
	"fmt"
	"time"
)

type Order struct {
	ID          int
	ProductId   int
	BuyerEmail  string
	BuyerAdrees string
	OrderDate   time.Time
	Product     Product
}

func (o *Order) PrintDetail() {
	fmt.Println("Email:", o.BuyerEmail)
	fmt.Println("Addrees:", o.BuyerAdrees)
	fmt.Println("Product:", o.Product.Name)
	fmt.Println("Order Date:", o.OrderDate)
}
