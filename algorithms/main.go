package main

import (
	"fmt"
	"strings"
)

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

//Some helpers for optimized solution fibonacci using dynamic programming
var fibonacciStore map[int]int

func init() {
	fibonacciStore = make(map[int]int)
}

func fibonacciOptimized(x int) int {
	if x <= 2 {
		return 1
	}
	if fibonacciStore[x] > 0 {
		return fibonacciStore[x]
	}
	result := fibonacciOptimized(x-1) + fibonacciOptimized(x-2)
	fibonacciStore[x] = result
	return result
}

func reverse(str string) string {
	runes := []rune(str)
	if len(runes) <= 1 {
		return str
	}
	return reverse(str[1:]) + str[:1]
}

func checkPalindrome(str string) bool {
	reversed := reverse(str)
	if reversed == str {
		return true
	} else {
		return false
	}
}

type callbackType func(int) bool

func someRecursive(arr []int, callback callbackType) bool {
	result := false
	first := arr[0]
	result = callback(first)
	if result {
		return result
	}
	if len(arr) == 1 {
		return result
	}
	return someRecursive(arr[1:], callback) || result
}

func isOdd(num int) bool {
	result := num%2 != 0
	return result
}

func capitalizeFirst(demoStrings []string) []string {
	str := demoStrings[0]
	result := strings.ToUpper(str[0:1]) + str[1:]
	results := []string{result}
	if len(demoStrings) == 1 {
		return results
	}
	return append(results, capitalizeFirst(demoStrings[1:])...)
}

func capitalizeAll(demoStrings []string) []string {
	str := demoStrings[0]
	result := strings.ToTitle(str)
	fmt.Println(result)
	results := []string{result}
	if len(demoStrings) == 1 {
		return results
	}
	return append(results, capitalizeAll(demoStrings[1:])...)
}

func main() {
	// arr := []int{44, 26, 79, 31, 12, 85, 0, 1, 2, 223}
	// newArr := []int{8, 1, 2, 3, 4, 5, 6, 7, 9}
	// demoStrings := []string{"hello", "taco", "bell", "sheldon cooper", "hello"}
	// names := []string{"Sourabh", "Anshul", "Max", "Colt", "Brad", "Hughie", "Starlight", "William"}
	// sortedNames := []string{"anshul", "brad", "colt", "hughie", "max", "sourabh", "starlight", "william"}
	// sortedNames := []int{1, 4, 9, 15, 19, 23, 28}
	// longString := "hello sourabh how are you, how is your dad?"
	// shortString := "how"
	// fmt.Println(factorial(5))
	// fmt.Println(inefficientFactorial(5))
	// fmt.Println(collectOddValues(arr))
	// fmt.Println(power(2, 0))
	// fmt.Println(productOfArray(arr))
	// fmt.Println(sumOfAllTill0(11))
	fmt.Println(fibonacciOptimized(44))
	// fmt.Println(fibonacci(44))
	// fmt.Println(reverse("Hello"))
	// fmt.Println(checkPalindrome("totatatot"))
	// fmt.Println(someRecursive(arr, isOdd)) recursive check using callbacks
	// fmt.Println(capitalizeFirst(demoStrings))
	// fmt.Println(capitalizeAll(demoStrings))
	// fmt.Println(search.LinearSearch(names, "Colt"))
	// fmt.Println(search.BinarySearch(sortedNames, 199))
	// fmt.Println(search.NaiveStringSearch(longString, shortString))
	// fmt.Println(sort.Bubble(arr))
	// fmt.Println(sort.Selection(arr))
	// fmt.Println(sort.Insertion(arr))
	// fmt.Println(sort.Merge(arr))
	// fmt.Println(sort.Quick(arr))
	// fmt.Println(sort.Radix(arr))
}
