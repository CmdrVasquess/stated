package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.ShutdownEvent.String()] = jehShutdown
}

func jehShutdown(ed *EDState, e events.Event) (chg att.Change) {
	must(ed.WrLocked(func() (err error) {
		if ed.ShutdownLogsOut {
			ed.SwitchCommander("", "")
		} else {
			err = ed.Save()
		}
		return err
	}))
	return ChgGame
}
