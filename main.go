package main

import (
	"flag"
	"log"
	"os"

	"github.com/senorprogrammer/gps-extractor/exporters"
)

var (
	outputFileFlag string
	targetDirFlag  string
)

func init() {
	flag.StringVar(&outputFileFlag, "o", "", "specifies the output file, with extension: csv, html (short-hand)")
	flag.StringVar(&outputFileFlag, "output", "", "specifies the output file, with extension: csv, html")

	flag.StringVar(&targetDirFlag, "t", "", "specifies the target directory key (short-hand)")
	flag.StringVar(&targetDirFlag, "target", "", "specifies the target directory key")
}

func run() {
	extractor := NewExtractor()

	images, err := extractor.Extract(targetDirFlag)
	if err != nil {
		log.Fatal(err)
	}

	expBase := exporters.NewBase()
	err = expBase.Export(images, outputFileFlag)
	if err != nil {
		log.Fatal(err)
	}
}

/* -------------------- Main -------------------- */

func main() {
	flag.Parse()

	requiredFlags(targetDirFlag)

	run()

	os.Exit(0)
}

/* -------------------- Unexported Functions -------------------- */

// requiredFlags raise an error if the flag is not set
func requiredFlags(targetDirFlag string) {
	if targetDirFlag == "" {
		log.Fatal("'target' is required")
	}
}
