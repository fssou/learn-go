package main

type sporting interface {
	turnTurbo()
}

type luxurious interface {
	autoPark()
}

type sportingLuxurious interface {
	sporting
	luxurious
}

type bmw struct{}

func (b bmw) turnTurbo() {
	println("Turbo...")
}

func (b bmw) autoPark() {
	println("Parking...")
}

func main() {
	var b sportingLuxurious = bmw{}
	b.turnTurbo()
	b.autoPark()
}
