package main

import (
	"flag"
	"log"
	"os"
)

var (
	targetDirFlag string
)

func init() {
	flag.StringVar(&targetDirFlag, "t", "", "specifies the target directory key (short-hand)")
	flag.StringVar(&targetDirFlag, "target", "", "specifies the target directory key")
}

func main() {
	flag.Parse()
	requireFlags(targetDirFlag)

	log.Printf("extracting GPS co-ordinates from %s....", targetDirFlag)

	extractor := NewExtractor()
	extractor.Extract(targetDirFlag)

	log.Printf("done")
	os.Exit(0)
}

func requireFlags(targetDirFlag string) {
	if targetDirFlag == "" {
		log.Fatal("targetDirFlag is required")
	}
}
