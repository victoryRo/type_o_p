package main

import (
	"fmt"
	"strings"
)

// ------------------------------------------------------------ empty interface
func testEmptyInterface(value interface{}) {
	fmt.Printf("empty interface type: %T value %v\n", value, value)
}
func empty() {
	testEmptyInterface(35)
	testEmptyInterface(35.04)
	testEmptyInterface("Hello love")
	testEmptyInterface(true)
	testEmptyInterface(-99)
	testEmptyInterface(map[int]string{})
	testEmptyInterface([4]string{})
}

func main() {
	//empty()
	//typeAssertions("mariana bella")
	typeAssertMany("hello how are you")
	typeAssertMany(99)
	typeAssertMany(77.89)
}

// ------------------------------------------------------------ type assertion

func typeAssertions(val interface{}) {
	v, ok := val.(string)
	if ok {
		fmt.Printf("%v\n", strings.ToUpper(v))
	}
}

func typeAssertMany(val interface{}) {
	switch v := val.(type) {
	case string:
		fmt.Printf("%v\n", strings.ToUpper(v))
	case int:
		fmt.Printf("value multiplication %d\n", v*2)
	case float64:
		fmt.Printf("float64 value %2.f\n", v)
	case bool:
		fmt.Printf("bool type %t\n", v)
	default:
		fmt.Println("No found type")
	}
}
