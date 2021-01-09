package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.ShutdownEvent.String()] = ehShutdown
}

func ehShutdown(ed *EDState, e events.Event) (chg att.Change) {
	must(ed.WrLocked(func() (err error) {
		if ed.ShutdownLogsOut {
			ed.SwitchCommander("", "")
		} else {
			err = ed.Save()
		}
		return err
	}))
	return 0 // TODO what if there was a commander
}
