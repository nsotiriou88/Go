package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20, // Bare in mind the commas required
		}, // before closing the {}.
		LicenseToKill: true, // Bare in mind the comma.
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		LicenseToKill: false,
	}
	// Generally we can do overriding in methods and fields.
	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
	fmt.Println("-----------------")
	// Pointers in struct
	p3 := &person{"Johnny", "-", 29}
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	fmt.Println(p3.First)
	fmt.Println(p1.person)
	// It is like using * for getting the value.
}
