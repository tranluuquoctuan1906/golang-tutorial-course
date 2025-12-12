package main

func main() {
	score := 9

	switch {
	case score > 8:
		println("Good score")
	case score > 6:
		println("Average score")
	case score > 4:
		println("Below average score")
	default:
		println("Poor score")
	}
}
