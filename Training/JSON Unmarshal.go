package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"Ηλικία"` // If you change the tag name here,
	// it won't accept input with the tag Age.
}

func main() {
	var p1 person // Initialise to 0 and space.
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "Ηλικία":20}`)
	check := json.Unmarshal(bs, &p1) // Need to point the variable. Returns
	// error(if exists).

	fmt.Println("--------------")
	fmt.Println(check) // <nil> if ok, and error message if anything wrong.
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age) // Check here we use the normal name of value.
	fmt.Printf("%T \n", p1)
}
