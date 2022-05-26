package main

import (
	"fmt"
)

func main() {
	array := []int{1,2,3,4}
	slice := array[:2]
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(cap(slice)) //len of original array
	fmt.Println(len(slice)) //len of slice

	//slice = slice[0:88]
	//fmt.Println(slice) FAILS
	//This fails because the cap of the slice
	// is 4, while this would extend its len
	// to 88
}