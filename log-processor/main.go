package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/while"
)

// Helper for creating simple commands
type simpleCommand func(ctx context.Context, input io.Reader, output, stderr io.Writer) error

func (f simpleCommand) Execute(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
	return f(ctx, input, output, stderr)
}

// ProcessLogLine extracts timestamp and level from a log line
func ProcessLogLine(line string) yup.Command {
	return simpleCommand(func(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
		// Simple field extraction by splitting on space
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			timestamp := fields[0]
			level := fields[1]
			fmt.Fprintf(output, "%s,%s\n", timestamp, level)
		}
		return nil
	})
}

func main() {
	ctx := context.Background()

	// Create logs directory and sample files for demo
	os.MkdirAll("logs", 0755)

	// Create sample log files
	sampleLogs := []struct {
		filename string
		content  string
	}{
		{
			"logs/app.log",
			"2024-01-01 ERROR Database connection failed\n2024-01-01 INFO Application started\n2024-01-01 WARNING Low disk space\n",
		},
		{
			"logs/system.log",
			"2024-01-01 INFO System healthy\n2024-01-01 ERROR Memory exhausted\n2024-01-01 WARNING CPU usage high\n",
		},
	}

	for _, log := range sampleLogs {
		os.WriteFile(log.filename, []byte(log.content), 0644)
	}

	// Read each log file and process error/warning lines
	var allLines []string

	// Process each log file
	for _, log := range sampleLogs {
		fmt.Printf("Processing %s\n", log.filename)

		content, err := os.ReadFile(log.filename)
		if err != nil {
			continue
		}

		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			// Check if line contains error or warning (case insensitive)
			lowerLine := strings.ToLower(line)
			if strings.Contains(lowerLine, "error") || strings.Contains(lowerLine, "warning") {
				allLines = append(allLines, line)
			}
		}
	}

	// Process all matching lines using the while command
	input := strings.Join(allLines, "\n")

	var csvOutput strings.Builder
	cmd := while.While(func(line string) yup.Command {
		return ProcessLogLine(line)
	})

	err := cmd.Execute(ctx, strings.NewReader(input), &csvOutput, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing lines: %v\n", err)
		os.Exit(1)
	}

	// Write results to file
	os.WriteFile("results.csv", []byte(csvOutput.String()), 0644)

	// Show results
	fmt.Print(csvOutput.String())
	fmt.Println("Results written to results.csv")
}
