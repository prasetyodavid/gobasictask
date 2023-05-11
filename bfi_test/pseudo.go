package main

import (
	"fmt"
)

func main() {

	var a int = 10
	var b int = 25
	var c int = 15
	var res int = 0
	for b > 0 {
		res += (a % c) + (c % a)
		b -= a % c
		a, c = c, a
	}
	fmt.Println(res)
}
