package journal

import "github.com/CmdrVasquess/stated/events"

type loadoutT string

const LoadoutEvent = loadoutT("Loadout")

func (t loadoutT) New() events.Event { return new(Loadout) }
func (t loadoutT) String() string    { return string(t) }

type Loadout struct {
	events.Common
	Ship          string
	ShipID        int
	ShipName      string
	ShipIdent     string
	MaxJumpRange  float32
	CargoCapacity int
}

func (_ *Loadout) EventType() events.Type { return LoadoutEvent }

func init() {
	events.MustRegisterType(string(LoadoutEvent), LoadoutEvent)
}
