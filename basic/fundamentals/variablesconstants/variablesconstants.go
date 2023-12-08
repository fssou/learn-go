package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var radius = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(radius, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
		c = 3
	)

	var (
		d = 3
		e = 4
	)

	fmt.Println(a, b, c, d, e)

	var f, g bool = true, false
	fmt.Println(f, g)

	h, i, j := 2, false, "epa!"
	fmt.Println(h, i, j)
}
