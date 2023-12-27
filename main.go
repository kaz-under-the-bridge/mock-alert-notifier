package main

import (
	"fmt"
	"os"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		// exit with error with printing error message
		fmt.Printf("Program aborted due to fatal error while executing command: %s", err)
		os.Exit(1)
	}
}
