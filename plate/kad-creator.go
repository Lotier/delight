package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/lotier/kad"
)

func CreateLayout(inputFile string, itemName string) {
	// you can define settings and the layout in JSON
	content, readFileErr := ioutil.ReadFile(inputFile)
	if readFileErr != nil {
		log.Fatal("Error when opening file: ", readFileErr)
	}

	// create a new KAD instance
	cad := kad.New()

	// populate the 'cad' instance with the JSON contents
	err := json.Unmarshal(content, cad)
	if err != nil {
		log.Fatalf("Failed to parse json data into the KAD file\nError: %s", err.Error())
	}

	// and you can define settings via the KAD instance
	cad.Hash = itemName             // the name of the design
	cad.FileStore = kad.STORE_LOCAL // store the files locally
	cad.FileDirectory = "./output/" // the path location where the files will be saved
	cad.FileServePath = "/"         // the url path for the 'results' (don't worry about this)

	// lets draw the SVG files now
	err = cad.Draw()
	if err != nil {
		log.Fatal("Failed to Draw the KAD file\nError: %s", err.Error())
	}
}
