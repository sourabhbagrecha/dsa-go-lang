package main

import "fmt"

// Person type
type Person struct {
	// alternate syntax
	// firstName string
	// lastName  string
	// gender    string
	// city      string
	// age       int

	firstName, lastName, gender, city string
	age                               int
}

func greet(p Person) string {
	return "Hello " + p.firstName + " " + p.lastName
}

// value Reciever
func (p Person) greetAsValueReceiver() string {
	return "Hello " + p.firstName + " " + p.lastName
}

// pointer Reciever
func (p *Person) gettingMarried(spouseLastName string) {
	if p.gender == "Female" {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Alexa", lastName: "Ryan", age: 21, gender: "Female", city: "Udaipur"}

	// alternate
	// person1 := Person{"Sourabh", "Bagrecha", "Male", "Udaipur", 21}
	// fmt.Println(person1.firstName)

	// fmt.Println(greet(person1))
	person1.gettingMarried("Greer")

	fmt.Println(person1.greetAsValueReceiver())
}
