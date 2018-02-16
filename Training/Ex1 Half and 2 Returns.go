package main

import "fmt"

func main() {
	var x int
	fmt.Println("Give an integer number:")
	fmt.Scanf("%d", &x)
	a, b := checker(x)
	fmt.Println("The number devided by 2 is:", a)
	fmt.Println("The number is even:", b)
}

func checker(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0 // Define the n we grabed as float right now.
}
