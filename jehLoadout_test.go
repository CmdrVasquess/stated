package stated

import (
	"fmt"
)

func ExampleParseModuleItem() {
	const example = "int_powerplant_size3_class5"
	base, size, class := parseModuleItem(example)
	fmt.Println(base, size, class)
	// Output:
	// int_powerplant 3 5
}
