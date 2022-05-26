package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "Linux" :
		fmt.Println("Linux: match found")
	case "Ubuntu" :
		fmt.Println("Ubuntu: match found")
	case "windows" :
		fmt.Println("Windows: match found")
	default:
		fmt.Printf("%s : match not found", os)
	}

}