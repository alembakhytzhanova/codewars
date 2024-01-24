package main

import (
	"strconv"
	"strings"
)

func Points(games []string) int {
	var result int
	for _, match := range games {
		scores := strings.Split(match, ":")
		x, y := scores[0], scores[1]

		// converting to int
		scoresX, _ := strconv.Atoi(x)
		scoresY, _ := strconv.Atoi(y)

		switch {
		case scoresX > scoresY:
			result += 3
		case scoresX == scoresY:
			result++

		}
	}
	return result
}

// func main() {
// 	fmt.Println(Points([]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}))
// }
