package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.SupercruiseExitEvent.String()] = jehSupercruiseExit
}

func jehSupercruiseExit(ed *EDState, e events.Event) (chg change.Flags) {
	ed.MustCommander(journal.SellShipOnRebuyEvent.String())
	evt := e.(*journal.SupercruiseExit)
	must(ed.WrLocked(func() error {
		sys, _ := ed.Loc.System().Update(
			evt.SystemAddress,
			evt.StarSystem,
		)
		if evt.BodyType != "Station" {
			ed.Loc.Location = sys
			return nil
		}
		port := &Port{
			Sys:  sys,
			Name: evt.Body,
		}
		ed.Loc.Location = port
		chg |= ChgLocation
		return nil
	}))
	return chg
}
