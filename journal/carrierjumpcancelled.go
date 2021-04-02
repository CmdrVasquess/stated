package journal

import "github.com/CmdrVasquess/stated/events"

type carrierjumpcancelledT string

const CarrierJumpCancelledEvent = carrierjumpcancelledT("CarrierJumpCancelled")

func (t carrierjumpcancelledT) New() events.Event { return new(CarrierJumpCancelled) }
func (t carrierjumpcancelledT) String() string    { return string(t) }

type CarrierJumpCancelled struct {
	events.Common
	CarrierID int64
}

func (_ *CarrierJumpCancelled) EventType() events.Type { return CarrierJumpCancelledEvent }

func init() {
	events.RegisterType(string(CarrierJumpCancelledEvent), CarrierJumpCancelledEvent)
}
