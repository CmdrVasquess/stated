package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.DockedEvent.String()] = ehDocked
}

func ehDocked(ed *EDState, e events.Event) (chg att.Change, err error) {
	evt := e.(*journal.Docked)
	sys := ed.Galaxy.EdgxSystem(evt.SystemAddress, evt.StarSystem, nil, evt.Time)
	loc := &Port{
		Sys:    sys,
		Name:   evt.StationName,
		Type:   evt.StationType,
		Docked: true,
	}
	err = ed.WrLocked(func() error {
		ed.Loc = JSONLocation{loc}
		chg = ChgLocation
		return nil
	})
	return chg, err
}
