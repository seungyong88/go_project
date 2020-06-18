package main

import "fmt"

func f2() (r int, s int) {
	r = 1 
	s = 2
	return
}

func makeEvenGenerator() func() int {
	i := int(0)
	// return func() (ret uint) {
	// 	ret = i
	// 	i += 2
	// 	return
	// }

	add := func() (rets int) {
		rets = i
		i += 2
		return
	}

	return add

}

func main() {
	nextEven := makeEvenGenerator()
	x, s := f2()
	fmt.Println(x)
	fmt.Println(s)
	fmt.Println(nextEven())
	// fmt.Println(nextEven())
}