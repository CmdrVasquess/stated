package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
)

func init() {
	evtHdlrs[events.StatusEvent.String()] = sehStatus
}

func sehStatus(ed *EDState, e events.Event) (chg att.Change) {
	ed.MustCommander(events.StatusEvent.String())
	// evt := e.(*events.Status)
	return chg
}
