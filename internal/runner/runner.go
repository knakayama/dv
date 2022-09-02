package runner

type Runner interface {
	Run(args []string) error
}
