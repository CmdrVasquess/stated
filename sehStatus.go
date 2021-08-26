package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[events.StatusEvent.String()] = sehStatus
}

func sehStatus(ed *EDState, e events.Event) (chg change.Flags) {
	ed.MustCommander(events.StatusEvent.String())
	// evt := e.(*events.Status)
	return chg
}
