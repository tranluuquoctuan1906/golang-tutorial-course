package service

type Animal interface {
	Speak() string
	GetName() string
}

type AnimalWithEat interface {
	Animal
	Eat(food string) string
}
