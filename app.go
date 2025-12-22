package main

import "fmt"

func main() {
	name := "World"
	fmt.Printf("%T\n", name)
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", &name)
}
