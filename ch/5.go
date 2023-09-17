package main

import "fmt"

type MyInterface interface {
	Method1() string
}

type Mystruct struct {
	Name string
}

func (mt Mystruct) Method1() string {
	return mt.Name
}

func main() {
	var inf MyInterface

	d := Mystruct{
		Name: "dauren",
	}

	inf = d
	result := inf.Method1()
	fmt.Println(result)

}
