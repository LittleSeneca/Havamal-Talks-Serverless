package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// main is the entry point of the application.
func main() {
	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())

	// Handle signals to cancel the context gracefully
	go handleSignals(cancel)

	severities := []string{"CRITICAL", "HIGH", "MEDIUM", "LOW"}
	for _, severity := range severities {
		fetchVulnerabilities(ctx, severity)
	}
}

// handleSignals listens for SIGINT and SIGTERM signals to cancel the context gracefully.
func handleSignals(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// When a signal is received, cancel the context to exit gracefully
	cancel()
}
