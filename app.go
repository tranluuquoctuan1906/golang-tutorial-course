package main

import "fmt"

func main() {
	s1 := 15
	s2 := 27
	sum := s1 + s2
	fmt.Println("Sum:", sum)
	diff := s1 - s2
	fmt.Println("Difference:", diff)
	prod := s1 * s2
	fmt.Println("Product:", prod)
	quot := float32(s1) / float32(s2)
	fmt.Printf("Quotient: %.2f\n", quot)
	rem := s2 % s1
	fmt.Println("Remainder:", rem)
}
