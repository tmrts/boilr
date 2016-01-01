package tabular

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Print writes the given header and data in a tabular format to stdout.
func Print(header []string, data [][]string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, datum := range data {
		table.Append(datum)
	}

	if len(data) == 0 {
		table.Append([]string{"", "", ""})
	}

	table.Render()

	return nil
}
