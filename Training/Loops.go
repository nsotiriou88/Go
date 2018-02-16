package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Println(i, "-", j)
		}
	}
	i := 17
	for i > 10 {
		if i == 14 {
			i--
			continue
		}
		if i == 12 {
			break
		}
		fmt.Println(i)
		i--
	}
	/* we can use continue to go back to the loop
	or break to immediately get out of it. In this
	way, we create Do & While with only the For in go. */
	for i := 50; i <= 120; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
		// with simple quote ' , we take the value of rune
	}
	fmt.Println('i') // i is 105 in UTF-8
}
