package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.LocationEvent.String()] = ehLocation
}

func ehLocation(ed *EDState, e events.Event) (chg att.Change) {
	ed.MustCommander(journal.LocationEvent.String())
	evt := e.(*journal.Location)
	sys := NewSystem(
		evt.SystemAddress,
		evt.StarSystem,
		evt.StarPos[:]...,
	)
	var loc Location
	switch {
	case evt.StationName != "":
		loc = &Port{
			Sys:    sys,
			Name:   evt.StationName,
			Type:   evt.StationType,
			Docked: evt.Docked,
		}
	default:
		loc = sys
	}
	must(ed.WrLocked(func() error {
		ed.Loc.Location = loc
		chg = ChgLocation
		return nil
	}))
	return chg
}
