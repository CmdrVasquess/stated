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
	"github.com/CmdrVasquess/stated/ships"
)

var (
	modSizeRegexp   = regexp.MustCompile(`size([\d]+)`)
	modClassRegexp  = regexp.MustCompile(`class([\d]+)`)
	modGradeRegexp  = regexp.MustCompile(`grade([\d]+)`)
	modHardptRegexp = regexp.MustCompile(`^([[:alpha:]]+)Hardpoint([\d]+)$`)
	modOptRegexp    = regexp.MustCompile(`^Slot([\d]+)_Size([\d]+)`)
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
	evt := e.(*journal.Loadout)
	ShipFromLoadout(evt, nil) // TODO Where to get ShipTypeRepos from
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
	} else {
		size = -1
	}
	match = modClassRegexp.FindStringSubmatchIndex(item)
	if match != nil {
		num := item[match[2]:match[3]]
		class, _ = strconv.Atoi(num)
		if match[0] < baseEnd {
			baseEnd = match[0]
		}
	} else {
		class = -1
	}
	base = strings.Trim(item[:baseEnd], "_")
	return base, size, class
}

func parseOptSlot(slotstr string) (slot, size int) {
	match := modOptRegexp.FindStringSubmatch(slotstr)
	if match == nil {
		return -1, -1
	}
	slot, _ = strconv.Atoi(match[1])
	size, _ = strconv.Atoi(match[2])
	return slot, size
}

func ShipFromLoadout(e *journal.Loadout, types ships.TypeRepo) (ship *ships.Ship, err error) {
	shty, err := types.Get(e.Ship)
	switch {
	case err != nil:
		return nil, err
	case shty == nil:
		return nil, fmt.Errorf("unknown ship type '%s'", e.Ship)
	}
	ship = shty.NewShip(nil)
	ship.ID = e.ShipID
	ship.Name = e.ShipName
	ship.Ident = e.ShipIdent
	ship.Armour, err = armourFromLoadout(e.Slot("Armour"))
	if err != nil {
		return ship, err
	}
	if err = coreModsFromLoadout(ship, e); err != nil {
		return ship, err
	}
	if err = toolsFromLoadout(ship, e); err != nil {
		return ship, err
	}
	if err = optModsFromLoadout(ship, e); err != nil {
		return ship, err
	}
	ship.Cargo = e.CargoCapacity
	ship.MaxJump = e.MaxJumpRange
	return ship, nil
}

func optModsFromLoadout(s *ships.Ship, e *journal.Loadout) error {
	slotIdxOffset := -1
	for i := range e.Modules {
		emod := &e.Modules[i]
		slot, ssz := parseOptSlot(emod.Slot)
		switch {
		case slot >= 0 || ssz >= 0:
			if slot == 0 { // Type-9 got an extra size 8 slot with index 0
				if slotIdxOffset == 0 {
					return fmt.Errorf("2nd slot with index 0: %s", emod.Slot)
				}
				copy(s.OptModules[1:], s.OptModules)
				slotIdxOffset = 0
			}
			item, msz, class := parseModuleItem(emod.Item)
			if sdef := s.Type.Type.OptSlot(slot + slotIdxOffset); sdef == nil {
				return fmt.Errorf("no definition for module slot '%s' (%s/%s)",
					emod.Slot,
					emod.Slot,
					emod.Item)
			} else if ssz != int(sdef.Size) {
				return fmt.Errorf(
					"module's slot size %d differs from definition %d (%s/%s)",
					ssz,
					sdef.Size,
					emod.Slot,
					emod.Item)
			} else if msz > int(sdef.Size) {
				return fmt.Errorf("module size %d exceeds slot size %d (%s/%s)",
					msz,
					sdef.Size,
					emod.Slot,
					emod.Item)
			}
			engnr, err := engnrFromLoadout(emod)
			if err != nil {
				return err
			}
			s.OptModules[slot+slotIdxOffset] = &ships.OptModule{
				Module: ships.Module{
					Size:  int8(msz),
					Class: int8(class),
					Engnr: engnr,
				},
				Type: item,
			}
		case strings.HasPrefix(emod.Item, "Military"):
			if err := militaryFromLoadout(s, emod); err != nil {
				return err
			}
		}
	}
	return nil
}

func militaryFromLoadout(s *ships.Ship, emod *journal.ShipModule) error {
	// TODO military modules
	return nil
}

func toolsFromLoadout(s *ships.Ship, e *journal.Loadout) error {
	for i := range e.Modules {
		emod := &e.Modules[i]
		match := modHardptRegexp.FindStringSubmatch(emod.Slot)
		if match == nil {
			continue
		}
		size := match[1]
		hpno, _ := strconv.Atoi(match[2])
		var hpsz ships.HardpointSize = -1
		switch size {
		case "Tiny":
			hpsz = ships.Utility
		case "Small":
			hpsz = ships.SmallWeapon
		case "Medium":
			hpsz = ships.MediumWeapon
		case "Large":
			hpsz = ships.LargeWeapon
		case "Huge":
			hpsz = ships.HugeWeapon
		}
		item, msz, class := parseModuleItem(emod.Item)
		if msz >= 0 && msz != int(hpsz) {
			return fmt.Errorf("module size %d of '%s' does not match hard-point size %d",
				msz,
				emod.Item,
				hpsz)
		}
		engnr, err := engnrFromLoadout(emod)
		if err != nil {
			return err
		}
		s.Tools[hpsz][hpno-1] = &ships.Tool{
			Type:  item,
			Size:  hpsz,
			Class: int8(class),
			Engnr: engnr,
		}
	}
	return nil
}

func armourFromLoadout(mod *journal.ShipModule) (ships.Armour, error) {
	res := ships.Armour{
		Alloy: ships.Alloy(intValRegexp(modGradeRegexp, mod.Item) - 1),
	}
	if res.Alloy < 0 {
		return res, fmt.Errorf("unknown alloy '%s'", mod.Item)
	}
	engnr, err := engnrFromLoadout(mod)
	if err == nil {
		res.Engnr = engnr
	}
	return res, err
}

func engnrFromLoadout(emod *journal.ShipModule) (res *ships.Engineering, err error) {
	jobj := &ggja.Obj{Bare: emod.Bare}
	jobj = jobj.Obj("Engineering")
	if jobj == nil {
		return nil, nil
	}
	jobj.OnError = ggja.SetError{Target: &err}.OnError
	res = &ships.Engineering{
		Blueprint: jobj.MStr("BlueprintName"),
		Level:     jobj.MInt("Level"),
		Quality:   jobj.MF32("Quality"),
	}
	return res, err
}

func coreModsFromLoadout(s *ships.Ship, e *journal.Loadout) error {
	count := 0
	var cmod ships.CoreSlot
	for i := range e.Modules {
		emod := &e.Modules[i]
		switch emod.Slot {
		case "PowerPlant":
			cmod = ships.PowerPlant
		case "MainEngines":
			cmod = ships.Thrusters
		case "FrameShiftDrive":
			cmod = ships.FSD
		case "LifeSupport":
			cmod = ships.LifeSupport
		case "PowerDistributor":
			cmod = ships.PowerDitsr
		case "Radar":
			cmod = ships.Sensors
		case "FuelTank":
			cmod = ships.FuelTank
		default:
			continue
		}
		_, sz, cls := parseModuleItem(emod.Item)
		engnr, err := engnrFromLoadout(emod)
		if err != nil {
			return err
		}
		s.CoreModules[cmod] = ships.CoreModule{
			Module: ships.Module{
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
