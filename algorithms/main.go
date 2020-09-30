package main

import "fmt"

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}

func inefficientFactorial(num int) int {
	if num <= 1 {
		return 1
	}
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	return factorial
}

func collectOddValues(arr []int) []int {
	var newArr []int
	if len(arr) == 0 {
		return newArr
	}
	if arr[0]%2 != 0 {
		fmt.Println("hey", arr[0])
		newArr = append(newArr, arr[0])
	}
	newArr = append(newArr, collectOddValues(arr[1:])...)
	return newArr
}

func power(x int, y int) int {
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}
	return x * power(x, y-1)
}

func productOfArray(arr []int) int {
	if len(arr) == 0 {
		return 1
	}
	return arr[0] * productOfArray(arr[1:])
}

func sumOfAllTill0(x int) int {
	if x == 0 {
		return 0
	}
	return x + sumOfAllTill0(x-1)
}

func fibonacci(x int) int {
	if x <= 2 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func reverse(str string) string {

}

func main() {
	// arr := []int{1, 2, 3, 4, 5, 20}
	// fmt.Println(factorial(5))
	// fmt.Println(inefficientFactorial(5))
	// fmt.Println(collectOddValues(arr))
	// fmt.Println(power(2, 0))
	// fmt.Println(productOfArray(arr))
	// fmt.Println(sumOfAllTill0(11))
	fmt.Println(fibonacci(4))
}
