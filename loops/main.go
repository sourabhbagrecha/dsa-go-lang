package main

import "fmt"

func main() {
	// long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// short method
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// fizzbuzz challenge
	// loop through 1 to 100
	for i := 1; i < 100; i++ {
		// if multiple of 15 replace with fizzbuzz
		if i%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
	// if multiple of 5 replace with buzz
	// if multiple of 3 replace with fizz

}
