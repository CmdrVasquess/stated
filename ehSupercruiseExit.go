package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.SupercruiseExitEvent.String()] = ehSupercruiseExit
}

func ehSupercruiseExit(ed *EDState, e events.Event) (chg att.Change, err error) {
	evt := e.(*journal.SupercruiseExit)
	err = ed.WrLocked(func() error {
		if ed.Loc.Location == nil {
			return nil
		}
		if evt.BodyType != "Station" {
			return nil
		}
		sys := ed.Loc.System()
		port := &Port{
			Sys:  sys,
			Name: evt.Body,
		}
		ed.Loc.Location = port
		chg = ChgLocation
		return nil
	})
	return chg, err
}
