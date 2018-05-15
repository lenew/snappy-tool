package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/snappy"
)

func main() {
	if len(os.Args) >= 2 {
		srcPath := os.Args[1]
		fr, err := os.Open(srcPath)
		if err != nil {
			log.Fatal(err)
		}
		defer fr.Close()
		sourceBuffer, err := ioutil.ReadAll(fr)
		outputLength, err := snappy.DecodedLen(sourceBuffer)
		if err != nil {
			log.Fatal(err)
		}
		outputBuffer := make([]byte, outputLength)

		snappy.Decode(outputBuffer, sourceBuffer)

		if len(os.Args) >= 3 {
			ioutil.WriteFile(os.Args[2], outputBuffer, 777)
		} else {
			os.Stdout.Write(outputBuffer)
		}
	}
}
