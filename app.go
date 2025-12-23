package main

import "encoding/json"

type User struct {
	Name  string
	Age   int
	Phone string
}

func (u *User) showInfo() {
	println("Name:", u.Name)
	println("Age:", u.Age)
	println("Phone:", u.Phone)
}

func (u *User) resetUser() {
	u.Name = ""
	u.Age = 0
	u.Phone = ""
}

func main() {
	user := User{Name: "Alice", Age: 30, Phone: "123-456-7890"}
	output, _ := json.Marshal(user)
	println(string(output))
}
