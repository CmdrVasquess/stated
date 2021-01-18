package ship

type Engineering struct {
	Blueprint string
	Level     int
	Quality   float64
	Effect    string
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
	Engi  *Engineering
}

type Module struct {
	Type  string
	Size  int8
	Class int8
	Mod   *Engineering
}

type CoreModulesSpec [FuelTank + 1]Module

type OptModule struct {
	Module
	Restriction OptSlotRestriction
}

type Tool struct {
	Type  string
	Size  HardpointSize
	Class int8
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
	MaxJump     float64
}
