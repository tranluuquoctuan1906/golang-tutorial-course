package main

func main() {
	score := 9

	if score > 8 {
		println("Good score")
	} else if score <= 8 && score >= 6 {
		println("Average score")
	} else if score < 6 && score >= 4 {
		println("Below average score")
	} else {
		println("Poor score")
	}
}
