package wc_test

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/yupsh/wc"
	"github.com/yupsh/wc/opt"
)

func ExampleWc() {
	ctx := context.Background()
	input := strings.NewReader("Hello World!\nThis is a test\nThird line\n")

	cmd := wc.Wc()
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("wc: %v", err)
	}
	// Output:
	//       3       8      39
}

func ExampleWc_withLines() {
	ctx := context.Background()
	input := strings.NewReader("Line one\nLine two\nLine three\n")

	cmd := wc.Wc(opt.Lines)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("wc: %v", err)
	}
	// Output:
	//        3
}

func ExampleWc_withWords() {
	ctx := context.Background()
	input := strings.NewReader("Hello world this is a test\n")

	cmd := wc.Wc(opt.Words)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("wc: %v", err)
	}
	// Output:
	//        6
}

func ExampleWc_withChars() {
	ctx := context.Background()
	input := strings.NewReader("Hello\n")

	cmd := wc.Wc(opt.Chars)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("wc: %v", err)
	}
	// Output:
	//        6
}

func ExampleWc_withBytes() {
	ctx := context.Background()
	input := strings.NewReader("Hello\n")

	cmd := wc.Wc(opt.Bytes)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("wc: %v", err)
	}
	// Output:
	//        6
}

func ExampleWc_withMaxLength() {
	ctx := context.Background()
	input := strings.NewReader("Short\nThis is a longer line\nMedium line\n")

	cmd := wc.Wc(opt.MaxLength)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("wc: %v", err)
	}
	// Output:
	//       21
}

func ExampleWc_withMultipleFlags() {
	ctx := context.Background()
	input := strings.NewReader("Hello world\nSecond line\n")

	cmd := wc.Wc(opt.Lines, opt.Words)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("wc: %v", err)
	}
	// Output:
	//       2       4
}
