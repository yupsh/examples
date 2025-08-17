package cat_test

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/yupsh/cat"
	"github.com/yupsh/cat/opt"
	yup "github.com/yupsh/framework"
)

func ExampleCat() {
	ctx := context.Background()
	input := strings.NewReader("Hello World!\nThis is a test\n")

	pipeline := yup.Exec(cat.Cat())
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("cat: %v", err)
	}
	// Output:
	// Hello World!
	// This is a test
}

func ExampleCat_withNumberLines() {
	ctx := context.Background()
	input := strings.NewReader("Line one\nLine two\nLine three\n")

	pipeline := yup.Exec(cat.Cat(opt.NumberLines))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("cat: %v", err)
	}
	// Output:
	//      1	Line one
	//      2	Line two
	//      3	Line three
}

func ExampleCat_withShowEnds() {
	ctx := context.Background()
	input := strings.NewReader("First line\nSecond line\n")

	pipeline := yup.Exec(cat.Cat(opt.ShowEnds))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("cat: %v", err)
	}
	// Output:
	// First line$
	// Second line$
}

func ExampleCat_withShowTabs() {
	ctx := context.Background()
	input := strings.NewReader("Line with\ttabs\there\n")

	pipeline := yup.Exec(cat.Cat(opt.ShowTabs))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("cat: %v", err)
	}
	// Output:
	// Line with	tabs	here
}

func ExampleCat_withSqueezeBlank() {
	ctx := context.Background()
	input := strings.NewReader("Line 1\n\n\nLine 2\n\n\n\nLine 3\n")

	pipeline := yup.Exec(cat.Cat(opt.SqueezeBlank))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("cat: %v", err)
	}
	// Output:
	// Line 1
	//
	// Line 2
	//
	// Line 3
}

func ExampleCat_withMultipleFlags() {
	ctx := context.Background()
	input := strings.NewReader("First line\n\n\nSecond line\n")

	pipeline := yup.Exec(cat.Cat(opt.NumberLines, opt.ShowEnds, opt.SqueezeBlank))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("cat: %v", err)
	}
	// Output:
	//      1	First line$
	//      2	$
	//      3	Second line$
}

func ExampleCat_withFiles() {
	ctx := context.Background()
	input := strings.NewReader("")

	// This would read from actual files if they existed - demonstrates pipeline pattern
	pipeline := yup.Exec(cat.Cat("testfile1.txt", "testfile2.txt"))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		// Expected to fail since files don't exist, but shows the pattern
		log.Panicf("cat: %v", err)
	}
}
