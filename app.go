package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) Area() {
	fmt.Println("Area:", r.width*r.height)
}

func (r *Rectangle) Perimeter() {
	fmt.Println("Perimeter:", 2*(r.width+r.height))
}

func main() {
	var rect Rectangle
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please enter width and height of the rectangle:")
		input, _ := reader.ReadString('\n')
		_, err := fmt.Sscanf(input, "%d %d", &rect.width, &rect.height)
		if err != nil {
			fmt.Println("⚠️ Invalid input")
		} else if rect.width <= 0 || rect.height <= 0 {
			fmt.Println("⚠️ Width and height must be positive integers")
			fmt.Println("Width and height must be positive integers")
		} else {
			rect.Area()
			rect.Perimeter()
			break
		}
	}
}
