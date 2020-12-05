package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.CommanderEvent.String()] = ehCommander
}

func ehCommander(ed *EDState, e events.Event) (chg att.Change, err error) {
	evt := e.(*journal.Commander)
	err = ed.WrLocked(func() error {
		ed.SwitchCommander(evt.FID, evt.Name)
		return nil
	})
	return chg, err
}
