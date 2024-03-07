package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Command line parsing
	durationFlag := flag.String("d", "", "Duration for countdown timer (e.g., 2h45m). Leave empty for stopwatch mode.")
	flag.Parse()

	// Determine mode and calculate end time if necessary
	var endTime time.Time
	countdownMode := false
	if *durationFlag != "" {
		countdownMode = true
		duration, err := time.ParseDuration(*durationFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid duration format: %v\n", err)
			os.Exit(1)
		}
		endTime = time.Now().Add(duration)
	}

	startTime := time.Now()
	fmt.Printf("Timer started at: %s\n", startTime.Format("2006-01-02 15:04:05"))

	// Setting up channel and signal notifier for capturing CTRL-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Using a ticker for the timer functionality, firing every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if countdownMode {
					remaining := endTime.Sub(time.Now()).Round(time.Second)
					if remaining <= 0 {
						fmt.Printf("\nCountdown finished at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
						os.Exit(0)
					}
					fmt.Printf("\rRemaining time: %s", remaining)
				} else {
					fmt.Printf("\rTime elapsed: %s", time.Since(startTime).Round(time.Second))
				}
			case <-sigChan:
				if countdownMode {
					fmt.Printf("\nCountdown stopped at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
				} else {
					fmt.Printf("\nTimer stopped at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
					fmt.Printf("Total time elapsed: %s\n", time.Since(startTime).Round(time.Second))
				}
				os.Exit(0)
			}
		}
	}()

	// Keeping the main goroutine alive
	select {}
}
