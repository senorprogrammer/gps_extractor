package main

import (
	"flag"
	"log"
	"os"

	"github.com/senorprogrammer/gps-extractor/exporters"
)

const (
	csvFilePath = "images.csv"
)

var (
	targetDirFlag string
)

func init() {
	flag.StringVar(&targetDirFlag, "t", "", "specifies the target directory key (short-hand)")
	flag.StringVar(&targetDirFlag, "target", "", "specifies the target directory key")
}

func run() {
	extractor := NewExtractor()

	images, err := extractor.Extract(targetDirFlag)
	if err != nil {
		log.Fatal(err)
	}

	err = exporters.ToCSV(images, csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
}

/* -------------------- Main -------------------- */

func main() {
	flag.Parse()
	requireFlags(targetDirFlag)

	run()

	log.Printf("done")
	os.Exit(0)
}

func requireFlags(targetDirFlag string) {
	if targetDirFlag == "" {
		log.Fatal("'target' is required")
	}
}
