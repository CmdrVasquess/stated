package journal

import "github.com/CmdrVasquess/stated/events"

type dockedT string

const DockedEvent = dockedT("Docked")

func (t dockedT) New() events.Event { return new(Docked) }
func (t dockedT) String() string    { return string(t) }

type Docked struct {
	events.Common
	SystemAddress  uint64
	StarSystem     string
	MarketID       int64
	StationName    string
	StationType    string
	DistFromStarLS float64
}

func (_ *Docked) EventType() events.Type { return DockedEvent }

func init() {
	events.MustRegisterType(string(DockedEvent), DockedEvent)
}
