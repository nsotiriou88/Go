package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 12
	var x2 = "12"
	var y = 12.5645645534758347
	var k rune = 'a' // rune is an alias for int32!
	var l int32 = 'b'
	fmt.Println("============")

	// We widenning the value (adding decimals)
	fmt.Println(y + float64(x)) // Can't use just x!!!
	// (widenning) convertion: int to float64
	fmt.Println("============")
	// We are narrowing the value (losing the decimals)
	fmt.Println(int(y) + x)
	fmt.Println("============")

	// Convertion rune to string
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(string(k))
	fmt.Println(string(l))
	fmt.Println("============")

	// Convertion []bytes(silce of bytes) to string and opposite
	fmt.Println(string([]byte{'a', 'b', 'c'}))
	fmt.Println([]byte("abc"))
	fmt.Println("============")

	// Convertion using strconv
	z, _ := strconv.Atoi(x2)
	z2 := "I have that many: " + strconv.Itoa(x)
	fmt.Println(x + z)
	fmt.Println(z2)
	fmt.Println("============")

	// Assertion examples
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
	fmt.Println("============")

	// var is an interface int
	var val interface{} = 7
	fmt.Println(val.(int) + 6)
	fmt.Printf("Val: %v is of type: %T\n", val.(int), val.(int))
	// fmt.Printf("%T\n", int(val)) // not working, needs assertion
	fmt.Println("============")

}

// Conversion is for values (int to float etc) and Assertion
// is only used for interfaces!!!
