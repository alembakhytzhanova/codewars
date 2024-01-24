package main

import (
	"strings"
)

func StringToArray(str string) []string {
	var result []string
	slovo := strings.Split(str, " ")
	for _, ch := range slovo {
		result = append(result, string(ch))
	}
	return result
}

// func StringToArray(str string) []string {
// 	return strings.Fields(str)
// }

// func main() {
// 	fmt.Println(StringToArray("Robin Good"))
// }
