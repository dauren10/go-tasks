package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello, World!"
	}()

	select {
	case value := <-ch1:
		fmt.Println("Получено из ch1:", value)
	case message := <-ch2:
		fmt.Println("Получено из ch2:", message)
	case <-time.After(3 * time.Second):
		fmt.Println("Прошло 3 секунды. Таймаут.")
	}
}
