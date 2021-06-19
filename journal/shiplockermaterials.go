package journal

import "github.com/CmdrVasquess/stated/events"

type shiplockermaterialsT string

const ShipLockerMaterialsEvent = shiplockermaterialsT("ShipLockerMaterials")

func (t shiplockermaterialsT) New() events.Event { return new(ShipLockerMaterials) }
func (t shiplockermaterialsT) String() string    { return string(t) }

type ShipLockerMaterials struct {
	events.Common
	Items       []SLMaterial
	Components  []SLMaterial
	Consumables []SLMaterial
	Data        []SLMaterial
}

type SLMaterial struct {
	Name    string
	NameL7d string `json:"Name_Localised"`
	OwnerID int
	Count   int
}

func (_ *ShipLockerMaterials) EventType() events.Type { return ShipLockerMaterialsEvent }

func init() {
	events.RegisterType(string(ShipLockerMaterialsEvent), ShipLockerMaterialsEvent)
}
