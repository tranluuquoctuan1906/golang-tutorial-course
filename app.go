package main

type User struct {
	name  string
	age   int
	phone string
}

func showUserInfo(u *User) {
	println("Name:", u.name)
	println("Age:", u.age)
	println("Phone:", u.phone)
}

func main() {
	user := User{name: "Alice", age: 30, phone: "123-456-7890"}
	showUserInfo(&user)

	mina := User{name: "Mina", age: 25, phone: "098-765-4321"}
	showUserInfo(&mina)
}
