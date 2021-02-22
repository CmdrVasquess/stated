package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/CmdrVasquess/stated"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/CmdrVasquess/stated/ship"
)

var typesDir = "../../ship"

func loadShip(file string) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var je journal.Loadout
	json.Unmarshal(raw, &je)
	types := ship.FsTypeRepo{Dir: typesDir}
	ship, err := stated.ShipFromLoadout(&je, &types)
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	enc.Encode(ship)
}

func main() {
	for _, arg := range os.Args[1:] {
		loadShip(arg)
	}
}
