package journal

import "github.com/CmdrVasquess/stated/events"

type approachsettlementT string

const ApproachSettlementEvent = approachsettlementT("ApproachSettlement")

func (t approachsettlementT) New() events.Event { return new(ApproachSettlement) }
func (t approachsettlementT) String() string    { return string(t) }

type ApproachSettlement struct {
	events.Common
	SystemAddress uint64
	Body          string
	BodyID        int
	Name          string
	Latitude      float32
	Longitude     float32
}

func (_ *ApproachSettlement) EventType() events.Type { return ApproachSettlementEvent }

func init() {
	events.MustRegisterType(string(ApproachSettlementEvent), ApproachSettlementEvent)
}
