package ship

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestType_read(t *testing.T) {
	raw, err := ioutil.ReadFile("eagle.json")
	if err != nil {
		t.Fatal(err)
	}
	var st Type
	err = json.Unmarshal(raw, &st)
	if err != nil {
		t.Fatal(err)
	}
}
