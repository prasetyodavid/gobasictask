package main

import "fmt"

func main() {
	ch := make(chan string)
	defer close(ch)

	go sendToChan("Dave", ch)
	go sendToChan("Mark", ch)
	go sendToChan("Jason", ch)

	fmt.Println(<-ch, <-ch, <-ch)
}

func sendToChan(name string, ch chan string) {
	msg := "Hello " + name
	ch <- msg
}
