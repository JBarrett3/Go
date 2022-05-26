package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v1 := Vertex{}
	fmt.Println("empty produces defaults", v1)
	v2 := Vertex{X : 1, Y: 2}
	fmt.Println("Colons assign by name,", v2)
	v3 := Vertex{1,2}
	fmt.Println("nameless assigns in order", v3)
}