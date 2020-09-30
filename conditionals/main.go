package main

import "fmt"

func main() {
	x := 4
	y := 5
	if x > y {
		fmt.Printf("%d is greater than %d", x, y)
	} else {
		fmt.Printf("%d is not greater than %d\n", x, y)
	}

	color := "red"
	if color == "red" {
		fmt.Printf("Trump sarkar because the color is %s", color)
	} else {
		fmt.Printf("Clinton sarkar because the color is %s", color)
	}
}
