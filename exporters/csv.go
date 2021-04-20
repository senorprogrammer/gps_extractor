package exporters

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/senorprogrammer/gps-extractor/filetypes"
)

// ToCSV exports data in CSV format
func ToCSV(images []*filetypes.Image, outputFilePath string) error {
	log.Println("> exporting to csv")

	// Define the file headers
	data := [][]string{{"file", "latitude", "longitude"}}

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

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := csv.NewWriter(outFile)

	err = w.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}
