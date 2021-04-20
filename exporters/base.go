package exporters

import (
	"fmt"
	"path/filepath"

	"github.com/senorprogrammer/gps-extractor/filetypes"
)

// Base is a terrible name for the thing that manages exporting
type Base struct{}

// NewBase creates and returns an instance of Base
func NewBase() *Base {
	return &Base{}
}

// Export writes the extracted data to file
func (base *Base) Export(images []*filetypes.Image, outputFilePath string) error {
	extn := filepath.Ext(outputFilePath)

	switch extn {
	case ".csv":
		return ToCSV(images, outputFilePath)
	case ".htm", ".html":
		return ToHTML(images, outputFilePath)
	default:
		return fmt.Errorf("export file type undefined or unknown: %s", extn)
	}
}
