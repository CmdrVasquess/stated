package ships

type Engineering struct {
	Blueprint string
	Level     int
	Quality   float32
	Effect    string `json:",omitempty"`
}

type Alloy int

const (
	LightweightAlloy Alloy = iota
	ReinforcedAlloay
	MilitaryCompoite
	MirroredComposite
	ReactiveComposite
)

type Armour struct {
	Alloy Alloy
	Engnr *Engineering `json:",omitempty"`
}

type Module struct {
	Size  int8
	Class int8
	Engnr *Engineering `json:",omitempty"`
}

type CoreModule struct {
	Module
	Type CoreSlot
}

type CoreModulesSpec [FuelTank + 1]CoreModule

type OptModule struct {
	Module
	Type        string
	Restriction OptSlotRestriction `json:",omitempty"`
}

type Tool struct {
	Type  string
	Size  HardpointSize
	Class int8
	Engnr *Engineering `json:",omitempty"`
}

type Ship struct {
	ID          int
	Type        TypeRef
	Name        string
	Ident       string
	Armour      Armour
	CoreModules CoreModulesSpec
	Tools       [HugeWeapon + 1][]*Tool
	OptModules  []*OptModule
	Cargo       int
	MaxJump     float32
}
