package main

import "fmt"

type sporting interface {
	turnTurbo()
}

type ferrari struct {
	model          string
	turboActivated bool
	currentSpeed   int
}

func (f *ferrari) turnTurbo() {
	f.turboActivated = true
}

func main() {
	f := ferrari{"F40", false, 0}
	f.turnTurbo()

	var s sporting = &ferrari{"F40", false, 0}
	s.turnTurbo()

	fmt.Println(f, s)
}
