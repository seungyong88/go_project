package main

import "fmt"

type Item struct {
	name string
	price float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

type DiscoutItem struct {
	Item
	discountRate float64
}

func main() {
	shoes := Item{"Women's Waling Shoes", 30000, 2}
	eventShoes := DiscoutItem{
		Item{"Sports Shoes", 50000, 3},
		10.00,
	}

	fmt.Println(shoes.Cost());
	fmt.Println(eventShoes.Cost());
}