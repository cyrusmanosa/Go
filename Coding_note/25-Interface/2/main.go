package main

import (
	"fmt"
)

type Member interface {
	GetName() string
	GetAge() int
}

type Robot struct {
	name  string
	age   int
	power int
}

func (r *Robot) Work() {
}

func (r *Robot) GetName() string {
	return r.name
}
func (r *Robot) GetAge() int {
	return r.age
}

type Car struct {
	name     string
	age      int
	odometer int
}

func (r *Car) Run() {
}

func (r *Car) GetName() string {
	return r.name
}
func (r *Car) GetAge() int {
	return r.age
}

func Cleaning(m Member) {
	fmt.Printf("Consumer Name:%s, Age:%d\n", m.GetName(), m.GetAge())
}

func main() {
	r := &Robot{
		name:  "GUMDAM",
		age:   1,
		power: 100,
	}
	c := &Car{
		name:     "BENZ",
		age:      2,
		odometer: 100,
	}
	Cleaning(c)
	Cleaning(r)
}
