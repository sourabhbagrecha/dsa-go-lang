package main

import "fmt"

func greetings(name string) string {
	return "Hello " + name + ", hope you are doing well!"
}

func main() {
	fmt.Println("Running inside main!")
	fmt.Println(greetings("Sourabh Bagrecha"))
}
