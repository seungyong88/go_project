package main

import "fmt"

// 사용자 정의 타입 정의(구조체)
type Item struct {
	name     string
	price    float64
	quantity int
}

// Item 타입에 Cost() 메서스 정의
func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func main() {
	shirt := Item{name: "asdsad", price: 25000, quantity: 3}
	fmt.Println(shirt.Cost())
}
