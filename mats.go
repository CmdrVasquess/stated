package stated

type Materials struct {
	Raw map[string]int `json:",omitempty"`
	Man map[string]int `json:",omitempty"`
	Enc map[string]int `json:",omitempty"`
}
