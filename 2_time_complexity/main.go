package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(isPrimeNumber(1000000007))
	duration := time.Since(start)
	fmt.Println("Time ", duration)
	start = time.Now()
	fmt.Println(isPrimeNumber(1500450271))
	duration = time.Since(start)
	fmt.Println("Time ", duration)

	fmt.Println(isPrimeNumber(100))
	fmt.Println(isPrimeNumber(35))
}

func isPrimeNumber(num int) string {
	//define prime number better

	fmt.Print(num)
	akarp2 := int(math.Sqrt(float64(num)))
	for i := 2; i <= akarp2; i++ {
		if num%i == 0 {
			return " Bukan Bilangan Prima"
		}
	}
	return " Bilangan Prima"
}
