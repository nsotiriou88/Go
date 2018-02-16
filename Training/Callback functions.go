package main

import "fmt"

func main() {
	visit([]int{1, 5, 3, 7}, func(n int) {
		fmt.Println(n)
	})
	xsm := filter([]int{1, 7, 3, 11}, func(m int) bool {
		return m > 3
	})
	fmt.Println(xsm) // [2 3 4]
}

func visit(numbers []int, callback func(int)) {
	fmt.Println(numbers)
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers2 []int, callback2 func(int) bool) []int {
	var xs []int
	for _, m := range numbers2 {
		if callback2(m) {
			xs = append(xs, m)
		}
	}
	return xs
}

// callback: passing a function as an argument and it is
// going back to the main where it is defined and passing
// a number and printing it.
