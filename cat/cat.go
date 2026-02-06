package cat

import (
	"errors"
	"strings"
)

type Cat struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

func New(name, color string, age int) (*Cat, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("Name cannot be empty")
	}
	if len(name) > 50 {
		return nil, errors.New("Name cannot exceed 50 characters")
	}
	return &Cat{
		Name:  name,
		Color: color,
		Age:   age,
	}, nil
}

func (c Cat) GetName() string {
	return c.Name
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Eat(food string) string {
	return c.Name + " is eating " + food
}
