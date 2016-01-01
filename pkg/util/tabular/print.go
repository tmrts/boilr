package tabular

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// Print writes the given header and data in a tabular format to stdout.
func Print(header []string, data [][]string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetCenterSeparator(color.GreenString("+"))
	table.SetColumnSeparator(color.GreenString("|"))
	table.SetRowSeparator(color.GreenString("-"))

	table.SetHeader(header)

	for _, datum := range data {
		datum[0] = color.RedString(datum[0])

		if strings.HasPrefix(datum[1], "local:") {
			datum[1] = color.BlueString(datum[1])
		} else {
			datum[1] = color.YellowString(datum[1])
		}

		table.Append(datum)
	}

	if len(data) == 0 {
		table.Append([]string{"", "", ""})
	}

	table.Render()

	return nil
}
