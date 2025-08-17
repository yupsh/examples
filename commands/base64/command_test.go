package base64_test

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/yupsh/base64"
	"github.com/yupsh/base64/opt"
)

func ExampleBase64() {
	ctx := context.Background()
	input := strings.NewReader("Hello World!")

	cmd := base64.Base64()
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("base64: %v", err)
	}
	// Output:
	// SGVsbG8gV29ybGQh
}

func ExampleBase64_withDecode() {
	ctx := context.Background()
	input := strings.NewReader("SGVsbG8gV29ybGQh")

	cmd := base64.Base64(opt.Decode)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("base64: %v", err)
	}
	// Output:
	// Hello World!
}

func ExampleBase64_withWrap() {
	ctx := context.Background()
	input := strings.NewReader("This is a longer string that will demonstrate wrapping functionality")

	cmd := base64.Base64(opt.Wrap)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("base64: %v", err)
	}
	// Output:
	// VGhpcyBpcyBhIGxvbmdlciBzdHJpbmcgdGhhdCB3aWxsIGRlbW9uc3RyYXRlIHdyYXBwaW5nIGZ1
	// bmN0aW9uYWxpdHk=
}

func ExampleBase64_withWrapWidth() {
	ctx := context.Background()
	input := strings.NewReader("This is a test string for custom wrap width")

	cmd := base64.Base64(opt.WrapWidth(20))
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("base64: %v", err)
	}
	// Output:
	// VGhpcyBpcyBhIHRlc3Qg
	// c3RyaW5nIGZvciBjdXN0
	// b20gd3JhcCB3aWR0aA==
}

func ExampleBase64_withIgnoreGarbage() {
	ctx := context.Background()
	input := strings.NewReader("SGVsbG8g!!!V29ybGQh")

	cmd := base64.Base64(opt.Decode, opt.IgnoreGarbage)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("base64: %v", err)
	}
	// Output:
	// Hello World!
}

func ExampleBase64_withNoWrap() {
	ctx := context.Background()
	input := strings.NewReader("This is a very long string that would normally be wrapped but we want to keep it on one line for this example")

	cmd := base64.Base64(opt.NoWrap)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("base64: %v", err)
	}
	// Output:
	// VGhpcyBpcyBhIHZlcnkgbG9uZyBzdHJpbmcgdGhhdCB3b3VsZCBub3JtYWxseSBiZSB3cmFwcGVkIGJ1dCB3ZSB3YW50IHRvIGtlZXAgaXQgb24gb25lIGxpbmUgZm9yIHRoaXMgZXhhbXBsZQ==
}

func ExampleBase64_withMultipleFlags() {
	ctx := context.Background()
	input := strings.NewReader("SGVsbG8g!!!V29ybGQ=")

	cmd := base64.Base64(opt.Decode, opt.IgnoreGarbage)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("base64: %v", err)
	}
	// Output:
	// Hello World
}

func ExampleBase64_withFiles() {
	ctx := context.Background()
	input := strings.NewReader("")

	// This would read from actual files if they existed
	// For testing purposes, this demonstrates the API
	cmd := base64.Base64("testfile1.txt", "testfile2.txt")
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		// Expected to fail since files don't exist, but shows the pattern
		log.Panicf("base64: %v", err)
	}
}
