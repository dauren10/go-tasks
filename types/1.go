var val interface{} = 42

// Безопасное утверждение типа
result := val.(int)
fmt.Println(result) // Вывод: 42

// Безопасное утверждение типа с проверкой
if result, ok := val.(int); ok {
	fmt.Println(result) // Вывод: 42
} else {
	fmt.Println("Не удалось привести тип")
}