package main

import "fmt"

type Course struct {
	Name     string
	Price    float64
	IsFree   bool
	UsersIDs []uint
	Classes  map[uint]string
}

// PrintClasses method for Course
func (c *Course) PrintClasses() {
	text := "The classes are: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

// ChangePrice updates the price of the course
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
