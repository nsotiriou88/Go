package main

import "fmt"

func main() {
	x := []int{4, 7, 22, 15, 37, 1}
	// Alternatively
	// max := varial(4, 7, 22, 15, 37, 1)
	max := varial(x...)
	fmt.Println(max)
}

func varial(n ...int) int {
	var max int
	for _, m := range n {
		if max < m {
			max = m
		}
	}
	return max
}
