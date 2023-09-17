package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello from goroutine 1!"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Hello from goroutine 2!"
	}()

	// Use select to wait for either ch1 or ch2 to become ready
	select {
	case message1 := <-ch1:
		fmt.Println(message1)
	case message2 := <-ch2:
		fmt.Println(message2)
	}

	fmt.Println("Done!")
}
