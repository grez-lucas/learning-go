package main

import (
	"fmt"

	"github.com/grez-lucas/learning-go/c5-functions/e3"
)

func main() {
	helloPrefix := e3.Prefixer("Hello")

	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))

}
