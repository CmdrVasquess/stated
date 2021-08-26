package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.CommanderEvent.String()] = jehCommander
}

func jehCommander(ed *EDState, e events.Event) change.Flags {
	evt := e.(*journal.Commander)
	must(ed.WrLocked(func() error {
		return ed.SwitchCommander(evt.FID, evt.Name)
	}))
	return ChgCommander // TODO changed more?
}
