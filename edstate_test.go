package stated

import (
	"encoding/json"
	"fmt"

	"github.com/CmdrVasquess/watched"
)

var _ watched.EventRecv = (*EDState)(nil)

func ExampleSave() {
	eds := NewEDState(&Config{
		CmdrFile: CmdrFile{Dir: "."}.Filename,
	})
	eds.Cmdr = NewCommander("F4711")
	eds.Cmdr.Name = "John Doe"
	_, err := json.Marshal(eds)
	fmt.Println(err)
	// Output:
	// <nil>
}
