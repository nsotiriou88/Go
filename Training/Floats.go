package main

import "fmt"

func zero(z *int) { // *int instead of int
	*z = 0 //*x instead of just x
	fmt.Printf("%p\n", &z)
	fmt.Println(z)
}

func main() {
	x := 2
	var y float64
	zero(&x) //passing the pointer's address to x in func
	fmt.Println(x)
	// Like this, we can change the value to zero
	// without passing the pointer to func zero, we can't get
	// the value zero back.
	fmt.Printf("%p\n", &x) //printing the address in Hexadecimal
	// Toggling commend line with cmd + /
	y = 15 / 2.0 //you need 2.0 instead of 2 in order to print this as
	// a float. When you declare a variable with Var, you use =
	// instead of := .
	fmt.Printf("%.2f\n", y)
}
