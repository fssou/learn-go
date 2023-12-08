package main

import "fmt"

type course struct {
	name string
}

func main() {
	var any interface{} = 3
	fmt.Println(any)

	any = "string"
	fmt.Println(any)

	any = true
	fmt.Println(any)

	any = course{"Golang: Exploring interfaces"}
	fmt.Println(any)
}
