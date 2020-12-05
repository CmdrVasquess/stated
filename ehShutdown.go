package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.ShutdownEvent.String()] = ehShutdown
}

func ehShutdown(ed *EDState, e events.Event) (chg att.Change, err error) {
	err = ed.WrLocked(func() error {
		if ed.ShutdownLogsOut {
			ed.SwitchCommander("", "")
		} else if ed.Cmdr != nil && ed.CmdrFile != nil {
			cmdrFile := ed.CmdrFile(ed.Cmdr.FID, ed.Cmdr.Name.Get())
			if err := ed.Cmdr.Save(cmdrFile); err != nil {
				return err
			}
		}
		return nil
	})
	return 0, err
}
