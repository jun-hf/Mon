package main

import "fmt"

func main() {
	var a = []int{}

	if a == nil {
		print("is nil")
	}

	fmt.Printf("%+v", a)
}
