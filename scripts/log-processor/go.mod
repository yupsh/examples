module github.com/yupsh/examples

go 1.25.0

require (
	github.com/yupsh/cut v0.0.0
	github.com/yupsh/framework v0.0.0-20250901202619-a6797080e919
	github.com/yupsh/grep v0.0.0
)

replace (
	github.com/yupsh/cut => ../../../cut
	github.com/yupsh/framework => ../../../framework
	github.com/yupsh/grep => ../../../grep
)
