package main

type User struct {
	Name string
	Age  int
}

func (u *User) changeName(name string) {
	u.Name = name
}

func Change(a *User, b *User) {
	temp := a.Name
	a.changeName(b.Name)
	b.changeName(temp)
}
