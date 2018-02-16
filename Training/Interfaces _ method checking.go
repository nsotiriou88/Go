package main

import (
	"fmt"
	"math"
	"reflect"
)

type square struct { // Implements shape interface,
	// by having the area(), which is a method.
	// This means that sqaure is also of type shape.
	side float64
}

// another shape
type circle struct {
	radius float64
}

type shape interface { // Type interface
	area() float64 // Includes the area() method signature
	// This means that anything that has this signature
	// implements the shape interface.
	// Gives polymorphic ability to our methods.
}

// We attach a method to type square. Generally, when
// we have a receiver in a function, we mean that it
// is a method for the specific type.
func (z square) area() float64 {
	return z.side * z.side
}

// which also implements the shape interface
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// So now, both cirlce and square, implement the
// shape interface.

func info(z shape) {

	// Checking what type of interface we passed in the
	// function.
	switch v := z.(type) { // myInterface.(type)
	case circle:
		// Pay attention here that the value of v is the z.
		// We could also use the z.(type) in switch and then
		// use z instead of v directly into the values.
		fmt.Printf("Circle with radius: %v\n", v)
	case square:
		fmt.Printf("Square with sides: %v\n", v)
	default:
		fmt.Printf("I don't know, ask stackoverflow.\n")
	}

	fmt.Println("<------------")
	// This way is not working to define which type is passing
	// in the function. We have to use the method above to check.
	if k := reflect.TypeOf(z).Kind(); k == reflect.Struct {
		fmt.Println(z, "is the length of Sides/Radius")
		fmt.Printf("The types of k are: %T - %s - %v\n", k, k, k)
		// it is always getting in this one.
	} else {
		fmt.Println(z, "is the length of Sides/Radius")
		fmt.Printf("The types of k are: %T - %s - %v\n", k, k, k)
		// This one never executes.
	}
	fmt.Println("------------>")
	fmt.Println(reflect.TypeOf(z)) // Not even with this value.
	fmt.Println("The area is:", z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	fmt.Println("***********************************")
	info(c)
}

// Method Sets; what we can pass in an interface when
// it has a certain receiver type. In pointer receivers
// we can send both values and pointers to a value!
// 
// Receivers	Values
// (t T)		T and *T
// (t *T)		*T
