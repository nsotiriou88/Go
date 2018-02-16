package main

import (
	"fmt"
)

func main() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T\n", letter)
	let := rune("A"[0])
	// note the difference with runes '' and strings "".
	fmt.Println(let)
	fmt.Printf("%T\n", let)

	// Remove the comment to run the program. It prints the whole book

	/* res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs) */
}
