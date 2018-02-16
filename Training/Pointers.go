package main

import "fmt"

func zero(z *int) { // *int instead of int
	fmt.Println(*z)
	*z = 0 //*x instead of just x
	fmt.Printf("%p\n", &z)
	fmt.Println(z)
	fmt.Println(*z)
}

func main() {
	x := 2
	zero(&x) //passing the pointer's address to x in func
	fmt.Println(x)
	// Like this, we can change the value to zero
	// without passing the pointer to func zero, we can't get
	// the value zero back.
	fmt.Printf("%p\n", &x) //printing the address in Hexadecimal
	// Toggling commend line with cmd + /
}
