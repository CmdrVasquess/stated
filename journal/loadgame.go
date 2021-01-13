package journal

import "github.com/CmdrVasquess/stated/events"

type loadgameT string

const LoadGameEvent = loadgameT("LoadGame")

func (t loadgameT) New() events.Event { return new(LoadGame) }
func (t loadgameT) String() string    { return string(t) }

type LoadGame struct {
	events.Common
	Commander string
	FID       string
	Horizons  bool
}

func (_ *LoadGame) EventType() events.Type { return LoadGameEvent }

func init() {
	events.MustRegisterType(string(LoadGameEvent), LoadGameEvent)
}
