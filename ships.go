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
