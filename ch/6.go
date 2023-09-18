package main

import (
	"fmt"
	"strings"
)

func main() {

	//data := []int{1, 2, 3, 4, 5, 6, 6}
	//res := isUnique(data)
	//fmt.Println(res)

	datas := []string{"dauren", "aizek", "amina"}
	result := startWith(datas, "a")
	fmt.Println(result)
}

func isUnique(data []int) bool {
	un := make(map[int]bool)
	for _, item := range data {
		_, ok := un[item]
		value, exist := un[item]
		if ok {
			return false
		}
		un[item] = true
	}
	return true
}

func startWith(slice []string, prefix string) []string {
	result := []string{}
	for _, item := range slice {
		if strings.HasPrefix(item, prefix) {
			result = append(result, item)
		}
	}
	return result
}
