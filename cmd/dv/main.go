package main

import (
	"os"

	"github.com/knakayama/dv/internal/command"
)

func main() {
	command.Execute(os.Args[1:])
}
