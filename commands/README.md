# yupsh Examples

This directory contains real, executable examples that demonstrate how to convert shell scripts into yupsh Go programs.

Each example includes:
- **Shell script** (`.sh`) - Traditional Unix shell implementation
- **Go program** (`main.go`) - Equivalent yupsh implementation
- **Go module** (`go.mod`) - Dependencies for the Go version
- **README** - Detailed comparison and explanation

## Available Examples

### 📊 [log-processor](./log-processor/)
**Shell Pattern**: File processing with grep, cut, and while-read loops
**Demonstrates**: Line-by-line processing, field extraction, CSV output

```bash
# Shell version
./log-processor/process-logs.sh

# yupsh version
cd log-processor && go run main.go
```

### 📈 [file-stats](./file-stats/)
**Shell Pattern**: File analysis with find, awk, sort, uniq
**Demonstrates**: Custom data processing, numeric sorting, aggregation

```bash
# Shell version
./file-stats/analyze-files.sh /path/to/analyze

# yupsh version
cd file-stats && go run main.go /path/to/analyze
```

## Key yupsh Patterns Shown

### 🔄 **Loop Patterns**
| Shell Pattern | yupsh Equivalent |
|---------------|------------------|
| `for file in *.txt` | `find.Find(".", find.Name("*.txt"))` |
| `while read line` | `While(func(line string) yup.Command {...})` |
| `find ... -exec` | `find.Find(...) \| ForEachFile(...)` |

### 🔧 **Processing Patterns**
| Shell Pattern | yupsh Equivalent |
|---------------|------------------|
| `cut -d' ' -f1` | `cut.Cut(cut.Fields(1), cut.Delimiter(" "))` |
| `grep "pattern"` | `grep.Grep("pattern")` |
| `sort \| uniq -c` | `sort.Sort() \| uniq.Uniq(uniq.Count)` |
| `awk '{print $1}'` | `awk.Awk("{print $1}")` |

### 🎯 **Control Flow Patterns**
| Shell Pattern | yupsh Equivalent |
|---------------|------------------|
| `[[ -f "$file" ]]` | `IfFileExists(file, ...)` |
| `>> output.txt` | `tee.Tee("output.txt")` |
| `command1 \| command2` | `yup.Pipe(command1, command2)` |

## Running Examples

Each example can be run independently:

```bash
# Run shell version
chmod +x */*.sh
./log-processor/process-logs.sh

# Run Go version
cd log-processor
go run main.go
```

## Creating New Examples

To add a new example:

1. **Create directory**: `mkdir examples/my-example`
2. **Add shell script**: Write `my-script.sh` with traditional approach
3. **Add Go program**: Write `main.go` with yupsh equivalent
4. **Add go.mod**: Include necessary yupsh dependencies
5. **Add README**: Document the comparison and patterns

### Example Template

```
examples/my-example/
├── my-script.sh      # Shell implementation
├── main.go           # yupsh implementation
├── go.mod            # Go dependencies
└── README.md         # Detailed comparison
```

## Philosophy

These examples demonstrate that yupsh programs:
- **Think like shell scripts** - Same logical flow and patterns
- **Work like Go programs** - Type safety, error handling, testing
- **Perform better** - No subprocess overhead, compiled performance
- **Deploy easier** - Single binary, no external dependencies

The goal is to show that shell scripting concepts translate naturally to yupsh, while gaining all the benefits of Go's ecosystem and tooling.

## Contributing Examples

Have a great shell → yupsh conversion? Please contribute!

1. Follow the template structure above
2. Ensure both versions produce identical output
3. Include comprehensive README with pattern explanations
4. Add any test data or setup needed to run the example

**Great example candidates:**
- Log processing and analysis
- File manipulation and organization
- Data extraction and transformation
- System monitoring scripts
- Build and deployment automation
- Text processing workflows
