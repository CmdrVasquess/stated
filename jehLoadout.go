package stated

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"git.fractalqb.de/fractalqb/ggja"

	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/CmdrVasquess/stated/ship"
)

var (
	modSizeRegexp  = regexp.MustCompile(`size([\d]+)`)
	modClassRegexp = regexp.MustCompile(`class([\d]+)`)
	modGradeRegexp = regexp.MustCompile(`grade([\d]+)`)
)

func intValRegexp(r *regexp.Regexp, s string) int {
	match := r.FindStringSubmatch(s)
	if match != nil {
		i, _ := strconv.Atoi(match[1])
		return i
	}
	return -1
}

func init() {
	evtHdlrs[journal.LoadoutEvent.String()] = jehLoadout
}

func jehLoadout(ed *EDState, e events.Event) (chg att.Change) {
	ed.MustCommander(journal.LoadoutEvent.String())
	return chg
}

func parseModuleItem(item string) (base string, size int, class int) {
	var baseEnd = len(item)
	match := modSizeRegexp.FindStringSubmatchIndex(item)
	if match != nil {
		num := item[match[2]:match[3]]
		size, _ = strconv.Atoi(num)
		if match[0] < baseEnd {
			baseEnd = match[0]
		}
	}
	match = modClassRegexp.FindStringSubmatchIndex(item)
	if match != nil {
		num := item[match[2]:match[3]]
		class, _ = strconv.Atoi(num)
		if match[0] < baseEnd {
			baseEnd = match[0]
		}
	}
	base = strings.Trim(item[:baseEnd], "_")
	return base, size, class
}

func shipFromLoadout(e *journal.Loadout, types ship.TypeRepo) (ship *ship.Ship, err error) {
	shty := types.Get(e.Ship)
	if shty == nil {
		return nil, fmt.Errorf("unknown ship type '%s'", e.Ship)
	}
	ship = shty.NewShip(nil)
	ship.Name = e.ShipName
	ship.Ident = e.ShipIdent
	ship.Armour, err = armourFromLoadout(e.Slot("Armour"))
	if err != nil {
		return ship, err
	}
	if err = coreModsFromLoadout(ship, e); err != nil {
		return ship, err
	}
	// TODO set ship details
	ship.Cargo = e.CargoCapacity
	ship.MaxJump = e.MaxJumpRange
	return ship, nil
}

func armourFromLoadout(mod *journal.ShipModule) (ship.Armour, error) {
	res := ship.Armour{
		Alloy: ship.Alloy(intValRegexp(modGradeRegexp, mod.Item) - 1),
	}
	if res.Alloy < 0 {
		return res, fmt.Errorf("unknown alloy '%s'", mod.Item)
	}
	return res, nil
}

func engnrFromLoadout(emod *journal.ShipModule) (res *ship.Engineering, err error) {
	jobj := &ggja.Obj{Bare: emod.Bare}
	jobj = jobj.Obj("Engineering")
	if jobj == nil {
		return nil, nil
	}
	jobj.OnError = ggja.SetError{&err}.OnError
	res = &ship.Engineering{
		Blueprint: jobj.MStr("BlueprintName"),
		Level:     jobj.MInt("Level"),
		Quality:   jobj.MF32("Quality"),
	}
	return res, err
}

func coreModsFromLoadout(s *ship.Ship, e *journal.Loadout) error {
	count := 0
	var cmod ship.CoreSlot
	for i := range e.Modules {
		emod := &e.Modules[i]
		switch emod.Slot {
		case "PowerPlant":
			cmod = ship.PowerPlant
		case "MainEngines":
			cmod = ship.Thrusters
		case "FrameShiftDrive":
			cmod = ship.FSD
		case "LifeSupport":
			cmod = ship.LifeSupport
		case "PowerDistributor":
			cmod = ship.PowerDitsr
		case "Radar":
			cmod = ship.Sensors
		case "FuelTank":
			cmod = ship.FuelTank
		default:
			continue
		}
		_, sz, cls := parseModuleItem(emod.Item)
		engnr, err := engnrFromLoadout(emod)
		if err != nil {
			return err
		}
		s.CoreModules[cmod] = ship.CoreModule{
			Module: ship.Module{
				Size:  int8(sz),
				Class: int8(cls),
				Engnr: engnr,
			},
			Type: cmod,
		}
		count++
		if count == 7 {
			break
		}
	}
	if count != 7 {
		return fmt.Errorf("unexpected number of core modules: %d", count)
	}
	return nil
}
