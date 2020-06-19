package main

import "fmt"

type Item struct {
	name     string
	price    int32
	quantity int
}

func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	shoes := Item{name: "Sports Shoes", price: 30000}
	dress := Item{name: "Stripe Shift Dress", quantity: 2}
	phone := Item{"Amazon Fire Phone, 32GB", 21900, 1}

	fmt.Printf("%#v", shirt) // {Men's Slim-Fit Shirt 25000 3}
	fmt.Println(shoes) // {Sports Shoes 30000 0}
	fmt.Println(dress) // {Stripe Shift Dress 0 2}
	fmt.Println(phone) // {Amazon Fire Phone, 32GB 21900 1}
}