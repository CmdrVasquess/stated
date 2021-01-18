package ship

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
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

func (st *Type) NewShip(reuse *Ship) *Ship {
	if reuse == nil {
		reuse = new(Ship)
	}
	reuse.Type = TypeRef{
		TypeName: st.Name,
		Type:     st,
	}
	for i, s := range st.Hardpoints {
		reuse.Tools[i] = make([]*Tool, s)
	}
	reuse.OptModules = make([]*OptModule, len(st.OptSlots))
	return reuse
}

type TypeRef struct {
	TypeName string
	Type     *Type
}

func (tr TypeRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(tr.TypeName)
}

func (tr *TypeRef) UnmarshalJSON(data []byte) error {
	tr.Type = nil
	return json.Unmarshal(data, &tr.TypeName)
}

type TypeRepo interface {
	Get(t string) *Type
}

type FsTypeRepo struct {
	Dir   string
	cache map[string]*Type
}

func (fsr *FsTypeRepo) Get(t string) *Type {
	res, ok := fsr.cache[t]
	if ok {
		return res
	}
	raw, err := ioutil.ReadFile(filepath.Join(fsr.Dir, t+".json"))
	if err != nil {
		fsr.cache[t] = nil
		return nil
	}
	res = new(Type)
	if err = json.Unmarshal(raw, res); err != nil {
		fsr.cache[t] = nil
		return nil
	}
	fsr.cache[t] = res
	return res
}
