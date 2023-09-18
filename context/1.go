package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled")
			return
		default:
			// Выполнение работы
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer cancel() // Важно вызвать cancel() при завершении

	go doWork(ctx)

	// Здесь можно добавить другие задачи

	time.Sleep(10 * time.Second) // Подождать некоторое время

	// ctx будет автоматически отменен после 5 секунд
}
