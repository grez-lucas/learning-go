package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"time"
)

func HandleDelay(w http.ResponseWriter, r *http.Request) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randDelay := rand.Intn(10) + 1

	fmt.Println("Generated random delay: ", randDelay)

	select {
	case <-time.After(time.Duration(randDelay) * time.Second):
		fmt.Println("Handler: Random delay operation finished")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Operation Completed"))
	case <-r.Context().Done():
		fmt.Println("Handler: Received context done signal. Stopping work.")
		if r.Context().Err() == context.DeadlineExceeded {
			fmt.Println("Handler: Context deadline exceeded (timeout).")
			return
		} else {
			fmt.Println("Context cancelled for some other reason.")
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /delay", HandleDelay)

	timeoutHandler := TimeoutMiddleware(5 * time.Second)(mux)
	s := http.Server{
		Addr:         ":8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      timeoutHandler,
	}

	if err := s.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
		slog.Info("Server shutting down")
	}
}

type Middleware func(http.Handler) http.Handler

func TimeoutMiddleware(timeout time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

			if ctx.Err() == context.DeadlineExceeded {
				fmt.Printf("Request to %s timed out after %d seconds\n", r.Pattern, int(timeout.Seconds()))
			}
		})
	}
}
