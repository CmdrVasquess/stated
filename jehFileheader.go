package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.FileheaderEvent.String()] = jehFileheader
}

func jehFileheader(ed *EDState, e events.Event) change.Flags {
	evt := e.(*journal.Fileheader)
	must(ed.WrLocked(func() error {
		ed.SetEDVersion(evt.GameVersion)
		ed.SetLanguage(evt.Language)
		return nil
	}))
	return ChgGame
}
