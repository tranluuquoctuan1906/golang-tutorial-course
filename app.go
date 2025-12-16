package main

import "strconv"

func main() {
	xhtml := ""
	numbers := []int{6, 48, 75, 89}
	for i := 1; i <= 100; i++ {
		skip := false
		for _, num := range numbers {
			if i == num {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		xhtml += strconv.Itoa(i)
		if i != 100 {
			xhtml += ","
		}
	}
	println(xhtml)
}
