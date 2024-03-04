package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type GreeterByer interface {
	Greeter
	Byer
}

//---------------------------------------------------------------- interface

// Person implement Greeter interface
// Person implement Byer interface
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello i am " + p.Name)
}

func (p Person) Bye() {
	fmt.Println("Bye " + p.Name)
}

func (p Person) String() string {
	return fmt.Sprintf("Hello i am from stringer interface Name: %s", p.Name)
}

//---------------------------------------------------------------- implement interface

// Text implement Greeter interface
// Text implement Byer interface
type Text string

func (t Text) Greet() {
	fmt.Printf("Greet from Text type: Hello %s\n", t)
}

func (t Text) Bye() {
	fmt.Printf("Bye from Text type: Bye %s\n", t)
}

//---------------------------------------------------------------- implement interface

func main() {
	//firstExample()
	//secondExample()
	p := Person{"Susana"}
	fmt.Println(p)
}

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

// repetitive code
func greeterInterface(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
	}
}

// repetitive code
func byerInterface(bs ...Byer) {
	for _, b := range bs {
		b.Bye()
	}
}

func secondExample() {
	p := Person{"Pedro"}
	var t Text = "Elizabeth"
	All(p, t)

	fmt.Println("----------------------------------------------------")

	//greeterInterface(p, t)
	//byerInterface(p, t)
}

func firstExample() {
	var g Greeter = Person{"Victory"}
	g.Greet()
	var gr Greeter = Text("Maria")
	gr.Greet()

	fmt.Println("----------------------------------------------------")

	var b Byer = Person{"Mariana"}
	b.Bye()
	var br Byer = Text("Vacations")
	br.Bye()
}
