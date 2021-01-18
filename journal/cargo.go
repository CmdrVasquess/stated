package journal

import "github.com/CmdrVasquess/stated/events"

type cargoT string

const CargoEvent = cargoT("Cargo")

func (t cargoT) New() events.Event { return new(Cargo) }
func (t cargoT) String() string    { return string(t) }

type CargoItem struct {
	Name    string
	NameL7d string `json:"Name_Localised"`
	Count   int16
	Stolen  int16
}

type Cargo struct {
	events.Common
	Vessel    string
	Count     int16
	Inventory []CargoItem
}

func (_ *Cargo) EventType() events.Type { return CargoEvent }

func init() {
	events.RegisterType(string(CargoEvent), CargoEvent)
}
