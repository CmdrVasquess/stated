package journal

import "github.com/CmdrVasquess/stated/events"

type locationT string

const LocationEvent = locationT("Location")

func (t locationT) New() events.Event { return new(Location) }
func (t locationT) String() string    { return string(t) }

type SSysInfo struct {
	SystemAddress       uint64
	StarSystem          string
	StarPos             [3]float32
	SystemAllegiance    string
	SystemEconomy1      string `json:"SystemEconomy"`
	SystemEconomy1L7d   string `json:"SystemEconomy_Localised"`
	SystemEconomy2      string `json:"SystemSecondEconomy"`
	SystemEconomy2L7d   string `json:"SystemSecondEconomy_Localised"`
	SystemGovernment    string
	SystemGovernmentL7d string `json:"SystemGovernment_Localised"`
	SystemSecurity      string
	SystemSecurityL7d   string `json:"SystemSecurity_Localised"`
	Population          int64
}

type Faction struct {
	Name         string
	FactionState string
	Government   string
	Influence    float32
	Allegiance   string
	Happiness    string
	HappinessL7d string `json:"Happiness_Localised"`
	MyReputation float32
}

type Location struct {
	events.Common
	SSysInfo
	DistFromStarLS  float32
	StationType     string
	StationName     string
	MarketID        int64
	StationServices []string
	Docked          bool
	Factions        []*Faction
}

func (_ *Location) EventType() events.Type { return LocationEvent }

func init() {
	events.MustRegisterType(string(LocationEvent), LocationEvent)
}
