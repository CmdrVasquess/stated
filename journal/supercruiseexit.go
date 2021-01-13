package journal

import "github.com/CmdrVasquess/stated/events"

type supercruiseexitT string

const SupercruiseExitEvent = supercruiseexitT("SupercruiseExit")

func (t supercruiseexitT) New() events.Event { return new(SupercruiseExit) }
func (t supercruiseexitT) String() string    { return string(t) }

type SupercruiseExit struct {
	events.Common
	StarSystem    string
	SystemAddress uint64
	Body          string
	BodyID        int
	BodyType      string
}

func (_ *SupercruiseExit) EventType() events.Type { return SupercruiseExitEvent }

func init() {
	events.MustRegisterType(string(SupercruiseExitEvent), SupercruiseExitEvent)
}
