package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
	name string
	age  int
}

func main() {
	c := cat{name: "Oreo", age: 9}
	fmt.Println(c.Speak())
	// c.Greeting()
	fmt.Println(c)
}

func (c cat) Speak() string {
	return "Purrrr Meow"
}

func (c cat) String() string {
	return fmt.Sprintf("%v (%v yesr old)", c.name, c.age)
}

// func (c cat) Greeting() {
// 	fmt.Println("Meow, mmmmeeeeeoooooowwwww!!!")
// }
