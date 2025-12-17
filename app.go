package main

import "fmt"

func sumToN(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumToN(n-1)
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer.")
		fmt.Println(err)
		return
	}
	result := sumToN(n)
	fmt.Printf("The sum of numbers from 1 to %d is %d\n", n, result)
}
