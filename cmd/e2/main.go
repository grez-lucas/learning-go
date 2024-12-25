package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grez-lucas/learning-go/c5-functions/e2"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("Currently on working dir:", cwd)
	filename := "c5-functions/e2/input"
	nBytes, err := e2.FileLen(filepath.Join(cwd, filename))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your file has %d bytes!\n", nBytes)
}
