package main

import "fmt"

func updateVariable(name *string) {
	*name = "Tran Thi B"
	fmt.Println("After changing value of name variable")
	fmt.Printf("Type of name variable: %T\n", name)
	fmt.Printf("Value of name variable: %v\n", *name)
	fmt.Printf("Address of name variable: %p\n", name)
}

func main() {
	name := "Nguyen Van A"
	fmt.Println("Information of variable")
	fmt.Printf("Type of name variable: %T\n", name)
	fmt.Printf("Value of name variable: %v\n", name)
	fmt.Printf("Address of name variable: %p\n", &name)
	updateVariable(&name)
}
