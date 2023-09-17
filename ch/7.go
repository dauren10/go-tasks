package main

import (
	"fmt"
	"sync"
)

func main() {
	readFromChannel()
}

func worker(item int, ch chan int, wg *sync.WaitGroup, resultChan chan []int) {
	defer wg.Done()
	value := <-ch
	result := []int{}
	result = append(result, value)
	resultChan <- result
}

func readFromChannel() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	var wg sync.WaitGroup

	ch := make(chan int)
	resultChan := make(chan []int)
	for _, item := range a {
		wg.Add(1)
		go worker(item, ch, &wg, resultChan)
	}

	for _, item := range b {
		wg.Add(1)
		go worker(item, ch, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(ch)
		close(resultChan) // Close result channel when all workers are done
	}()

	results := []int{}

	for _, item := range a {
		ch <- item
	}

	for _, item := range b {
		ch <- item
	}

	for result := range resultChan {
		results = append(results, result...)
	}

	fmt.Println(results)
}
