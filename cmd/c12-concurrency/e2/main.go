package main

import (
	"fmt"
)

func launchTwoRoutines() (chan int, chan int) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	// goroutine 1
	go func() {
		defer close(ch1)
		for _, num := range nums {
			ch1 <- num
		}
	}()
	// goroutine 2
	go func() {
		defer close(ch2)
		for _, num := range nums {
			ch2 <- num
		}
	}()

	return ch1, ch2
}

func main() {
	ch1, ch2 := launchTwoRoutines()

	for i := 0; i <= 20; i++ {
		select {
		case v1, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Printf("Read %d from goroutine 1!\n", v1)
		case v2, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Printf("Read %d from goroutine 2!\n", v2)
		}
	}
}
