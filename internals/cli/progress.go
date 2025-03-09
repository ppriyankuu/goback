package cli

import "fmt"

func TrackProgress(message string, args ...interface{}) {
	fmt.Printf(message, args...)
	fmt.Println()
}
