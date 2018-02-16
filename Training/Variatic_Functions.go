package main

import "fmt"

func main() {
	n := average(43, 56, 85, 12, 44, 57)
	fmt.Println(n)
	data := []float64{46, 56, 85, 15, 44, 57}
	m := average(data...)
	fmt.Println(m)
	k := average2(data)
	fmt.Println(k)
	// About a function within a function-only if set on a variable.
	greeting := func() {
		fmt.Println("Hello World!")
	}
	greeting()                    // It is not printed twice, only once. Not when defining above!
	fmt.Printf("%T \n", greeting) // Print the type of the greeting variable
	// which is function()...
	greet := makeGreeter() // Check the function below!
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}

// This is if we pull out of the data structure one by one
// with ... or if we manually pass one by one the arguments.
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// This is if you pass a data structure
func average2(sf []float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))

}
func makeGreeter() func() string {
	return func() string {
		return "Hello world!!!"
	}
}

// This function is interesting for the structure of return.
// Returns func() string TYPE.
