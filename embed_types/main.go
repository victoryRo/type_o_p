package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hello Peson")
}

type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age, children uint) Human {
	return Human{age, children}
}

type Employee struct {
	Person
	Human
	Salary float64
}

func NewEmployee(name string, age, children uint, salary float64) Employee {
	return Employee{
		Person: Person{Name: name, Age: age},
		Human:  Human{Age: age, Children: children},
		Salary: salary,
	}
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func (e Employee) Greet() {
	fmt.Println("Hello Employee")
}

func main() {
	e := NewEmployee("Victor", 35, 3, 1000000)
	e.Payroll()
	fmt.Println(e.Name)
	fmt.Println(e.Human.Age)
	e.Greet()
	e.Person.Greet()
}
