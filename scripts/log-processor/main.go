package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	logFiles, err := filepath.Glob("logs/*.log")
	if err != nil {
		log.Fatalf("Error finding log files: %v", err)
	}

	if len(logFiles) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "No log files found in logs/ directory")
		return
	}
}
