package main

import (
	"fmt"
	"reflect"
)

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastname string
}

func (p person) toString() string {
	return fmt.Sprintf("%s %s", p.name, p.lastname)
}

type product struct {
	name  string
	price float64
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Printf("[%s] - %s\n", reflect.TypeOf(x), x.toString())
}

func main() {
	var thing printable = product{"Pen", 2.99}
	print(thing)

	thing = person{"John", "Doe"}
	print(thing)

	// Interface is a type
	var thing2 printable = product{"Pen", 2.99}
	print(thing2)

	thing2 = person{"John", "Doe"}
	print(thing2)

	// Interface is a type
	var thing3 printable = product{"Pen", 2.99}
	print(thing3)

	thing3 = person{"John", "Doe"}
	print(thing3)
}
