package e2

import (
	"fmt"
	"os"
)

func FileLen(filename string) (int, error) {
	fmt.Printf("Trying to open file %s\n", filename)

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 2048)

	nBytes, err := f.Read(buf)
	if err != nil {
		return 0, err
	}

	return nBytes, nil
}
