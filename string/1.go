str1 := "abcdef"
str2 := "abcxyz"

// Проверка, имеют ли строки одинаковую длину
if len(str1) != len(str2) {
	fmt.Println("Строки разной длины и не могут быть равными")
	return
}

// Итерируемся по символам и сравниваем их
for i := 0; i < len(str1); i++ {
	if str1[i] != str2[i] {
		fmt.Printf("Символ на позиции %d отличается: %c != %c\n", i, str1[i], str2[i])
		return
	}
}

fmt.Println("Строки идентичны")