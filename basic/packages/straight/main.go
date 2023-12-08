package main

import "fmt"

func main() {
	p1 := point{2.0, 2.0}
	p2 := point{2.0, 4.0}
	fmt.Println(sides(p1, p2))
	fmt.Println(distance(p1, p2))
}
