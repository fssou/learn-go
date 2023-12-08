package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastname string
}

func (p person) fullName() string {
	return p.name + " " + p.lastname
}

func (p *person) setFullName(fullname string) {
	names := strings.Split(fullname, " ")
	p.name = names[0]
	p.lastname = names[1]
}

func main() {
	p1 := person{"João", "Pedro"}
	fmt.Println("Antes", p1.fullName())

	p1.setFullName("Maria Junqueira")
	fmt.Println("Depois", p1.fullName())
}
