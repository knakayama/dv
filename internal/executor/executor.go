package executor

type Executor interface {
	Run(args []string) error
}
