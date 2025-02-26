package main

import (
	"fmt"

	c6 "github.com/grez-lucas/learning-go/chapters/c6-pointers/e2"
)

func main() {
	sSlice := []string{"first", "second", "third"}
	s := "fourth"

	fmt.Println("Slice before UpdateSlice: ", sSlice)
	c6.UpdateSlice(sSlice, s)
	fmt.Println("Slice after UpdateSlice: ", sSlice)

	sSlice = []string{"first", "second", "third"}

	fmt.Println("Slice before GrowSlice: ", sSlice)
	c6.GrowSlice(sSlice, s)
	fmt.Println("Slice after UpdateSlice: ", sSlice)
}
