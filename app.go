package main

import (
	"strconv"
	"strings"
)

func main() {
	var xhtml strings.Builder
	for i := 1; i <= 100; i += 2 {
		xhtml.WriteString(strconv.Itoa(i))
		if (i+1)%3 == 0 {
			xhtml.WriteString("\n")
		} else if i != 99 {
			xhtml.WriteString(",")
		}
	}
	println(xhtml.String())
}
