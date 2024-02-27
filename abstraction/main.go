package main

import "fmt"

func main() {
	methodsGolang()
	instancesOfTypeCourse()
}

func methodsGolang() {
	Cs1 := Course{
		Name:   "Go from scratch",
		Price:  12.34,
		IsFree: false,
		UsersIDs: []uint{
			12, 56, 89,
		},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Structures",
			3: "Maps",
			4: "Interfaces",
		},
	}

	// I instantiate the course type and access its method
	Cs1.PrintClasses()
	Cs1.ChangePrice(30.45)
	fmt.Println(Cs1.Price)
}

func instancesOfTypeCourse() {
	// different ways to instantiate the Courses type
	Cs1 := Course{
		Name:   "Go from scratch",
		Price:  12.34,
		IsFree: false,
		UsersIDs: []uint{
			12, 56, 89,
		},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Structures",
			3: "Maps",
		},
	}

	Cs2 := Course{Name: "CSS from scratch", IsFree: true}

	Js := Course{}
	Js.Name = "Javascript from scratch"
	Js.Price = 11.22
	Js.IsFree = true
	Js.UsersIDs = []uint{1, 2, 3}
	Js.Classes = map[uint]string{
		1: "Functions",
		2: "Types",
		3: "Export Import",
	}

	fmt.Println(Cs1.Name)
	fmt.Printf("%+v\n", Cs1)
	fmt.Println(Cs2.Name, Cs2.IsFree)
	fmt.Printf("%+v\n", Cs2)
	fmt.Printf("%+v\n", Js)
}
