package main

import "fmt"

func main() {
	// One way
	greetings := make(map[string]string)
	greetings["Nick"] = "Hello World!"
	fmt.Println(greetings)
	// Second way
	var greeting = map[string]string{ // Or without Var, but change = to :=
		"Kate": "Hi!",
		"Mike": "Hello!"}
	fmt.Println(greeting)
	greeting["Bob"] = "Hey!"
	fmt.Println(greeting)
	delete(greeting, "Bob") // Deleting a key
	fmt.Println(greeting)

	fmt.Println("-----------")

	// New about checking Maps
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}
	fmt.Println(myGreeting)
	// delete(myGreeting, 2) // Use it do delete and run again
	// If and checking the Type of variables
	if val, exists := myGreeting[2]; exists { // The Comma ok idiom
		// Go determines the type of the variable, and because we placed
		// exists after the ";" in the if statement, it is a bool.
		fmt.Println("That value exists.")
		fmt.Printf("val is %T and exists is %T.\n", val, exists)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Printf("val is %T and exists is %T.\n", val, exists)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)

	fmt.Println("-----------")
	// Multi-dimentional Maps
	multiMap := map[int]map[string]string{
		1: {
			"H":       "Hydrogen",
			"State":   "Gas",
			"Protons": "One",
		},
		2: {
			"He":    "Helium",
			"State": "Gas",
		},
	}
	fmt.Println(multiMap)
	fmt.Println(multiMap[1]["State"])
	fmt.Println()

	for k, l := range multiMap {
		fmt.Println(k)
		fmt.Println(l)
		for i, j := range l {
			fmt.Println(i)
			fmt.Println(j)
		}
	}
}
