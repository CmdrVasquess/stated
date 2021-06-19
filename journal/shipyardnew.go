package journal

import "github.com/CmdrVasquess/stated/events"

type shipyardnewT string

const ShipyardNewEvent = shipyardnewT("ShipyardNew")

func (t shipyardnewT) New() events.Event { return new(ShipyardNew) }
func (t shipyardnewT) String() string    { return string(t) }

type ShipyardNew struct {
	events.Common
	ShipType    string
	ShipTypeL7d string `json:"ShipType_Localised"`
	NewShipID   int
}

func (_ *ShipyardNew) EventType() events.Type { return ShipyardNewEvent }

func init() {
	events.MustRegisterType(string(ShipyardNewEvent), ShipyardNewEvent)
}
