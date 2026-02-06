package dog

import (
	"errors"
	"strings"
)

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func New(name, breed string, age int) (*Dog, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("Name cannot be empty")
	}
	if len(name) > 50 {
		return nil, errors.New("Name cannot exceed 50 characters")
	}
	return &Dog{
		Name:  name,
		Breed: breed,
		Age:   age,
	}, nil
}

func (d Dog) GetName() string {
	return d.Name
}

func (d Dog) Speak() string {
	return "Woof!"
}
