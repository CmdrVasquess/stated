package ships

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestType_read(t *testing.T) {
	des, err := os.ReadDir("types")
	if err != nil {
		t.Fatal(err)
	}
	var types []string
	for _, de := range des {
		if de.IsDir() {
			continue
		}
		if nm := de.Name(); filepath.Ext(nm) == ".json" {
			types = append(types, nm)
		}
	}

	for _, typ := range types {
		t.Run(typ, func(t *testing.T) {
			rd, err := os.Open(filepath.Join("types", typ))
			if err != nil {
				t.Fatal(err)
			}
			defer rd.Close()
			dec := json.NewDecoder(rd)
			dec.DisallowUnknownFields()
			var st Type
			err = dec.Decode(&st)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
