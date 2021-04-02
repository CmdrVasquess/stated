package journal

import "github.com/CmdrVasquess/stated/events"

type carrierjumpT string

const CarrierJumpEvent = carrierjumpT("CarrierJump")

func (t carrierjumpT) New() events.Event { return new(CarrierJump) }
func (t carrierjumpT) String() string    { return string(t) }

type CarrierJump struct {
	events.Common
	StationName   string
	MarketID      int64 // Same as CarrierID
	StarSystem    string
	SystemAddress uint64
	StarPos       [3]float32
	Body          string
	BodyID        int
	BodyType      string
}

func (_ *CarrierJump) EventType() events.Type { return CarrierJumpEvent }

func init() {
	events.RegisterType(string(CarrierJumpEvent), CarrierJumpEvent)
}
