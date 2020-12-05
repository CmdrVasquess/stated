package stated

import (
	"time"

	"github.com/CmdrVasquess/stated/att"
)

type Ship struct {
	Type     string
	Ident    att.String
	Name     att.String
	Cargo    att.Int
	MaxRange att.Float32
	MaxJump  att.Float32
	Berth    *Port      `json:",omitempty"`
	Sold     *time.Time `json:",omitempty"`
}

type Materials struct {
	Raw map[string]*Material
	Man map[string]*Material
	Enc map[string]*Material
}

type Material struct {
	Stock  int
	Demand int
}
