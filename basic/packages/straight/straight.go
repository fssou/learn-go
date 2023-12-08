package main

import (
	"math"
)

type point struct {
	x float64
	y float64
}

func sides(a, b point) (cx, cy float64) {
	cx = math.Abs(a.x - b.x)
	cy = math.Abs(a.y - b.y)
	return
}

func distance(a, b point) float64 {
	cx, cy := sides(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
