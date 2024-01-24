package main

func combat(health, damage float64) float64 {
	if health >= damage {
		return damage
	} else {
		return 0.0
	}
}

// func main() {
// 	fmt.Println((combat(1.0, 100)))
// }
