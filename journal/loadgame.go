package journal

import "github.com/CmdrVasquess/stated/events"

type loadgameT string

const LoadGameEvent = loadgameT("LoadGame")

func (t loadgameT) New() events.Event { return new(LoadGame) }
func (t loadgameT) String() string    { return string(t) }

type LoadGame struct {
	events.Common
	FID       string
	Commander string
	Horizons  bool
	Odyssey   bool
	Ship      string
	ShipL7d   string `json:"Ship_Localised"`
	ShipID    int
	ShipName  string
	ShipIdent string
}

func (_ *LoadGame) EventType() events.Type { return LoadGameEvent }

func init() {
	events.MustRegisterType(string(LoadGameEvent), LoadGameEvent)
}
