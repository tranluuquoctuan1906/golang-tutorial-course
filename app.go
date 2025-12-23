package main

type User struct {
	name  string
	age   int
	phone string
}

func (u *User) showInfo() {
	println("Name:", u.name)
	println("Age:", u.age)
	println("Phone:", u.phone)
}

func (u *User) resetUser() {
	u.name = ""
	u.age = 0
	u.phone = ""
}

func main() {
	user := User{name: "Alice", age: 30, phone: "123-456-7890"}
	user.showInfo()
	user.resetUser()
	user.showInfo()
}
