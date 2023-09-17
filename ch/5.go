package main

import (
	"fmt"
)

// Определим интерфейс MyInterface
type MyInterface interface {
	Method1() string
	Method2(int) int
}

// Создадим структуру MyType, реализующую интерфейс MyInterface
type MyType struct {
	value int
}

func (mt MyType) Method1() string {
	return fmt.Sprintf("Method1 called with value %d", mt.value)
}

func (mt MyType) Method2(arg int) int {
	return mt.value + arg
}

func main() {
	// Создадим экземпляр MyType и присвоим его переменной интерфейса MyInterface
	var intf MyInterface
	mt := MyType{value: 42}
	intf = mt

	// Вызовем методы через интерфейс
	result1 := intf.Method1()
	result2 := intf.Method2(10)

	fmt.Println(result1) // Вывод: Method1 called with value 42
	fmt.Println(result2) // Вывод: 52
}
