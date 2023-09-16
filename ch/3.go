package main

import (
	"context"
)

func g(ctx context.Context) {
	ch := make(chan string)
	go func() { close(ch) }()
	select {
	case <-ch:
	case <-ctx.Done():
	}
}

func main() {
	ctx := context.Background()
	g(ctx)
}
