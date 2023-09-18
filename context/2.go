package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	originalString := "Пример строки для изменения"

	// Создаем новую строку с измененным содержимым
	modifiedString := strings.Replace(originalString, "строки", "текста", -1)

	// Выводим измененную строку
	fmt.Println(modifiedString)
	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}

	// str1 := "abcdef"
	// str2 := "abc"

	// // Проверка, является ли str2 префиксом str1
	// isPrefix := strings.HasPrefix(str1, str2)

	// if isPrefix {
	// 	fmt.Println(str2, "является префиксом", str1)
	// } else {
	// 	fmt.Println(str2, "не является префиксом", str1)
	// }

	str1 := "abcdef"
	str2 := "def"

	// Проверка, является ли str2 суффиксом str1
	isSuffix := strings.HasSuffix(str1, str2)

	if isSuffix {
		fmt.Println(str2, "является суффиксом", str1)
	} else {
		fmt.Println(str2, "не является суффиксом", str1)
	}
}
