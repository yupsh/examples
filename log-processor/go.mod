module log-processor

go 1.23.0

require (
	github.com/yupsh/cat v0.0.0
	github.com/yupsh/cut v0.0.0
	github.com/yupsh/echo v0.0.0
	github.com/yupsh/find v0.0.0
	github.com/yupsh/framework v0.1.0
	github.com/yupsh/grep v0.0.0
	github.com/yupsh/tee v0.0.0
	github.com/yupsh/while v0.0.0
)

replace github.com/yupsh/framework => ../../framework

replace github.com/yupsh/find => ../../find

replace github.com/yupsh/cat => ../../cat

replace github.com/yupsh/grep => ../../grep

replace github.com/yupsh/cut => ../../cut

replace github.com/yupsh/echo => ../../echo

replace github.com/yupsh/tee => ../../tee

replace github.com/yupsh/while => ../../while
