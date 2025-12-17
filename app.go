package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Printf("Enter a positive integer: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Please enter a non-negative integer.")
		return
	}
	result := fibonacci(n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, result)
}
