package main

import "fmt"

func main() {
	a := 5000
	b := &a
	fmt.Printf("%T is the type for %T, %v, %v", b, a, b, a)
}
