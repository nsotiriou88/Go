package main

import (
	"encoding/json" // For json implementation.
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	Kakaota     int `json:"-"`        // This means it is not exported.
	Kakaota2    int `json:"NEW_NAME"` // It changes the name.
	notExported int // Not exported because it is not with
	// capital letter starting for extraction of data!
}

func main() {
	p1 := person{"James", "Bond", 20, 72, 69, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
