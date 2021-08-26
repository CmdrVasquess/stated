package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.SupercruiseEntryEvent.String()] = jehSupercruiseEntry
}

func jehSupercruiseEntry(ed *EDState, e events.Event) (chg change.Flags) {
	ed.MustCommander(journal.SupercruiseEntryEvent.String())
	evt := e.(*journal.SupercruiseEntry)
	must(ed.WrLocked(func() error {
		sys, _ := ed.Loc.System().Update(
			evt.SystemAddress,
			evt.StarSystem,
		)
		ed.Loc.Location = sys
		chg |= ChgLocation
		return nil
	}))
	return chg
}
