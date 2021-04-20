package exporters

import (
	"path/filepath"

	"github.com/senorprogrammer/gps-extractor/filetypes"
)

// Base is the thing that manages exporting
type Base struct{}

// NewBase creates and returns an instance of Base
func NewBase() *Base {
	return &Base{}
}

// Export writes the extracted data to file by passing the responsibility off to
// file-specific exporters. If the file type is unknown, it raises an error
func (base *Base) Export(images []*filetypes.Image, outputFilePath string) error {
	extn := filepath.Ext(outputFilePath)

	switch extn {
	case ".csv":
		return ToCSV(images, outputFilePath)
	case ".htm", ".html":
		return ToHTML(images, outputFilePath)
	default:
		return ToTable(images)
	}
}
