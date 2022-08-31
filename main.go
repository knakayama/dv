package main

import (
	"os"

	"github.com/knakayama/delete-aws-default-vpc/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
