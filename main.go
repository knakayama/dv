package main

import (
	"os"

	"github.com/knakayama/remove-aws-default-vpcs/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
