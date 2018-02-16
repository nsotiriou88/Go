package main

import "fmt"

func main() {
	var s string
	i := 2
	switch i {
	case 1:
		fmt.Println("i is", i)
	case 2:
		fmt.Println("i is", i)
	default:
		fmt.Println("error")
	}
	fmt.Println("Give me your full Name")
	fmt.Scanf("%q", &s)
	fmt.Println(s) // You have to input quotes "" and the whole
	// string that you want inside them
}
