package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.UndockedEvent.String()] = ehUndocked
}

func ehUndocked(ed *EDState, e events.Event) (chg att.Change, err error) {
	evt := e.(*journal.Undocked)
	err = ed.WrLocked(func() error {
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
		chg = ChgLocation
		return nil
	})
	return chg, err
}
