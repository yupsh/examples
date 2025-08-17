package exec_test

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/yupsh/exec"
	"github.com/yupsh/exec/opt"
)

func ExampleExec() {
	ctx := context.Background()
	input := strings.NewReader("")

	cmd := exec.Exec("echo", "Hello from exec!")
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("exec: %v", err)
	}
	// Output: Hello from exec!
}

func ExampleExec_withShell() {
	ctx := context.Background()
	input := strings.NewReader("")

	cmd := exec.Exec("echo 'Hello from shell' | tr a-z A-Z", opt.UseShell)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("exec: %v", err)
	}
	// Output: HELLO FROM SHELL
}

func ExampleExec_withWorkingDir() {
	ctx := context.Background()
	input := strings.NewReader("")

	cmd := exec.Exec("pwd", opt.WorkingDir("/tmp"))
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("exec: %v", err)
	}
	// Output: /private/tmp
}

func ExampleExec_quiet() {
	ctx := context.Background()
	input := strings.NewReader("")

	// This command would normally output to stderr, but with Quiet flag it's suppressed
	cmd := exec.Exec("echo", "This goes to stdout", opt.Quiet)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("exec: %v", err)
	}
	// Output: This goes to stdout
}

func ExampleExec_withEnvironmentVariable() {
	ctx := context.Background()
	input := strings.NewReader("")

	cmd := exec.Exec("echo", "${TEST_VAR}", opt.EnvVar("TEST_VAR=Hello World"), opt.UseShell)
	if err := cmd.Execute(ctx, input, os.Stdout, os.Stderr); err != nil {
		log.Panicf("exec: %v", err)
	}
	// Output: Hello World
}
