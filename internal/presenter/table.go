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

func TableFrom(kv map[string]string, headers [2]string) {
	t := newTable()
	t.AppendHeader(table.Row{headers[0], headers[1]})

	for k, v := range kv {
		t.AppendRows([]table.Row{
			{k, v},
		})
		t.AppendSeparator()
	}

	t.SortBy([]table.SortBy{
		{
			Name: headers[0],
		},
	})

	t.Render()
}
