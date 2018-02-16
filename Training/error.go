package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Negative square root")
	}
	return math.Pow(x, 0.5), nil
}

func main() {
	i := 0
	for i < 10 {
		k, err := sqrt(-7)
		if err != nil {
			fmt.Println(err)
		}
		println(k)
		i++
	}
}
