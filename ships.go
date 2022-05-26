package stated

import (
	"time"

	"github.com/CmdrVasquess/stated/att"
	"github.com/fractalqb/change"
)

type Ship struct {
	Type     string
	Ident    change.Val[string]
	Name     change.Val[string]
	Cargo    change.Val[int]
	MaxRange att.Float32
	MaxJump  att.Float32
	Berth    *Port      `json:",omitempty"`
	Sold     *time.Time `json:",omitempty"`
}
