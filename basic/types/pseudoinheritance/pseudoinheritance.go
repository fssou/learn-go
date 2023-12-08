package main

import "fmt"

type car struct {
	name  string
	speed int
}

type ferrari struct {
	car
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.speed = 0
	f.turboOn = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.name, f.turboOn)
}
