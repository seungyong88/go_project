package main

import "fmt"

// 사용자 정의 타입 정의(구조체)
type item struct {
	name string
	price float64
	quantity int
}

// item 타입에 Cost() 메서드 정의

func (t item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func main() {
	// 아이템 값 생성
	shirt := item{name:"seungyong", price: 25000, quantity:3}
	fmt.Println(shirt.Cost())
}