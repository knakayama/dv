package presenter

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func newTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	return t
}

func PrintTableIn(keyValue map[string]string) {
	t := newTable()
	t.AppendHeader(table.Row{"Region", "Default VPC"})

	for k, v := range keyValue {
		t.AppendRows([]table.Row{
			{k, v},
		})
		t.AppendSeparator()
	}

	t.Render()
}
