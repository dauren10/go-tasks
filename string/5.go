package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаем корневой контекст с сроком выполнения 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Гарантирует освобождение ресурсов

	// Создаем горутину, которая работает с контекстом
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Операция отменена или срок выполнения истек")
		case <-time.After(3 * time.Second):
			fmt.Println("Операция выполнена успешно")
		}
	}(ctx)

	// Ожидаем завершения горутины
	<-ctx.Done()
}
