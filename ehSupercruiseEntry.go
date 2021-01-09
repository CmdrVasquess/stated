package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.SupercruiseEntryEvent.String()] = ehSupercruiseEntry
}

func ehSupercruiseEntry(ed *EDState, e events.Event) (chg att.Change) {
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
