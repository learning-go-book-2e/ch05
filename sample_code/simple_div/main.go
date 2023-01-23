package main

import "fmt"

func main() {
	result := div(5, 2)
	fmt.Println(result)
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
