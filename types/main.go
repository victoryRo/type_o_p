package main

import "fmt"

type tutorial struct {
	name string
}

func (t tutorial) Print() {
	fmt.Printf("%+v\n", t)
}

// declarar alias
type alias = tutorial

// definicion de tipo
type newTutorial tutorial

func main() {
	t := newTutorial{"Go"}
	// t.Print()
	fmt.Printf("the type is: %T\n", t)

	fmt.Println("----------------------------------------")

	var v newBool = false
	fmt.Println(v.String())
}

type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}
