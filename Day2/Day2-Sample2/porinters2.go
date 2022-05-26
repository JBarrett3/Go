package main

import "fmt"

func main() {
	i := 5
	fmt.Println("val of i is", i)
	p := &i
	*p = 9
	fmt.Println("val of i is now ", i)
}