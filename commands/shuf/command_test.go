package shuf_test

import (
	"context"
	"log"
	"os"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/shuf"
	"github.com/yupsh/shuf/opt"
)

func ExampleShuf() {
	ctx := context.Background()
	input := strings.NewReader("line1\nline2\nline3\nline4\nline5\n")

	pipeline := yup.Exec(shuf.Shuf())
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show shuffled lines in random order
}

func ExampleShuf_withCount() {
	ctx := context.Background()
	input := strings.NewReader("line1\nline2\nline3\nline4\nline5\n")

	pipeline := yup.Exec(shuf.Shuf(opt.Count(3)))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show 3 randomly selected lines
}

func ExampleShuf_withInputRange() {
	ctx := context.Background()
	input := strings.NewReader("")

	pipeline := yup.Exec(shuf.Shuf(opt.InputRange("1-10")))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show numbers 1-10 in random order
}

func ExampleShuf_withEcho() {
	ctx := context.Background()
	input := strings.NewReader("")

	pipeline := yup.Exec(shuf.Shuf("apple", "banana", "cherry", "date", opt.Echo))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show the arguments in random order
}

func ExampleShuf_withRepeat() {
	ctx := context.Background()
	input := strings.NewReader("a\nb\nc\n")

	pipeline := yup.Exec(shuf.Shuf(opt.Repeat, opt.Count(10)))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show 10 lines with possible repetitions
}

func ExampleShuf_withZero() {
	ctx := context.Background()
	input := strings.NewReader("")

	// This would use NUL character as line delimiter - demonstrates pipeline pattern
	pipeline := yup.Exec(shuf.Shuf("file.txt", opt.Zero))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		// Expected to fail since files don't exist, but shows the pattern
		log.Panicf("shuf: %v", err)
	}
}

func ExampleShuf_withRandomSource() {
	ctx := context.Background()
	input := strings.NewReader("line1\nline2\nline3\n")

	// This would use a specific file as random source - demonstrates pipeline pattern
	pipeline := yup.Exec(shuf.Shuf(opt.RandomSource("/dev/urandom")))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show shuffled lines using specified random source
}

func ExampleShuf_rangeWithCount() {
	ctx := context.Background()
	input := strings.NewReader("")

	pipeline := yup.Exec(shuf.Shuf(opt.InputRange("1-100"), opt.Count(5)))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show 5 random numbers from 1-100
}

func ExampleShuf_multipleOptions() {
	ctx := context.Background()
	input := strings.NewReader("")

	pipeline := yup.Exec(shuf.Shuf("red", "green", "blue", "yellow", opt.Echo, opt.Count(2), opt.Repeat))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("shuf: %v", err)
	}
	// Output would show 2 colors with possible repetition
}
