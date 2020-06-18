package main

import "fmt"

type quantity int
type dozen []quantity

var d dozen

func main() {
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}

	fmt.Println(d)
}