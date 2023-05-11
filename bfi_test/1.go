package main

import "fmt"

func main() {

	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {

		if i%3 == 0 {
			// Multiple of 3
			fmt.Printf("Fizz")
		}
		if i%5 == 0 {
			// Multiple of 5
			fmt.Printf("Buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			// Neither, so print the number itself
			fmt.Printf("%d", i)
		}

		// A trailing new line (so both fizz + buzz can be printed on the same line)
		fmt.Printf("\n")

	}
}
