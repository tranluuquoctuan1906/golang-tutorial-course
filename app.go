package main

func sum(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func countdown(n int) {
	if n <= 0 {
		println("Done!")
		return
	}
	println(n)
	countdown(n - 1)
}

func main() {
	sum, diff := sum(3, 5)
	println("The sum is:", sum)
	println("The difference is:", diff)
	countdown(10)
}
