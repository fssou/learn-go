package main

type grade float64

func (g grade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func gradeConcept(g grade) string {
	switch {
	case g.between(9.0, 10.0):
		return "A"
	case g.between(7.0, 8.99):
		return "B"
	case g.between(5.0, 7.99):
		return "C"
	case g.between(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	println(gradeConcept(9.8))
	println(gradeConcept(6.9))
	println(gradeConcept(2.1))
}
