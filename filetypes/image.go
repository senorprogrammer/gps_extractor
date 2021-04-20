package filetypes

import (
	"image"
	"io/ioutil"
	"os"

	"github.com/h2non/filetype"
	"github.com/rwcarlsen/goexif/exif"
)

// Image is a wrapper struct that defines an image on disk
type Image struct {
	Hash  string
	Image image.Image
	Name  string
	Path  string
	Size  int64
}

// NewImage creates and returns an instance of ImageFile
func NewImage(path string) *Image {
	imgFile := Image{
		Hash: "",
		Name: "",
		Path: path,
		Size: 0,
	}

	imgFile.loadFileInfo()

	return &imgFile
}

/* -------------------- Global Functions -------------------- */

// IsImage returns TRUE if this represents a known image type,
// FALSE if it does not
func IsImage(path string) bool {
	buf, _ := ioutil.ReadFile(path)
	return filetype.IsImage(buf)
}

/* -------------------- Unexported Functions -------------------- */

// HasGPS returns TRUE if this image file contains GPS EXIF data,
// FALSE if it does not
func (imgFile *Image) HasGPS() bool {
	lat, lon := imgFile.LatLon()
	if lat != 0 && lon != 0 {
		return true
	}
	return false
}

// LatLon returns the Latitude and Longitude read from the EXIF data
func (imgFile *Image) LatLon() (lat, lon float64) {
	file, err := os.Open(imgFile.Path)
	if err != nil {
		return
	}
	defer file.Close()

	exifData, err := exif.Decode(file)
	if err != nil {
		return
	}

	lat, lon, _ = exifData.LatLong()

	return lat, lon
}

/* -------------------- Unexported Functions -------------------- */

func (imgFile *Image) loadFileInfo() {
	info, err := os.Stat(imgFile.Path)
	if err == nil {
		imgFile.Name = info.Name()
		imgFile.Size = info.Size()
	}
}
