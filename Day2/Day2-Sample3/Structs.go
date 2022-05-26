package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1,2}
	v.X = 3
	fmt.Println(v)
	p := &v
	p.X = 5 //Note that p is not derefrenced
	// (*p).X = 5 would also work
	// but *p.X would not work because it 
	//    computs as *(p.X)
	fmt.Println(v)
}