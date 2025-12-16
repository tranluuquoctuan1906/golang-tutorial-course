package main

import (
	"slices"
	"strconv"
	"strings"
)

func main() {
	var xhtml strings.Builder
	numbers := []int{6, 48, 75, 89}
	for i := 1; i <= 100; i++ {
		if slices.Contains(numbers, i) {
			continue
		}
		xhtml.WriteString(strconv.Itoa(i))
		if i != 100 {
			xhtml.WriteString(",")
		}
	}
	println(xhtml.String())
}
