package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fullName string
	fmt.Print("Enter your full name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fullName = scanner.Text()
	}
	fmt.Printf("Hello, %s!\n", fullName)
}
