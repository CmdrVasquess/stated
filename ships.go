package stated

import (
	"time"

	"github.com/fractalqb/change/chgv"

	"github.com/CmdrVasquess/stated/att"
)

type Ship struct {
	Type     string
	Ident    chgv.String
	Name     chgv.String
	Cargo    chgv.Int
	MaxRange att.Float32
	MaxJump  att.Float32
	Berth    *Port      `json:",omitempty"`
	Sold     *time.Time `json:",omitempty"`
}
