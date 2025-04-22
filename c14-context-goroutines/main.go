package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var ErrSomethingWentWrong = errors.New("something critical went wrong in processing")

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	numOfWorkers := 2
	errChan := make(chan error, numOfWorkers)
	respChan := make(chan string, numOfWorkers)
	var wg sync.WaitGroup
	wg.Add(numOfWorkers)

	go getRandomStatus(ctx, &wg, respChan, errChan)
	go getDelay(ctx, &wg, respChan, errChan)

	// Listen for errors and results
loop:
	for {
		select {
		case err := <-errChan:
			fmt.Printf("Received an error: %v\n", err)
			cancelFunc()
		case res := <-respChan:
			fmt.Printf("Received a result: %v\n", res)
		case <-ctx.Done():
			fmt.Println("\nContext already done, exiting channel listener.")
			break loop
		}
	}

	fmt.Println("Main goroutine waiting for workers to complete...")
	wg.Wait()

	close(errChan)
	close(respChan)

	fmt.Println("All goroutines finished. Main goroutine exiting.")
}

func getRandomStatus(ctx context.Context, wg *sync.WaitGroup, respChan chan string, errChan chan error) {
	defer wg.Done()
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	for {
		req, err := http.NewRequest("GET", "http://httpbin.org/status/200,200,200,500", nil)
		if err != nil {
			errChan <- ErrSomethingWentWrong
			return
		}
		req = req.WithContext(ctx)
		resp, err := client.Do(req)
		if err != nil {
			errChan <- ErrSomethingWentWrong
			return
		}

		if resp.StatusCode != http.StatusOK {
			errChan <- errors.New("got a non-200 HTTP status")
			return
		}

		select {
		case <-ctx.Done():
			fmt.Println("Goroutine getRandomStatus received cancellation")
			return
		case respChan <- "Success from status":
		}
		time.Sleep(1 * time.Second)
	}
}

func getDelay(ctx context.Context, wg *sync.WaitGroup, respChan chan string, errChan chan error) {
	defer wg.Done()

	for {

		req, err := http.NewRequest("GET", "http://httpbin.org/delay/1", nil)
		if err != nil {
			errChan <- ErrSomethingWentWrong
			return
		}

		req = req.WithContext(ctx)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			errChan <- ErrSomethingWentWrong
			return
		}

		select {
		case <-ctx.Done():
			fmt.Println("Goroutine getDelay received cancellation")
			return
		case respChan <- "Success from delay: " + resp.Header.Get("date"):
		}
	}
}
