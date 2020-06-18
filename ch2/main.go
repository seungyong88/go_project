package main

import (
	"fmt"
	mylib "./lib"
)

var v rune

func init() {
	v = '1'
}

func main() {
	fmt.Println("잠시 제거", v)
	fmt.Println(mylib.IsDigit(v))
// 	fmt.Println(mylib.IsDigit('a'))
// 	fmt.Println(mylib.IsSpace('\t'))
}