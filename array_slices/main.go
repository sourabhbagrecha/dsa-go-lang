package main

import "fmt"

func main() {
	// var fruits [2]string
	// fruits[0] = "apple"
	// fruits[1] = "mango"

	// shorthand
	// fruits := [2]string{"Apple", "Mango"}

	// dynamic length
	fruits := []string{"apple", "mango", "orange"}

	fmt.Println(fruits[2])
}
