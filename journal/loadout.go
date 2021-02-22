package journal

import (
	"encoding/json"

	"git.fractalqb.de/fractalqb/ggja"
	"github.com/CmdrVasquess/stated/events"
)

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
	Modules       []ShipModule
}

func (l *Loadout) Slot(named string) *ShipModule {
	for i := range l.Modules {
		m := &l.Modules[i]
		if m.Slot == named {
			return m
		}
	}
	return nil
}

type ShipModule struct {
	Slot string
	Item string
	Bare ggja.BareObj
}

func (m *ShipModule) UnmarshalJSON(data []byte) (err error) {
	var bare ggja.BareObj
	if err = json.Unmarshal(data, &bare); err != nil {
		return err
	}
	m.Bare = bare
	obj := ggja.Obj{Bare: bare, OnError: ggja.SetError{&err}.OnError}
	m.Slot = obj.MStr("Slot")
	if err != nil {
		return err
	}
	m.Item = obj.MStr("Item")
	return err
}

func (_ *Loadout) EventType() events.Type { return LoadoutEvent }

func init() {
	events.MustRegisterType(string(LoadoutEvent), LoadoutEvent)
}
