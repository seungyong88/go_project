package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(6))
}

// package main

// import "fmt"

// func factorial(n uint64) uint64 {
// 	if n == 0 {
// 		return 1
// 	}

// 	return n * factorial(n-1)
// }

// func main() {
// 	fmt.Println(factorial(5)) // 120
// }
