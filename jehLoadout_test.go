package stated

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CmdrVasquess/stated/ship"

	"github.com/CmdrVasquess/stated/journal"
)

func ExampleParseModuleItem() {
	const example = "int_powerplant_size3_class5"
	base, size, class := parseModuleItem(example)
	fmt.Println(base, size, class)
	// Output:
	// int_powerplant 3 5
}
