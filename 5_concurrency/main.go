package main

import (
	"fmt"
	"time"
)

func main() {
	go hello("hello")
	//time.Sleep(10000)
	hello("world")
	channel()
}

func hello(str string) {
	fmt.Println("Hello Go Concurrent")
	for i := 0; i < 5; i++ {
		time.Sleep(10000)
		fmt.Println(str)
	}
}

func channel() {
	fmt.Println("start main")
	c := make(chan string)
	go greet(c)
	c <- "Hello World"
	fmt.Println("end main")
}

func greet(c chan string) {
	c <- "Hello Chan"
}
