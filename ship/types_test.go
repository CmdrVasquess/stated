package ship

import (
	"encoding/json"
	"os"
	"testing"
)

func TestType_read(t *testing.T) {
	types := []string{
		"sidewinder.json",
		"eagle.json",
	}
	for _, typ := range types {
		t.Run(typ, func(t *testing.T) {
			rd, err := os.Open(typ)
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
