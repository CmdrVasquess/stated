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
	"github.com/CmdrVasquess/stated/ships"
)

func Example_parseModuleItem() {
	const example = "int_powerplant_size3_class5"
	base, size, class := parseModuleItem(example)
	fmt.Println(base, size, class)
	// Output:
	// int_powerplant 3 5
}

func loadTestLoadout(t *testing.T, name string) *journal.Loadout {
	rd, err := os.Open(filepath.Join("testdata/loadouts", name))
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
	filter := map[string]bool{}
	if len(filter) > 0 {
		t.Error("incomlete test due to loadout filters")
	}

	const loadoutDir = "testdata/loadouts"
	dirls, err := os.ReadDir(loadoutDir)
	if err != nil {
		t.Fatal(err)
	}
	typeRepo := ships.FsTypeRepo{Dir: "ships/types"}
	for _, dire := range dirls {
		if filepath.Ext(dire.Name()) != ".json" {
			continue
		}
		name := dire.Name()
		name = name[:len(name)-5]
		generate := false
		if len(filter) > 0 {
			gen, ok := filter[name]
			if !ok {
				continue
			}
			generate = gen
		}
		t.Run(name, func(t *testing.T) {
			loadout := loadTestLoadout(t, dire.Name())
			ship, err := ShipFromLoadout(loadout, &typeRepo)
			if err != nil {
				t.Fatal(err)
			}
			var out bytes.Buffer
			enc := json.NewEncoder(&out)
			enc.SetIndent("", "  ")
			if err = enc.Encode(ship); err != nil {
				t.Fatal(err)
			}
			texstFile := filepath.Join("testdata", "texst", name+".texst")
			if generate {
				t.Errorf("generating texst file: %s", texstFile)
				texst.PrepareFile(texstFile, &out)
			} else {
				cmpr := texst.Compare{MismatchLimit: 1}
				if testing.Verbose() {
					cmpr.OnMismatch = texsting.MismatchError(t, name, false)
				}
				err = cmpr.RefFile(texstFile, &out)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
