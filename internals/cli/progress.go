package cli

import "fmt"

// TrackProgress logs the progress of an operation with a formatted message.

// Parameters:
// - message: The format string for the message to be logged.
// - args: Additional arguments to format the message.
func TrackProgress(message string, args ...any) {
	fmt.Printf(message, args...)
	fmt.Println()
}
