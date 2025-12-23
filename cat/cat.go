package cat

type Cat struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

func New(name, color string, age int) Cat {
	return Cat{
		Name:  name,
		Color: color,
		Age:   age,
	}
}

func (c Cat) Meow() string {
	return "Meow!"
}
