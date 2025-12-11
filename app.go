package main

import "fmt"

func main() {
	var fullName, address string
	fmt.Print("Enter your full name: ")
	fmt.Scanln(&fullName)
	fmt.Print("Enter your address: ")
	fmt.Scanln(&address)
	fmt.Printf("Hello, %s! Your address is %s.\n", fullName, address)
}
