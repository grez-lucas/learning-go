package main

import (
	"fmt"

	c6 "github.com/grez-lucas/learning-go/chapters/c6-pointers/e1"
)

func main() {
	p := c6.MakePerson("Lucas", "Grez", 24)
	pp := c6.MakePersonPointer("Lucas", "Grez", 24)

	fmt.Println("Person: ", p)
	fmt.Println("Person Pointer: ", pp)
}
