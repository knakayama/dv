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

func TableFrom(regionVpc map[string]string) {
	t := newTable()
	t.AppendHeader(table.Row{"Region", "Default VPC"})

	for region, vpc := range regionVpc {
		t.AppendRows([]table.Row{
			{region, vpc},
		})
		t.AppendSeparator()
	}

	t.SortBy([]table.SortBy{
		{
			Name: "Region",
		},
	})

	t.Render()
}
