package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	fmt.Println("Encoding from here!")
	p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
	/* NewEncoder returns a pointer to an encoder (*encoder)
	so that next we can use it to Encode (needs *encoder). The
	Stdout opens a writer with a pointer to a file (*file) and therefore
	we can send the data there. */
	fmt.Println("----------------")
	fmt.Println("Decoding from here!")

	// Decoding from here
	var p2 person
	rdr := strings.NewReader(`{"First":"Coco","Last":"Blocko","Age":29,"notExported":20}`)
	// Notice that I added the last string about notExported, though it is not passing,
	// because it is not exported from the struct (not capital).
	er1 := json.NewDecoder(rdr).Decode(&p2) // If you want, you check error.

	fmt.Println(er1)
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
	fmt.Println(p2.notExported) // No value passing here.
}
