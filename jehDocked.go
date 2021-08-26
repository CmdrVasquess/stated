package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.DockedEvent.String()] = jehDocked
}

func jehDocked(ed *EDState, e events.Event) (chg change.Flags) {
	ed.MustCommander(journal.DockedEvent.String())
	evt := e.(*journal.Docked)
	must(ed.WrLocked(func() error {
		sys := ed.Loc.System()
		sys, _ = sys.Update(evt.SystemAddress, evt.StarSystem)
		loc := &Port{
			Sys:    sys,
			Name:   evt.StationName,
			Type:   evt.StationType,
			Docked: true,
		}
		ed.Loc = JSONLocation{loc}
		chg |= ChgLocation
		return nil
	}))
	return chg
}
