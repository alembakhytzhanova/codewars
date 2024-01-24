package main

import (
	"strings"
)

func Order(sentence string) string {
	str1 := strings.Split(sentence, " ")
	m := map[string]string{}
	ints := "123456789"
	result := ""
	for _, v := range str1 {
		for _, h := range v {
			for _, v2 := range ints {
				if h == v2 {
					m[string(v2)] = v
					break
				}
			}
		}
	}

	for i, v := range ints {
		if m[string(v)] == ""{
			break
		}
		result += m[string(v)]
		if i != len(str1)-1 {
			result += " "
		}
	}
	return result
}

func main() {
	Order("is2 Thi1s T4est 3a")
}
