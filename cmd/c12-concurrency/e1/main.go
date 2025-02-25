package main

import (
	"fmt"
	"sync"
)

func processAndGather() {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(2)
	for i := 0; i < 2; i++ {
		// Launch two goroutines that write to chan
		go func() {
			defer wg.Done()
			for _, num := range nums {
				ch <- num
			}
		}()
	}

	// Go routine to clean up shared chan
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Read Go routine
	for num := range ch {
		fmt.Printf("Read num from chan: %d\n", num)
	}
}

func main() {
	processAndGather()
}
