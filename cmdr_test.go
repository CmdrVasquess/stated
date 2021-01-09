package stated

import (
	"encoding/json"
	"os"
	"testing"
)

func ExampleRanks() {
	var rs Ranks
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(rs)
	// Output:
	// {"CQC":{"L":0,"P":0},"Combat":{"L":0,"P":0},"Empire":{"L":0,"P":0},"Explore":{"L":0,"P":0},"Federation":{"L":0,"P":0},"Trade":{"L":0,"P":0}}
}

func TestRanksUnmarshal(t *testing.T) {
	var rs Ranks
	json.Unmarshal([]byte(`{
		"Explore":{"L":4,"P":0}
	}`), &rs)
	if l := rs[Explore].Level; l != 4 {
		t.Errorf("unexpected explorer level %d", l)
	}
}
