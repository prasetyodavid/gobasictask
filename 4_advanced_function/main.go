package main

import "fmt"

func main() {

	/*defer function*/
	defer fmt.Println("Run at the end of main function")

	defer func() {
		fmt.Println("Run at the end of main function as function")
	}()

	/*real sample of defer*/
	fmt.Println("Open DB")
	defer fmt.Println("Close DB")
	fmt.Println("Transaction DB")

	/*standart function*/
	fmt.Println(sum([]int{12, 4}))

	/*variadic function*/
	fmt.Println(sum_variadic(1, 2, 3, 4))

	/*anonymous function*/

	func() {
		fmt.Println("Anon Func 1")
	}()

	anonfunc2 := func() {
		fmt.Println("Anon Func 2")
	}

	go func() {
		fmt.Println("Anon Func Concurrent")
	}()

	anonfunc2()

	/*closure function*/
	counterrun := newCounter()
	fmt.Println(counterrun()) //return 1
	fmt.Println(counterrun()) //return 2
	fmt.Println(counterrun()) //return 3

}

func sum(data []int) int {
	result := 0
	for _, v := range data {
		result += v
	}
	return result
}

func sum_variadic(data ...int) int {
	result := 0
	for _, val := range data {
		result += val
	}
	return result
}

func newCounter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}
