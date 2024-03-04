package main

import "fmt"

type Storage interface {
	Get() string
	Set(string)
}

//---------------------------------------------------------- interface

type User struct {
	Name string
}

// NewUser constructor
func NewUser(name string) *User {
	return &User{name}
}

func (u *User) Set(s string) {
	u.Name = s
}

func (u *User) Get() string {
	return "My name is: " + u.Name
}

func Exec(s Storage, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	u := NewUser("Victor")
	Exec(u, "Marjean")
}
