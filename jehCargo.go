package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.CargoEvent.String()] = jehCargo
}

func jehCargo(ed *EDState, e events.Event) (chg change.Flags) {
	ed.MustCommander(journal.DockedEvent.String())
	evt := e.(*journal.Cargo)
	if evt.Vessel != "Ship" {
		return 0
	}
	must(ed.WrLocked(func() error {
		ed.Cargo = make(map[string]*Cargo)
		for _, item := range evt.Inventory {
			c := &Cargo{
				Name:   item.Name,
				Count:  item.Count,
				Stolen: item.Stolen,
			}
			ed.Cargo[c.Name] = c
		}
		return nil
	}))
	return chg
}
