package main

import "fmt"

func logical(b1, b2 bool) (bool, bool, bool) {
	and := b1 && b2
	xor := b1 != b2 // ou exclusivo
	or := b1 || b2

	return and, xor, or
}

func main() {
	and, xor, or := logical(true, true)
	fmt.Printf(
		"AND: %t, XOR: %t, OR: %t, NOT OR: %t",
		and, xor, or, !or,
	)
}
