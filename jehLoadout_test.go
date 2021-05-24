package stated

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/fractalqb/texst"
	"github.com/fractalqb/texst/texsting"

	"github.com/CmdrVasquess/stated/journal"
	"github.com/CmdrVasquess/stated/ship"
)

func ExampleParseModuleItem() {
	const example = "int_powerplant_size3_class5"
	base, size, class := parseModuleItem(example)
	fmt.Println(base, size, class)
	// Output:
	// int_powerplant 3 5
}

func loadTestLoadout(t *testing.T, name string) *journal.Loadout {
	rd, err := os.Open(filepath.Join("test/loadouts", name))
	if err != nil {
		t.Fatal(err)
	}
	defer rd.Close()
	res := new(journal.Loadout)
	if err = json.NewDecoder(rd).Decode(res); err != nil {
		t.Fatal(err)
	}
	return res
}

func TestShipFromLoadout(t *testing.T) {
	tyr := ship.FsTypeRepo{Dir: "ship/types"}
	loadout := loadTestLoadout(t, "krait_mkii.1.json")
	ship, err := ShipFromLoadout(loadout, &tyr)
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	enc := json.NewEncoder(&out)
	enc.SetIndent("", "  ")
	if err = enc.Encode(ship); err != nil {
		t.Fatal(err)
	}
	cmpr := texst.Compare{MismatchLimit: 1}
	if testing.Verbose() {
		cmpr.OnMismatch = texsting.MismatchError(t, "krait_mkii.1", false)
	}
	err = cmpr.RefFile("test/texst/krait_mkii.1.texst", &out)
	if err != nil {
		t.Fatal(err)
	}
}
