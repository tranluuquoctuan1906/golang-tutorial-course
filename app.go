package main

import (
	"golang-tutorial-course/cat"
	"golang-tutorial-course/dog"
)

func main() {
	myDog := dog.New(
		"Buddy",
		"Golden Retriever",
		3,
	)
	myCat := cat.New(
		"Whiskers",
		"Tabby",
		2,
	)
	println(myDog.Bark())
	println(myCat.Meow())
}
