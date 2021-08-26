package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.UndockedEvent.String()] = jehUndocked
}

func jehUndocked(ed *EDState, e events.Event) (chg change.Flags) {
	ed.MustCommander(journal.UndockedEvent.String())
	evt := e.(*journal.Undocked)
	must(ed.WrLocked(func() error {
		if port := ed.Loc.Port(); port == nil {
			port := &Port{
				Name:   evt.StationName,
				Type:   evt.StationType,
				Docked: false,
			}
			ed.Loc.Location = port
		} else {
			port.Docked = false
			if port.Name != evt.StationName {
				port.Name = evt.StationName
				port.Type = evt.StationType
				port.Sys = nil
			}
		}
		chg |= ChgLocation
		return nil
	}))
	return chg
}
