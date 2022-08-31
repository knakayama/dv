package main

import (
	"os"

	"github.com/knakayama/dv/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
