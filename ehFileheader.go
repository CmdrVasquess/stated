package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.FileheaderEvent.String()] = ehFileheader
}

func ehFileheader(ed *EDState, e events.Event) (chg att.Change, err error) {
	evt := e.(*journal.Fileheader)
	err = ed.WrLocked(func() error {
		ed.SetEDVersion(evt.GameVersion)
		ed.SetLanguage(evt.Language)
		return nil
	})
	return chg, err
}
