# Log Processor Example

This example demonstrates how to convert a common shell script pattern into a yupsh Go program.

## What it does

Processes log files to extract error and warning entries, outputting timestamp and log level in CSV format.

## Shell Script Version

```bash
./process-logs.sh
```

The shell script:
1. Finds all `*.log` files in the `logs/` directory
2. For each file, greps for lines containing "error" or "warning" (case insensitive)
3. Extracts the timestamp (1st field) and level (2nd field) from each matching line
4. Writes results in CSV format to `results.csv`

## yupsh Go Version

```bash
go run main.go
```

The Go program does exactly the same thing using yupsh commands:
- Uses `find.Find()` to locate log files
- Uses `grep.Grep()` to filter log lines
- Uses `cut.Cut()` to extract specific fields
- Uses `tee.Tee()` to write output to file
- Uses the real `while.While()` command for line processing
- Creates reusable patterns like `ForEachFile()` for file iteration

## Key yupsh Patterns Demonstrated

### 1. Shell Loop → Go Function
```bash
# Shell
for file in logs/*.log; do
    # process file
done
```

```go
// yupsh
ForEachFile(func(filename string) yup.Command {
    return /* process file */
})
```

### 2. While Read Loop → While Command
```bash
# Shell
grep "pattern" file | while read line; do
    # process line
done
```

```go
// yupsh - using the real while.While command
import "github.com/yupsh/while"

yup.Pipe(
    grep.Grep("pattern"),
    while.While(func(line string) yup.Command {
        return /* process line */
    }),
)
```

### 3. Command Substitution → Field Extraction
```bash
# Shell
timestamp=$(echo "$line" | cut -d' ' -f1)
level=$(echo "$line" | cut -d' ' -f2)
```

```go
// yupsh
yup.Pipe(
    echo.Echo(line),
    cut.Cut(cut.Fields(1), cut.Delimiter(" ")),  // timestamp
)
yup.Pipe(
    echo.Echo(line),
    cut.Cut(cut.Fields(2), cut.Delimiter(" ")),  // level
)
```

## Sample Data

To test both versions, create some sample log files:

```bash
mkdir -p logs
echo "2024-01-01 ERROR Database connection failed" > logs/app.log
echo "2024-01-01 INFO Application started" >> logs/app.log
echo "2024-01-01 WARNING Low disk space" >> logs/app.log
echo "2024-01-01 ERROR Authentication failed" >> logs/app.log

echo "2024-01-01 INFO System healthy" > logs/system.log
echo "2024-01-01 ERROR Memory exhausted" >> logs/system.log
echo "2024-01-01 WARNING CPU usage high" >> logs/system.log
```

Both versions should produce identical `results.csv` output:

```csv
2024-01-01,ERROR
2024-01-01,WARNING
2024-01-01,ERROR
2024-01-01,ERROR
2024-01-01,WARNING
```

## Benefits of yupsh Version

- **Type Safety**: Compile-time checking vs runtime shell errors
- **Testing**: Easy to unit test individual components
- **Reusability**: `While()`, `ForEachFile()` patterns can be packaged
- **Performance**: No subprocess overhead
- **Portability**: Single binary works anywhere Go runs
- **Maintainability**: Clear structure and explicit error handling
