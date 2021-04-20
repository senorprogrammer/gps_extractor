package main

type ExtractedData struct {
	Latitude  string
	Longitude string
}

func NewExtractedData() *ExtractedData {
	return &ExtractedData{}
}
