package stated

import (
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
)

func init() {
	evtHdlrs[journal.FSDJumpEvent.String()] = ehFSDJump
}

func ehFSDJump(ed *EDState, e events.Event) (chg att.Change) {
	cmdr := ed.MustCommander(journal.FSDJumpEvent.String())
	evt := e.(*journal.FSDJump)
	chg = ChgSystem
	sys := NewSystem(
		evt.SystemAddress,
		evt.StarSystem,
		evt.StarPos[:]...,
	)
	must(ed.WrLocked(func() error {
		ed.Loc.Location = sys
		if cmdr.inShip != nil && evt.JumpDist > float32(cmdr.inShip.MaxJump) {
			chg |= cmdr.inShip.MaxJump.Set(evt.JumpDist, ChgShip)
		}
		// TODO be more precise
		return nil
	}))
	return chg
}
