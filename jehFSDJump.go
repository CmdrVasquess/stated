package stated

import (
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"
)

func init() {
	evtHdlrs[journal.FSDJumpEvent.String()] = jehFSDJump
}

func jehFSDJump(ed *EDState, e events.Event) (chg change.Flags) {
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
		if cmdr.inShip != nil && evt.JumpDist > cmdr.inShip.MaxJump.Get() {
			chg |= cmdr.inShip.MaxJump.Set(evt.JumpDist, ChgShip)
		}
		// TODO be more precise
		return nil
	}))
	return chg
}
