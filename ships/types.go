package ships

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Manufacturer string

const (
	CoreDynamics   Manufacturer = "Core Dynamics"
	FaulconDeLacy  Manufacturer = "Faulcon DeLacy"
	Gutamaya       Manufacturer = "Gutamaya"
	Lakon          Manufacturer = "Lakon Spaceways"
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

//go:generate stringer -type CoreSlot
type CoreSlot int

func (cs CoreSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(cs.String())
}

func (cs *CoreSlot) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for i := 0; i < len(_CoreSlot_index); i++ {
		if tmp := CoreSlot(i); tmp.String() == s {
			*cs = tmp
			return nil
		}
	}
	return fmt.Errorf("unknown core slot '%s'", s)
}

const (
	PowerPlant CoreSlot = iota
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

func (hps HardpointSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(hps.String())
}

func (hps *HardpointSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for i := 0; i < len(_HardpointSize_index); i++ {
		if tmp := HardpointSize(i); tmp.String() == s {
			*hps = tmp
			return nil
		}
	}
	return fmt.Errorf("unknown hardpoint size '%s'", s)
}

const (
	Utility HardpointSize = iota
	SmallWeapon
	MediumWeapon
	LargeWeapon
	HugeWeapon
)

//go:generate stringer -type OptSlotRestriction
type OptSlotRestriction int

func (osr OptSlotRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(osr.String())
}

func (osr *OptSlotRestriction) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for i := 0; i < len(_OptSlotRestriction_index); i++ {
		if tmp := OptSlotRestriction(i); tmp.String() == s {
			*osr = tmp
			return nil
		}
	}
	return fmt.Errorf("unknown opt-slot restriction '%s'", s)
}

const (
	OptAll OptSlotRestriction = iota
	OptMilitary
)

type OptSlots struct {
	Size     int8
	Count    int8
	Restrict OptSlotRestriction `json:",omitempty"`
}

type Type struct {
	Tag        string `json:"-"`
	Name       string
	Manf       Manufacturer
	Size       Size
	Cost       int
	Crew       int8
	CoreSlots  CoreSlotsSpec
	Hardpoints [HugeWeapon + 1]int8
	OptSlots   []OptSlots
	Fighter    bool
}

func (st *Type) NewShip(reuse *Ship) *Ship {
	if reuse == nil {
		reuse = new(Ship)
	}
	reuse.Type = TypeRef{
		TypeName: st.Tag,
		Type:     st,
	}
	for i, s := range st.Hardpoints {
		reuse.Tools[i] = make([]*Tool, s)
	}
	reuse.OptModules = make([]*OptModule, st.OptSlotNo())
	return reuse
}

func (st *Type) OptSlot(idx int) *OptSlots {
	sum := 0
	for i := range st.OptSlots {
		def := &st.OptSlots[i]
		sum += int(def.Count)
		if idx < sum {
			return def
		}
	}
	return nil
}

func (st *Type) OptSlotNo() (res int) {
	for _, s := range st.OptSlots {
		res += int(s.Count)
	}
	return res
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
	Get(t string) (*Type, error)
}

type FsTypeRepo struct {
	Dir   string
	cache map[string]*Type
}

func (fsr *FsTypeRepo) Get(t string) (*Type, error) {
	if fsr.cache == nil {
		fsr.cache = make(map[string]*Type)
	}
	res, ok := fsr.cache[t]
	if ok {
		return res, nil
	}
	raw, err := ioutil.ReadFile(filepath.Join(fsr.Dir, t+".json"))
	if err != nil {
		fsr.cache[t] = nil
		return nil, err
	}
	res = new(Type)
	if err = json.Unmarshal(raw, res); err != nil {
		fsr.cache[t] = nil
		return nil, err
	}
	res.Tag = t
	fsr.cache[t] = res
	return res, nil
}
