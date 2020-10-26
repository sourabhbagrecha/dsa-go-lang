package main

import "fmt"

func main() {
	// arr := []int{33, 43, 47, 69, 67}

	// for i, ele := range arr {
	// 	fmt.Println(ele, i)
	// }

	//get sum
	// sum := 0
	// for _, ele := range arr {
	// 	sum += ele
	// }
	// fmt.Println("Sum: ", sum)

	user := map[string]string{"email": "sourabhbagrecha@gmail.com", "name": "Sourabh Bagrecha", "password": "98123hbr1i42rh34bt8234hbi4h"}
	for key, value := range user {
		fmt.Print(key, ": ", value, "\n")
	}

}
