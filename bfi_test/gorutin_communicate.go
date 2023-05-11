package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chn := make(chan int, 1)
	defer close(chn)

	i := 0
	for i < 3 {
		fmt.Println("Sending value to channel.")
		go send(chn)
		fmt.Println("Receiving from channel.")
		go receive(chn)
		i++
		fmt.Println("Sleeping 3 seconds.")
		time.Sleep(time.Second * 3)
	}
}
func send(ch chan int) {
	val := rand.Intn(100)
	ch <- val
	fmt.Printf("Value %d sent.\n", val)
}

func receive(ch chan int) {
	time.Sleep(time.Second * 2)
	val := <-ch
	fmt.Printf("Value received: %d\n", val)

}
