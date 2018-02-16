package main

import (
	"fmt"
	"log"
	"os"
	// "errors" // For creating custom/specific errors
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	// This is for printing to a file log, instead of screen.
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		log.Println("err happened", err)
		// log.Fatalln(err)
		// panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard
error and prints the date and time of each logged message ... the Fatal
functions call os.Exit(1) after writing the log message ... the Panic
functions call panic after writing the log message.
*/

// log.Println calls Output to print to the standard logger. Arguments
// are handled in the manner of fmt.Println.

//
// ======= Example with custom errors ======

// func main() {
// 	_, err := Sqrt(-10)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		// First step for creating custom errors: errors.New or Errorf for more info
// 		return 0, errors.New("norgate math: square root of negative number")
// 		// return 0, fmt.Errorf("norgate math: square root of negative number: %v", f)
// 	}
// 	// implementation
// 	return 42, nil
// }

//
// ========== Create variables for errors ==========

// var ErrNorgateMath = errors.New("norgate math: square root of negative number")

// func main() {
// 	fmt.Printf("%T\n", ErrNorgateMath)
// 	_, err := Sqrt(-10)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, ErrNorgateMath
// 	}
// 	// implementation
// 	return 42, nil
// }

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go

// =========== Custom type Error (struct) ============
//

// type NorgateMathError struct {
// 	lat, long string
// 	err       error
// }

// func (n *NorgateMathError) Error() string {
// 	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
// }

// func main() {
// 	_, err := Sqrt(-10.23)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
// 		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
// 	}
// 	// implementation
// 	return 42, nil
// }

// see use of structs with error type in standard library:
// http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
