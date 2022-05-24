package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Please provide positive number: ")
	var input float64
	fmt.Scanln(&input);
	fmt.Println("Math Sqrt is ", math.Sqrt((input)))
	fmt.Println("My Sqrt is equal to or just less than ", sqrt(input))
}

func sqrt(input float64) float64 {
	var x float64
	for (x * x) < input {
		x++
	}
	return x
}