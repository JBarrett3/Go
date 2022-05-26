package main

import (
	"fmt"
)

func main() {
	a := make([]int,3,7) 
	lenA := len(a)
	capA := cap(a)
	fmt.Println("length is first: ", lenA)
	fmt.Println("cap is second: ", capA)
}