package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/h2non/filetype"
)

// Extractor extracts GPS data from image files
type Extractor struct {
}

// NewExtractor creates and returns an instance of GPSExtractor
func NewExtractor() *Extractor {
	return &Extractor{}
}

// Extract loops over a directory, reading any image files in it, looking for GPS data
func (ext *Extractor) Extract(targetDirPath string) ([]*ExtractedData, error) {
	data := []*ExtractedData{}

	err := filepath.Walk(targetDirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !ext.isImage(path) {
			return nil
		}

		log.Println(path)
		gpsData, _ := ext.extractGPS(info)
		if data != nil {
			data = append(data, gpsData)
		}
		return nil
	})
	if err != nil {
		return data, err
	}

	return data, nil
}

/* -------------------- Unexported Functions -------------------- */

// extractGPS checks this image for GPS data
func (ext *Extractor) extractGPS(info fs.FileInfo) (*ExtractedData, error) {
	data := NewExtractedData()
	return data, nil
}

//isImage returns TRUE if this file is an image file, FALSE if it is not
func (ext *Extractor) isImage(filePath string) bool {
	buf, _ := ioutil.ReadFile(filePath)
	return filetype.IsImage(buf)
}
