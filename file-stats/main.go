package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/awk"
	"github.com/yupsh/echo"
	"github.com/yupsh/find"
	"github.com/yupsh/head"
	"github.com/yupsh/sort"
	"github.com/yupsh/uniq"
)

// ExtractExtension extracts file extension from filename
func ExtractExtension() yup.Command {
	return yup.CommandFunc(func(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
		return yup.ProcessLinesSimple(ctx, input, output,
			func(ctx context.Context, lineNum int, filename string, output io.Writer) error {
				filename = strings.TrimSpace(filename)
				ext := filepath.Ext(filename)
				if ext != "" {
					// Remove the leading dot
					ext = ext[1:]
				} else {
					ext = "no-extension"
				}
				fmt.Fprintln(output, ext)
				return nil
			})
	})
}

// FileWithSize runs ls -la on each file and outputs size and name
func FileWithSize() yup.Command {
	return yup.CommandFunc(func(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
		return yup.ProcessLinesSimple(ctx, input, output,
			func(ctx context.Context, lineNum int, filename string, output io.Writer) error {
				filename = strings.TrimSpace(filename)
				if filename == "" {
					return nil
				}

				// Get file info
				info, err := os.Stat(filename)
				if err != nil {
					return nil // Skip files we can't stat
				}

				// Output size and filename (tab separated like awk would)
				fmt.Fprintf(output, "%d\t%s\n", info.Size(), filename)
				return nil
			})
	})
}

func main() {
	ctx := context.Background()

	// Get directory from command line argument, default to current directory
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	fmt.Printf("Analyzing files in: %s\n", dir)

	// === File Count by Type ===
	fmt.Println("=== File Count by Type ===")

	countByType := yup.Pipe(
		find.Find(dir, find.Type("f"), find.Name("*.*")), // Find files with extensions
		ExtractExtension(),                                 // Extract extensions
		sort.Sort(),                                        // Sort extensions
		uniq.Uniq(uniq.Count),                             // Count unique extensions
		sort.Sort(sort.Numeric, sort.Reverse),             // Sort by count (descending)
	)

	err := countByType.Execute(ctx, nil, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in file count analysis: %v\n", err)
	}

	// === Largest Files ===
	fmt.Println("\n=== Largest Files ===")

	largestFiles := yup.Pipe(
		find.Find(dir, find.Type("f")),    // Find all files
		FileWithSize(),                     // Get size and name for each file
		sort.Sort(sort.Numeric, sort.Reverse), // Sort by size (descending)
		head.Head(head.Lines(10)),         // Take top 10
	)

	err = largestFiles.Execute(ctx, nil, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in largest files analysis: %v\n", err)
	}

	// === Total Size ===
	fmt.Println("\n=== Total Size ===")

	totalSize := yup.Pipe(
		find.Find(dir, find.Type("f")),    // Find all files
		FileWithSize(),                     // Get size and name
		awk.Awk("{sum += $1} END {print \"Total: \" sum \" bytes\"}"), // Sum sizes
	)

	err = totalSize.Execute(ctx, nil, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in total size calculation: %v\n", err)
	}
}
