package main

import (
	"fmt"
	"sync"
)

var links = []string{"https://example.com",
	"https://google.com",
	"https://github.com"}

func main() {
	counter := int64(0)
	mu := &sync.Mutex{}
	ch := make(chan string, 10)
	wg := &sync.WaitGroup{}
	for i := range links {
		ch <- links[i]
	}

	close(ch)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := range ch {
				if err := checkUrl(l); err == nil {
					mu.Lock()
					counter++
					mu.Unlock()
				}
			}

		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func checkUrl(l string) (err error) {
	fmt.Println(l)
	return nil
}
