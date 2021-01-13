package ship

import (
	"encoding/json"
)

type Manufacturer string

const (
	CoreDynamics   Manufacturer = "Core Dynamics"
	FaulconDeLacy  Manufacturer = "Faulcon DeLacy"
	Gutamaya       Manufacturer = "Gutamaya"
	Lakon          Manufacturer = "Lakon"
	SaudKruger     Manufacturer = "Saud Kruger"
	ZorgonPeterson Manufacturer = "Zorgon Peterson"
)

//go:generate stringer -type Size
type Size int

const (
	SmallShip Size = (iota + 1)
	MediumShip
	LargeShip
)

//go:generate stringer -type CoreModule
type CoreModule int

const (
	PowerPlant CoreModule = iota
	Thrusters
	FSD
	LifeSupport
	PowerDitsr
	Sensors
	FuelTank
)

type CoreSlotsSpec [FuelTank + 1]int8

func (css CoreSlotsSpec) MarshalJSON() ([]byte, error) {
	tmp := make(map[string]int8)
	for m := PowerPlant; m <= FuelTank; m++ {
		tmp[m.String()] = css[m]
	}
	return json.Marshal(tmp)
}

func (css *CoreSlotsSpec) UnmarshalJSON(data []byte) error {
	tmp := make(map[string]int8)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	for m := PowerPlant; m <= FuelTank; m++ {
		(*css)[m] = tmp[m.String()]
	}
	return nil
}

//go:generate stringer -type HardpointSize
type HardpointSize int

const (
	Utility HardpointSize = iota
	SmallWeapon
	MediumWeapon
	LargeWeapon
	HugeWeapon
)

//go:generate stringer -type OptSlotRestriction
type OptSlotRestriction int

const (
	OptAll OptSlotRestriction = iota
	OptMilitary
)

type OptSlot struct {
	Size     int8
	Count    int8
	Restrict OptSlotRestriction `json:",omitempty"`
}

type Type struct {
	Name       string
	Manf       Manufacturer
	Size       Size
	Cost       int
	Crew       int8
	CoreSlots  CoreSlotsSpec
	Hardpoints [HugeWeapon + 1]int8
	OptSlots   []OptSlot
	Fighter    bool
}
