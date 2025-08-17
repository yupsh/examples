# File Statistics Example

This example shows how to convert a shell script that analyzes files and directories into a yupsh Go program.

## What it does

Analyzes files in a directory and generates three types of statistics:
1. **File count by type** - Groups files by extension and counts them
2. **Largest files** - Shows the 10 largest files by size
3. **Total size** - Calculates total size of all files

## Shell Script Version

```bash
./analyze-files.sh [directory]
```

The shell script uses classic Unix pipeline patterns:
- `find | while read` loops
- `awk` for field processing and calculations
- `sort | uniq -c` for counting
- `sort -nr | head` for top-N results

## yupsh Go Version

```bash
go run main.go [directory]
```

The Go program replicates the exact same logic using yupsh commands while adding benefits like type safety and better error handling.

## Key yupsh Patterns Demonstrated

### 1. Custom Data Processing
```bash
# Shell: Extract file extensions in while loop
find . -name "*.*" | while read file; do
    echo "${file##*.}"
done
```

```go
// yupsh: Custom command for extension extraction
func ExtractExtension() yup.Command {
    return yup.CommandFunc(func(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
        return yup.ProcessLinesSimple(ctx, input, output,
            func(ctx context.Context, lineNum int, filename string, output io.Writer) error {
                ext := filepath.Ext(filename)
                // Process extension...
                return nil
            })
    })
}
```

### 2. File Information Processing
```bash
# Shell: ls -la with awk processing
find . -type f -exec ls -la {} \; | awk '{print $5 "\t" $9}'
```

```go
// yupsh: Native Go file operations
func FileWithSize() yup.Command {
    return yup.CommandFunc(func(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
        return yup.ProcessLinesSimple(ctx, input, output,
            func(ctx context.Context, lineNum int, filename string, output io.Writer) error {
                info, err := os.Stat(filename)
                if err != nil {
                    return nil
                }
                fmt.Fprintf(output, "%d\t%s\n", info.Size(), filename)
                return nil
            })
    })
}
```

### 3. Classic Unix Pipeline Patterns
```bash
# Shell
find . -name "*.*" | while read file; do echo "${file##*.}"; done | sort | uniq -c | sort -nr
```

```go
// yupsh: Direct pipeline translation
pipeline := yup.Pipe(
    find.Find(dir, find.Type("f"), find.Name("*.*")),
    ExtractExtension(),
    sort.Sort(),
    uniq.Uniq(uniq.Count),
    sort.Sort(sort.Numeric, sort.Reverse),
)
```

### 4. AWK-style Calculations
```bash
# Shell
awk '{sum += $5} END {print "Total: " sum " bytes"}'
```

```go
// yupsh: Using awk command directly
awk.Awk("{sum += $1} END {print \"Total: \" sum \" bytes\"}")
```

## Sample Run

Both versions produce identical output:

```
Analyzing files in: .
=== File Count by Type ===
      5 go
      3 md
      2 txt
      1 json

=== Largest Files ===
45231   ./main.go
12485   ./README.md
8934    ./data.json
3421    ./config.txt
1205    ./notes.txt

=== Total Size ===
Total: 71276 bytes
```

## Benefits of yupsh Version

- **Cross-platform**: Works on Windows, macOS, Linux
- **No external dependencies**: No need for `find`, `awk`, `sort`, `uniq` tools
- **Type safety**: File operations use Go's `os.Stat()` with proper error handling
- **Performance**: No subprocess overhead for each file operation
- **Maintainable**: Clear separation of concerns with reusable components
- **Testable**: Each component can be unit tested independently
- **Extensible**: Easy to add new analysis types or modify existing ones
