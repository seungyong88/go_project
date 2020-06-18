package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("안녕, 내 이름은 ", p.Name, "야")
}

type Android struct {
	Person Person
	Modal  string
}

func main() {
	p := Person{"승용"}
	a := new(Android)
	a.Talk()
}
