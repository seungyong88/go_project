package main

import (
	"fmt"
)


type item struct {
	name string
	price float64
	quantity int
}

func New(name string, price float64, quantity int) *item {
	if price <= 0 || quantity <= 0 || len(name) == 0 {
		return nil
	}

	return &item{name, price, quantity}
} 

func main() {
	shirt := item.New("mens", 25000, 3)

	fmt.Println(shirt);
}