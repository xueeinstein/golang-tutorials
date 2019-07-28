package main

import "fmt"

func main() {
	var a = 100
	var p = &a

	fmt.Println("a = ", a)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	// Changing the value stored in the pointed variable through the pointer
	*p = 2000
	fmt.Println("a (after) = ", a)

	var s = "abcdefg"
	var ps = &s
	fmt.Println("s = ", s)
	fmt.Println("*ps = ", *ps)
	// fmt.Println("*(ps+1) = ", *(ps + 1))
}
