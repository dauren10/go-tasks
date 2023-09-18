package main

import "fmt"

func d() (string, error) {
	a := "Hello"
	return a, nil
}

func main() {

	if _, err := d(); err != nil {
		fmt.Println("da")
	}

	var b byte = 'A'
	fmt.Printf("%c\n", b)

	emoji := []rune("Привет")
	fmt.Printf("%c", emoji)

	data := map[string]interface{}{
		"name": "Bob",
		"age":  25,
	}
	jsonStr := `{"name":"Bob","age":25}`
	fmt.Println(jsonStr)
	fmt.Printf("%v", data)
}
