package att

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

func TestInt_Set(t *testing.T) {
	var i Int
	if i.Set(1, 1) != 1 {
		t.Fatal("Change not detected")
	}
	if g := i.Get(); g != 1 {
		t.Fatalf("Got wrong value %d", g)
	}
	if i.Set(1, 2) != 0 {
		t.Fatal("Invalid change detected")
	}
	if g := i.Get(); g != 1 {
		t.Fatalf("Got wrong value %d", g)
	}
}

func TestInt_JSON(t *testing.T) {
	i := Int(4711)
	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	if err := enc.Encode(i); err != nil {
		t.Fatal(err)
	}
	jstr := sb.String()
	if jstr != "4711\n" {
		log.Fatalf("Wrote wrong JSON: '%s'", jstr)
	}
	dec := json.NewDecoder(strings.NewReader(sb.String()))
	if err := dec.Decode(&i); err != nil {
		t.Fatal(err)
	}
	if g := i.Get(); g != 4711 {
		t.Fatalf("Read wrong from JSON: %d", g)
	}
}
