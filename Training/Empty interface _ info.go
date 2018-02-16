package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}
type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	// No/empty interface
	fmt.Println("==============")
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	shadow := dog{animal{"woof"}, true}
	critters := []interface{}{fido, fifi, shadow}
	fmt.Println(critters)
	fmt.Println("==============")

	// No/empty interface 2nd example
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	cars := []car{prius, tacoma, bmw528}

	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	planes := []plane{boeing747, boeing757, boeing767}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	boats := []boat{sanger, nautique, malibu}

	for key, value := range cars {
		fmt.Println(key, " - ", value)
	}

	fmt.Println()
	for key, value := range planes {
		fmt.Println(key, " - ", value)
	}

	fmt.Println()
	for key, value := range boats {
		fmt.Println(key, " - ", value)
	}
	fmt.Println("==============")
}

// Do not use many (empty) interfaces, as it will get
// complicated to process data/variables afterwards;
// they are different interface and variable types and
// this makes it difficult to work with them with
// conventional functions.

// Also, check "golint" & "go fmt", commonly used with
// "./...", for formating and searching for issues in
// your written code. It will search the whole root folder
// and the rest of the folders in it (all folders).
