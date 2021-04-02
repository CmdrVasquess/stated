package journal

import "github.com/CmdrVasquess/stated/events"

type carrierjumprequestT string

const CarrierJumpRequestEvent = carrierjumprequestT("CarrierJumpRequest")

func (t carrierjumprequestT) New() events.Event { return new(CarrierJumpRequest) }
func (t carrierjumprequestT) String() string    { return string(t) }

type CarrierJumpRequest struct {
	events.Common
	CarrierID     int64
	SystemName    string
	SystemAddress uint64
	Body          string
	BodyID        int
}

func (_ *CarrierJumpRequest) EventType() events.Type { return CarrierJumpRequestEvent }

func init() {
	events.RegisterType(string(CarrierJumpRequestEvent), CarrierJumpRequestEvent)
}
