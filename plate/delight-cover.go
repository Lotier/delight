package main

import (
	"encoding/json"
	"log"

	"github.com/lotier/kad"
)

func BuildCover() {
	// you can define settings and the layout in JSON
	json_bytes := []byte(`{"top-padding":0,"left-padding":0,"right-padding":0,"bottom-padding":0,"custom":[{"layers":["switch"],"op":"add","polygon":"custom-rounded-rectangle","radius":3,"width":284.13,"height":17.684,"rel_to":"[0,0]","points":"[0,0]"},{"layers":["switch"],"op":"cut","polygon":"custom-circle","diameter":2.2,"rel_to":"[0,0]","points":"[x-137,y-2.715];[x-137,y+8.715];[x-19.398,y-2.715];[x-19.398,y+8.715];[x+141.13,y-2.715];[x+141.13,y+8.715];"}]}`)

	// create a new KAD instance
	cad := kad.New()

	// populate the 'cad' instance with the JSON contents
	err := json.Unmarshal(json_bytes, cad)
	if err != nil {
		log.Fatalf("Failed to parse json data into the KAD file\nError: %s", err.Error())
	}

	// and you can define settings via the KAD instance
	cad.Hash = "delight_cover"      // the name of the design
	cad.FileStore = kad.STORE_LOCAL // store the files locally
	cad.FileDirectory = "./output/" // the path location where the files will be saved
	cad.FileServePath = "/"         // the url path for the 'results' (don't worry about this)

	// lets draw the SVG files now
	err = cad.Draw()
	if err != nil {
		log.Fatal("Failed to Draw the KAD file\nError: %s", err.Error())
	}
}
