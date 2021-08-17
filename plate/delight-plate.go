package main

import (
	"encoding/json"
	"log"

	"github.com/lotier/kad"
)

func BuildPlate() {
	// you can define settings and the layout in JSON
	json_bytes := []byte(`{"case":{"case-type":""},"switch-type":3,"stab-type":1,"layout":[[{"a":7},"esc"],["","1","2","3","4","5","6","7","8","9","0","-",": ",{"w":2},"backspace","home"],[{"w":1.5},"tab","q","w","e","r","t","y","u","i","o","p","[","]",{"w":1.5},"\\","pgup"],[{"w":1.75},"cap","a","s","d","f","g","h","j","k","l",";","'",{"w":2.25},"enter","pgdn"],[{"w":2.25},"shift","z","x","c","v","b","n","m",",",".","/",{"w":1.75},"shift","up","end"],[{"w":1.25},"ctrl",{"w":1.25},"win",{"w":1.25},"alt",{"w":6.25},"","alt","fn","ctrl","left","down","right"]],"top-padding":0,"left-padding":0,"right-padding":0,"bottom-padding":0,"custom":[{"layers":["switch"],"op":"cut","polygon":"custom-rounded-rectangle","radius":3,"width":300,"height":21,"rel_to":"[0,0]","points":"[-x+169,-y+9]"},{"layers":["switch"],"op":"cut","polygon":"custom-superellipse","radius":3,"rel_to":"[0,0]","points":"[-x+19,-y];[x,-y+19]"},{"layers":["switch"],"op":"cut","polygon":"custom-circle","diameter":4,"rel_to":"[0,0]","points":"[-x+40.9,-y+94.23];[-x+6.6,-y+57.4];[-x+53.85,-y+38.1];[-x+143.26,-y+42.16];[-x+152.37,-y+100.92];[-x+244.35,-y+38.26];[-x+262.38,-y+95.11]"}],"fillet":3}`)

	// create a new KAD instance
	cad := kad.New()

	// populate the 'cad' instance with the JSON contents
	err := json.Unmarshal(json_bytes, cad)
	if err != nil {
		log.Fatalf("Failed to parse json data into the KAD file\nError: %s", err.Error())
	}

	// and you can define settings via the KAD instance
	cad.Hash = "delight_plate"      // the name of the design
	cad.FileStore = kad.STORE_LOCAL // store the files locally
	cad.FileDirectory = "./output/" // the path location where the files will be saved
	cad.FileServePath = "/"         // the url path for the 'results' (don't worry about this)

	// lets draw the SVG files now
	err = cad.Draw()
	if err != nil {
		log.Fatal("Failed to Draw the KAD file\nError: %s", err.Error())
	}
}
