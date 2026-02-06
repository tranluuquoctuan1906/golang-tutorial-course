package main

import (
	"fmt"
	"golang-tutorial-course/cat"
	"golang-tutorial-course/dog"
	"golang-tutorial-course/service"
)

func makeSound(a service.Animal) {
	fmt.Printf("The animal's name is: %s\n", a.GetName())
	fmt.Printf("%s says: %s\n", a.GetName(), a.Speak())
}

func makeEat(a service.AnimalWithEat, food string) {
	fmt.Printf("The animal's name is: %s\n", a.GetName())
	fmt.Printf("%s\n", a.Eat(food))
}

func main() {
	myDog, err := dog.New(
		"Buddy",
		"Golden Retriever",
		3,
	)
	if err != nil {
		panic(err)
	}
	myCat, err := cat.New(
		"Whiskers",
		"Tabby",
		2,
	)
	if err != nil {
		panic(err)
	}
	makeSound(myDog)
	makeEat(myCat, "fish")
}
