package main

import (
	"strconv"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string { // (p person) is a receiver.
	return p.first + " " + p.last + ", age: " + strconv.Itoa(p.age)
	// age needs to be a string because we want to return a string
}

// This is a method. This means that this function will be attached
// to the type person and any person value will have access to type
// person.
// Notice that this function has receiver (receiver, name(parameters)
// and returns). Any value of type person can call this function.

func main() {
	p1 := person{"James", "Bond", 20} // p1 and p2 are of type person.
	p2 := person{"Miss", "MoneyPenny", 18}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p1.fullName()) // Check how we call the function
	fmt.Println(p2.fullName()) // because of the receiver.
}
