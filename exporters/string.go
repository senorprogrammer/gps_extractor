package exporters

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/senorprogrammer/gps-extractor/filetypes"
)

// ToString writes the file data out to the console as a string
func ToString(images []*filetypes.Image) error {
	// Define the file headers
	data := [][]string{}

	for _, image := range images {
		lat, lon := image.LatLon()
		data = append(
			data, []string{
				image.Path,
				strconv.FormatFloat(lat, 'f', -1, 64),
				strconv.FormatFloat(lon, 'f', -1, 64),
			},
		)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"file", "latitude", "longitude"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}
