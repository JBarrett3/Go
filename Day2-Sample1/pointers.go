package main

import "fmt"

func main() {
	i := 5
	p := &i
	fmt.Println("p points to val", *p)
	fmt.Println("i has val", i)
	fmt.Println("p points to address", p)
	fmt.Println("i has address", &i)
}