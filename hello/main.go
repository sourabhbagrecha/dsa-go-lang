package main

import (
	"fmt"
)

var count int = 10

func main() {
	for i := 0; i < count; i++ {
		fmt.Println("hello", i)
	}
}
