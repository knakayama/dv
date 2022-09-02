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

func TableFrom(m map[string]string) {
	t := newTable()
	t.AppendHeader(table.Row{"Region", "Default VPC"})

	for k, v := range m {
		t.AppendRows([]table.Row{
			{k, v},
		})
		t.AppendSeparator()
	}

	t.Render()
}
