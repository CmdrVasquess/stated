package journal

import "github.com/CmdrVasquess/stated/events"

type shipyardswapT string

const ShipyardSwapEvent = shipyardswapT("ShipyardSwap")

func (t shipyardswapT) New() events.Event { return new(ShipyardSwap) }
func (t shipyardswapT) String() string    { return string(t) }

type ShipyardSwap struct {
	events.Common
	MarketID    int64
	ShipType    string
	ShipTypeL7d string `json:"ShipType_Localised"`
	ShipID      int
	StoreShipID int
}

func (_ *ShipyardSwap) EventType() events.Type { return ShipyardSwapEvent }

func init() {
	events.MustRegisterType(string(ShipyardSwapEvent), ShipyardSwapEvent)
}
