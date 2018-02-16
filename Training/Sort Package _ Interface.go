package main

import (
	"fmt"
	"sort"
)

// Creating my own type
type people []string

type person struct {
	Name string
	Age  int
}

// Building the required methods for the interface Interface, that we need in
// order to use the Sort method.
func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// This method does everything at the same time(one line), so there is no
// need for another temp variable. All methods use single one line presentation
// better structure. The below, would be the same:

// func (p people) Swap(i, j int) {
// 	temp := p[i]
// 	p[i] = p[j]
// 	p[j] = temp
// }

// Method for printing when string method is used.
func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Age)
}

// =============================
// ByAge implements sort.Interface for []person based on
// the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }

//===================

func main() {
	// Setting the initial values for sorting
	studygroup := people{"Zeno", "John", "Al", "Coco", "Blocko", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Coco", "Jenny"}
	n := []int{7, 6, 4, 8, 9, 19, 32, 23, 7, 14, 20}

	// Sorting groups. Sort needs the interface "Interface"
	fmt.Println("--------------------")
	fmt.Println(studygroup)
	sort.Sort(studygroup)
	// sort.Sort(sort.Reverse(studygroup)) // For reverse order.
	fmt.Println(studygroup)
	fmt.Println("--------------------")

	// Now we are not using a type to attach a method, so we
	// are sending a variable, therefore, we need individual
	// functions that accept the type of our variable (eg. []string)
	// and they also implement the methods Len, Less & Swap.
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	// sort.Sort(sort.StringSlice(s)) //====> same result; it is an interface
	// This interface implements the methods needed in the Interface
	// interface, in order to call the sort method (no need to build them).
	fmt.Println(s)

	fmt.Println("--------------------")

	fmt.Println(n)
	sort.IntSlice(n).Sort()
	// sort.Sort(sort.IntSlice(n)) //====> same result!
	fmt.Println(n)
	fmt.Println("--------------------")

	// Experimenting
	fmt.Println("Experimenting here!")
	t := sort.StringSlice(s)
	fmt.Printf("s is type: %T\nt is type: %T\n", s, t)
	fmt.Println("--------------------")

	// Sorting the complex struct.
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("--------------------")

}

// Check https://godoc.org/sort for more info.

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface

// String() string
// https://golang.org/doc/effective_go.html#printing
