package sort_test

import (
	"context"
	"log"
	"os"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/sort"
	"github.com/yupsh/sort/opt"
)

func ExampleSort() {
	ctx := context.Background()
	input := strings.NewReader("zebra\napple\nbanana\ncherry\n")

	pipeline := yup.Exec(sort.Sort())
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// apple
	// banana
	// cherry
	// zebra
}

func ExampleSort_withReverse() {
	ctx := context.Background()
	input := strings.NewReader("apple\nbanana\ncherry\n")

	pipeline := yup.Exec(sort.Sort(opt.Reverse))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// cherry
	// banana
	// apple
}

func ExampleSort_withNumeric() {
	ctx := context.Background()
	input := strings.NewReader("10\n2\n1\n20\n3\n")

	pipeline := yup.Exec(sort.Sort(opt.Numeric))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// 1
	// 2
	// 3
	// 10
	// 20
}

func ExampleSort_withUnique() {
	ctx := context.Background()
	input := strings.NewReader("apple\nbanana\napple\ncherry\nbanana\n")

	pipeline := yup.Exec(sort.Sort(opt.Unique))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// apple
	// banana
	// cherry
}

func ExampleSort_withIgnoreCase() {
	ctx := context.Background()
	input := strings.NewReader("Apple\nbanana\nCherry\napple\n")

	pipeline := yup.Exec(sort.Sort(opt.IgnoreCase))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// apple
	// Apple
	// banana
	// Cherry
}

func ExampleSort_withField() {
	ctx := context.Background()
	input := strings.NewReader("John,30,Engineer\nAlice,25,Designer\nBob,35,Manager\n")

	pipeline := yup.Exec(sort.Sort(opt.Delimiter(","), opt.Field(2)))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output sorted by age (second field)
}

func ExampleSort_withHumanNumeric() {
	ctx := context.Background()
	input := strings.NewReader("1K\n2M\n500\n1.5K\n")

	pipeline := yup.Exec(sort.Sort(opt.HumanNumeric))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// 500
	// 1K
	// 1.5K
	// 2M
}

func ExampleSort_withVersionSort() {
	ctx := context.Background()
	input := strings.NewReader("v1.10.0\nv1.2.0\nv1.9.0\nv2.0.0\n")

	pipeline := yup.Exec(sort.Sort(opt.VersionSort))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// v1.2.0
	// v1.9.0
	// v1.10.0
	// v2.0.0
}

func ExampleSort_withRandom() {
	ctx := context.Background()
	input := strings.NewReader("apple\nbanana\ncherry\ndate\n")

	pipeline := yup.Exec(sort.Sort(opt.Random))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output would show lines in random order
}

func ExampleSort_withIgnoreLeadingBlanks() {
	ctx := context.Background()
	input := strings.NewReader("  apple\n banana\n   cherry\n")

	pipeline := yup.Exec(sort.Sort(opt.IgnoreLeadingBlanks))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	//   apple
	//  banana
	//    cherry
}

func ExampleSort_multipleOptions() {
	ctx := context.Background()
	input := strings.NewReader("10\n2\n10\n3\n2\n")

	pipeline := yup.Exec(sort.Sort(opt.Numeric, opt.Reverse, opt.Unique))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("sort: %v", err)
	}
	// Output:
	// 10
	// 3
	// 2
}

func ExampleSort_withFiles() {
	ctx := context.Background()
	input := strings.NewReader("")

	// This would sort files if they existed - demonstrates pipeline pattern
	pipeline := yup.Exec(sort.Sort("file1.txt", "file2.txt", opt.Unique))
	if err := pipeline.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		// Expected to fail since files don't exist, but shows the pattern
		log.Panicf("sort: %v", err)
	}
}
