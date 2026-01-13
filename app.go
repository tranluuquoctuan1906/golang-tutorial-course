package main

import (
	"fmt"
	"golang-tutorial-course/cat"
	"golang-tutorial-course/dog"
)

type Animal interface {
	Speak() string
	GetName() string
}

func makeSound(a Animal) {
	fmt.Printf("The animal's name is: %s\n", a.GetName())
	fmt.Printf("%s says: %s\n", a.GetName(), a.Speak())
}

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
	makeSound(myDog)
	makeSound(myCat)
}
