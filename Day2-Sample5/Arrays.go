package main

import "fmt"

func main() {
	nums := [6]int{1,2,3,4,5,6} //[size]type{vals}
	fmt.Println(nums)
	slice := nums[1:3] //start incl, end not
	fmt.Println(slice)

	fmt.Println()
	fmt.Println("Note that changing slice changes the array")
	fmt.Println()

	slice[0] = 7
	fmt.Println(nums)
	fmt.Println(slice)
	fmt.Println("changed ind0slice, ind1array")
}