package main

import "fmt"

func main() {
	for x := 0; x < 5; x++ {
		fmt.Println(x) //Remember that x is contained only in this loop
	}				   //And note that y and z are accessible
	fmt.Println("OR")
	y := 0
	for ; y < 5 ; y++ {
		fmt.Println(y)
	}
	fmt.Println("OR")
	z := 0
	for z < 5 {
		fmt.Println(z)
		z++
	}
	//fmt.Println(x,y,z) produces undefined variable error for x
}