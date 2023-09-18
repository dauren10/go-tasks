str1 := "Hello"
str2 := " World"
concatenatedStr := str1 + str2


import "fmt"

str1 := "Hello"
str2 := " World"
concatenatedStr := fmt.Sprintf("%s%s", str1, str2)



import "strings"

var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" World")
concatenatedStr := builder.String()