package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	// Second part of Arrays with binary array.
	var y [256]byte
	fmt.Println(len(y))
	fmt.Println(y[42])
	for k := 0; k < 256; k++ {
		y[k] = byte(k)
	}

	for k, v := range y { // Notice the way range works here.
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if k > 50 {
			break
		}
	}
}
