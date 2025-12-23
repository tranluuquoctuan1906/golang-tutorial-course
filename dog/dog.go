package dog

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func New(name, breed string, age int) Dog {
	return Dog{
		Name:  name,
		Breed: breed,
		Age:   age,
	}
}

func (d Dog) Bark() string {
	return "Woof!"
}
