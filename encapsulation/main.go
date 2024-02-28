package main

import (
	"fmt"
	"github.com/victoryRo/type_o_p/encapsulation/course"
)

func main() {
	Cs := course.New("Advanced golang", 0, false)
	usersIds := []uint{
		11, 22, 33,
	}
	Cs.SetUsersIDs(usersIds)

	class := map[uint]string{
		1: "Introducción",
		2: "Structures",
		3: "Maps",
		4: "Interfaces",
	}
	Cs.SetClasses(class)

	// I instantiate the course type and access its method
	Cs.PrintClasses()
	//Cs.ChangePrice(42.01)
	//fmt.Println(Cs.price)
	fmt.Printf("%+v\n", Cs)
	Cs.SetName("Go professional programming")
	fmt.Printf("%+v\n", Cs.Name())
	fmt.Printf("%+v\n", Cs.UsersIDs())
}
