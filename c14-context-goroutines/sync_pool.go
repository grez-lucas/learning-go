package main

import (
	"fmt"
	"sync"
	"time"
)

var logEntryPool = sync.Pool{
	New: func() interface{} {
		return &LogEntry{}
	},
}

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

// Log uses a shared pool to help the program reduce the amount of times it
// must allocate new memoery for a LogEntry struct. Instead it uses a pre-
// allocated one.
func Log(msg string) {
	entry := logEntryPool.Get().(*LogEntry)
	defer logEntryPool.Put(entry)

	entry.Message = msg
	entry.Timestamp = time.Now()
	entry.Level = "DEBUG"

	fmt.Printf("[%s] %s: %s\n", entry.Timestamp.Format(time.Kitchen), entry.Level, entry.Message)
}

func main() {
	fmt.Println("Starting pooled logging simulation...")
	// Simulate frequent logging from multiple places with goroutines
	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Log(fmt.Sprintf("Log message #%d", i))
		}(i)

	}
	wg.Wait()
	fmt.Println("Pooled logging simulation done.")
}
