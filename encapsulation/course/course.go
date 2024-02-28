package course

import "fmt"

type Course struct {
	name     string
	price    float64
	isFree   bool
	usersIDs []uint
	classes  map[uint]string
}

func New(name string, price float64, isFree bool) *Course {
	if price == 0 {
		price = 30
	}

	return &Course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

func (c *Course) SetName(name string) { c.name = name }

func (c *Course) Name() string { return c.name }

func (c *Course) SetPrice(price float64) { c.price = price }

func (c *Course) Price() float64 { return c.price }

func (c *Course) SetIsFree(isFree bool) { c.isFree = isFree }

func (c *Course) IsFree() bool { return c.isFree }

func (c *Course) SetUsersIDs(usersIDs []uint) { c.usersIDs = usersIDs }

func (c *Course) UsersIDs() []uint { return c.usersIDs }

func (c *Course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

// PrintClasses method for course
func (c *Course) PrintClasses() {
	text := "The classes are: "
	for _, class := range c.classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
