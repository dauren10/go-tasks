package main

import (
	"fmt"
	"time"
)

func worker(i int, ch chan int) {
	value := <-ch
	fmt.Printf("received %d %d", i, value)
}

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
	time.Sleep(time.Second)
}
