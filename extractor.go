package main

import (
	"io/fs"
	"path/filepath"

	"github.com/senorprogrammer/gps-extractor/filetypes"
)

// Extractor extracts GPS data from image files
type Extractor struct {
}

// NewExtractor creates and returns an instance of GPSExtractor
func NewExtractor() *Extractor {
	return &Extractor{}
}

// Extract loops over a directory, reading any image files in it, looking for GPS data
func (ext *Extractor) Extract(targetDirPath string) ([]*filetypes.Image, error) {
	images := []*filetypes.Image{}

	err := filepath.Walk(targetDirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		imgFile := filetypes.NewImage(path)
		if imgFile.HasGPS() {
			images = append(images, imgFile)
		}

		return nil
	})

	if err != nil {
		return images, err
	}

	return images, nil
}
