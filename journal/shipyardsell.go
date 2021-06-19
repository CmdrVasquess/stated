package journal

import "github.com/CmdrVasquess/stated/events"

type shipyardsellT string

const ShipyardSellEvent = shipyardsellT("ShipyardSell")

func (t shipyardsellT) New() events.Event { return new(ShipyardSell) }
func (t shipyardsellT) String() string    { return string(t) }

type ShipyardSell struct {
	events.Common
	SellShipID int
	ShipType   string
	ShipPrice  int64
}

func (_ *ShipyardSell) EventType() events.Type { return ShipyardSellEvent }

func init() {
	events.MustRegisterType(string(ShipyardSellEvent), ShipyardSellEvent)
}
