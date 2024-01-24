package main

func WordsToMarks(s string) int {
	var result int
	for _, str := range s {
		result += int(str) - 'a' + 1
	}
	return result
}

// func main() {
// 	fmt.Println(WordsToMarks("attitude"))
// }
