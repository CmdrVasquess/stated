package stated

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CmdrVasquess/stated/ship"

	"github.com/CmdrVasquess/stated/journal"
)

func ExampleParseModuleItem() {
	const example = "int_powerplant_size3_class5"
	base, size, class := parseModuleItem(example)
	fmt.Println(base, size, class)
	// Output:
	// int_powerplant 3 5
}

// func ExampleShipFromLoadout() {
// 	var je journal.Loadout
// 	json.Unmarshal([]byte(testNugget), &je)
// 	fmt.Println(&je)
// 	types := ship.FsTypeRepo{Dir: "ship"}
// 	ship, err := shipFromLoadout(&je, &types)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else if raw, err := json.Marshal(ship); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		os.Stdout.Write(raw)
// 	}
// 	// Output:
// 	// ...
// }

const testNugget = `{
  "timestamp": "2021-01-23T20:15:58Z",
  "event": "Loadout",
  "Ship": "asp",
  "ShipID": 3,
  "ShipName": "Nugget",
  "ShipIdent": "jv-ax1",
  "HullValue": 5521649,
  "ModulesValue": 36374809,
  "HullHealth": 1.0,
  "UnladenMass": 324.357971,
  "CargoCapacity": 16,
  "MaxJumpRange": 68.463455,
  "FuelCapacity": {
    "Main": 32.0,
    "Reserve": 0.63
  },
  "Rebuy": 1571119,
  "Modules": [
    {
      "Slot": "ShipCockpit",
      "Item": "asp_cockpit",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "CargoHatch",
      "Item": "modularcargobaydoor",
      "On": true,
      "Priority": 2,
      "Health": 1.0
    },
    {
      "Slot": "TinyHardpoint1",
      "Item": "hpt_heatsinklauncher_turret_tiny",
      "On": true,
      "Priority": 0,
      "AmmoInClip": 1,
      "AmmoInHopper": 2,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "Ram Tah",
        "EngineerID": 300110,
        "BlueprintID": 128731471,
        "BlueprintName": "HeatSinkLauncher_LightWeight",
        "Level": 1,
        "Quality": 0.0,
        "Modifiers": [
          {
            "Label": "Mass",
            "Value": 0.787459,
            "OriginalValue": 1.3,
            "LessIsGood": 1
          },
          {
            "Label": "Integrity",
            "Value": 35.367107,
            "OriginalValue": 45.0,
            "LessIsGood": 0
          },
          {
            "Label": "PowerDraw",
            "Value": 0.182056,
            "OriginalValue": 0.2,
            "LessIsGood": 1
          }
        ]
      }
    },
    {
      "Slot": "TinyHardpoint2",
      "Item": "hpt_chafflauncher_tiny",
      "On": true,
      "Priority": 0,
      "AmmoInClip": 1,
      "AmmoInHopper": 10,
      "Health": 1.0
    },
    {
      "Slot": "TinyHardpoint3",
      "Item": "hpt_chafflauncher_tiny",
      "On": true,
      "Priority": 0,
      "AmmoInClip": 1,
      "AmmoInHopper": 10,
      "Health": 1.0
    },
    {
      "Slot": "TinyHardpoint4",
      "Item": "hpt_plasmapointdefence_turret_tiny",
      "On": true,
      "Priority": 0,
      "AmmoInClip": 12,
      "AmmoInHopper": 10000,
      "Health": 1.0
    },
    {
      "Slot": "PaintJob",
      "Item": "paintjob_asp_metallic_gold",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Armour",
      "Item": "asp_armour_grade1",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "PowerPlant",
      "Item": "int_powerplant_size3_class5",
      "On": true,
      "Priority": 1,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "Marco Qwent",
        "EngineerID": 300200,
        "BlueprintID": 128673765,
        "BlueprintName": "PowerPlant_Boosted",
        "Level": 1,
        "Quality": 1.0,
        "Modifiers": [
          {
            "Label": "Integrity",
            "Value": 66.5,
            "OriginalValue": 70.0,
            "LessIsGood": 0
          },
          {
            "Label": "PowerCapacity",
            "Value": 13.440001,
            "OriginalValue": 12.0,
            "LessIsGood": 0
          },
          {
            "Label": "HeatEfficiency",
            "Value": 0.42,
            "OriginalValue": 0.4,
            "LessIsGood": 1
          }
        ]
      }
    },
    {
      "Slot": "MainEngines",
      "Item": "int_engine_size4_class2",
      "On": true,
      "Priority": 0,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "Professor Palin",
        "EngineerID": 300220,
        "BlueprintID": 128673669,
        "BlueprintName": "Engine_Tuned",
        "Level": 5,
        "Quality": 0.234,
        "ExperimentalEffect": "special_engine_lightweight",
        "ExperimentalEffect_Localised": "Stripped Down",
        "Modifiers": [
          {
            "Label": "Mass",
            "Value": 3.6,
            "OriginalValue": 4.0,
            "LessIsGood": 1
          },
          {
            "Label": "Integrity",
            "Value": 53.760002,
            "OriginalValue": 64.0,
            "LessIsGood": 0
          },
          {
            "Label": "PowerDraw",
            "Value": 4.2804,
            "OriginalValue": 3.69,
            "LessIsGood": 1
          },
          {
            "Label": "EngineOptimalMass",
            "Value": 283.5,
            "OriginalValue": 315.0,
            "LessIsGood": 0
          },
          {
            "Label": "EngineOptPerformance",
            "Value": 124.290016,
            "OriginalValue": 100.0,
            "LessIsGood": 0
          },
          {
            "Label": "EngineHeatRate",
            "Value": 0.61958,
            "OriginalValue": 1.3,
            "LessIsGood": 1
          }
        ]
      }
    },
    {
      "Slot": "FrameShiftDrive",
      "Item": "int_hyperdrive_size5_class5",
      "On": true,
      "Priority": 0,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "Elvira Martuuk",
        "EngineerID": 300160,
        "BlueprintID": 128673694,
        "BlueprintName": "FSD_LongRange",
        "Level": 5,
        "Quality": 1.0,
        "Modifiers": [
          {
            "Label": "Mass",
            "Value": 26.0,
            "OriginalValue": 20.0,
            "LessIsGood": 1
          },
          {
            "Label": "Integrity",
            "Value": 102.0,
            "OriginalValue": 120.0,
            "LessIsGood": 0
          },
          {
            "Label": "PowerDraw",
            "Value": 0.69,
            "OriginalValue": 0.6,
            "LessIsGood": 1
          },
          {
            "Label": "FSDOptimalMass",
            "Value": 1627.5,
            "OriginalValue": 1050.0,
            "LessIsGood": 0
          }
        ]
      }
    },
    {
      "Slot": "LifeSupport",
      "Item": "int_lifesupport_size4_class2",
      "On": true,
      "Priority": 0,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "Bill Turner",
        "EngineerID": 300010,
        "BlueprintID": 128731493,
        "BlueprintName": "LifeSupport_LightWeight",
        "Level": 3,
        "Quality": 0.0,
        "Modifiers": [
          {
            "Label": "Mass",
            "Value": 1.961555,
            "OriginalValue": 4.0,
            "LessIsGood": 1
          },
          {
            "Label": "Integrity",
            "Value": 53.320938,
            "OriginalValue": 72.0,
            "LessIsGood": 0
          }
        ]
      }
    },
    {
      "Slot": "PowerDistributor",
      "Item": "int_powerdistributor_size3_class2",
      "On": true,
      "Priority": 0,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "The Dweller",
        "EngineerID": 300180,
        "BlueprintID": 128673742,
        "BlueprintName": "PowerDistributor_PriorityEngines",
        "Level": 3,
        "Quality": 0.6029,
        "Modifiers": [
          {
            "Label": "WeaponsCapacity",
            "Value": 16.379999,
            "OriginalValue": 18.0,
            "LessIsGood": 0
          },
          {
            "Label": "WeaponsRecharge",
            "Value": 2.037,
            "OriginalValue": 2.1,
            "LessIsGood": 0
          },
          {
            "Label": "EnginesCapacity",
            "Value": 19.153399,
            "OriginalValue": 14.0,
            "LessIsGood": 0
          },
          {
            "Label": "EnginesRecharge",
            "Value": 1.2722,
            "OriginalValue": 1.0,
            "LessIsGood": 0
          },
          {
            "Label": "SystemsCapacity",
            "Value": 12.74,
            "OriginalValue": 14.0,
            "LessIsGood": 0
          },
          {
            "Label": "SystemsRecharge",
            "Value": 0.91,
            "OriginalValue": 1.0,
            "LessIsGood": 0
          }
        ]
      }
    },
    {
      "Slot": "Radar",
      "Item": "int_sensors_size5_class2",
      "On": true,
      "Priority": 0,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "Lei Cheung",
        "EngineerID": 300120,
        "BlueprintID": 128740673,
        "BlueprintName": "Sensor_Sensor_LightWeight",
        "Level": 5,
        "Quality": 0.0,
        "Modifiers": [
          {
            "Label": "Mass",
            "Value": 1.788959,
            "OriginalValue": 8.0,
            "LessIsGood": 1
          },
          {
            "Label": "Integrity",
            "Value": 24.948948,
            "OriginalValue": 77.0,
            "LessIsGood": 0
          },
          {
            "Label": "PowerDraw",
            "Value": 0.362113,
            "OriginalValue": 0.37,
            "LessIsGood": 1
          },
          {
            "Label": "SensorTargetScanAngle",
            "Value": 21.422466,
            "OriginalValue": 30.0,
            "LessIsGood": 0
          }
        ]
      }
    },
    {
      "Slot": "FuelTank",
      "Item": "int_fueltank_size5_class3",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Decal1",
      "Item": "decal_explorer_elite",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Decal2",
      "Item": "decal_explorer_elite",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Decal3",
      "Item": "decal_explorer_elite",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "ShipName0",
      "Item": "nameplate_practical02_black",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "ShipName1",
      "Item": "nameplate_practical02_black",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "ShipID0",
      "Item": "nameplate_shipid_black",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "ShipID1",
      "Item": "nameplate_shipid_black",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Slot01_Size6",
      "Item": "int_fuelscoop_size6_class5",
      "On": true,
      "Priority": 0,
      "Health": 1.0
    },
    {
      "Slot": "Slot02_Size5",
      "Item": "int_guardianfsdbooster_size5",
      "On": true,
      "Priority": 0,
      "Health": 1.0
    },
    {
      "Slot": "Slot03_Size3",
      "Item": "int_shieldgenerator_size3_class2",
      "On": true,
      "Priority": 0,
      "Health": 1.0,
      "Engineering": {
        "Engineer": "Didi Vatermann",
        "EngineerID": 300000,
        "BlueprintID": 128673827,
        "BlueprintName": "ShieldGenerator_Optimised",
        "Level": 3,
        "Quality": 1.0,
        "Modifiers": [
          {
            "Label": "Mass",
            "Value": 1.32,
            "OriginalValue": 2.0,
            "LessIsGood": 1
          },
          {
            "Label": "Integrity",
            "Value": 32.299999,
            "OriginalValue": 38.0,
            "LessIsGood": 0
          },
          {
            "Label": "PowerDraw",
            "Value": 1.008,
            "OriginalValue": 1.44,
            "LessIsGood": 1
          },
          {
            "Label": "ShieldGenOptimalMass",
            "Value": 158.399994,
            "OriginalValue": 165.0,
            "LessIsGood": 0
          },
          {
            "Label": "ShieldGenStrength",
            "Value": 98.099998,
            "OriginalValue": 90.0,
            "LessIsGood": 0
          }
        ]
      }
    },
    {
      "Slot": "Slot04_Size3",
      "Item": "int_cargorack_size3_class1",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Slot05_Size3",
      "Item": "int_cargorack_size3_class1",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Slot06_Size2",
      "Item": "int_supercruiseassist",
      "On": true,
      "Priority": 0,
      "Health": 1.0
    },
    {
      "Slot": "Slot07_Size2",
      "Item": "int_dockingcomputer_standard",
      "On": true,
      "Priority": 0,
      "Health": 1.0
    },
    {
      "Slot": "Slot08_Size1",
      "Item": "int_detailedsurfacescanner_tiny",
      "On": true,
      "Priority": 0,
      "Health": 1.0
    },
    {
      "Slot": "PlanetaryApproachSuite",
      "Item": "int_planetapproachsuite",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Bobble03",
      "Item": "bobble_planet_mars",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "Bobble08",
      "Item": "bobble_pilotmale",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "ShipKitSpoiler",
      "Item": "asp_shipkit1_spoiler1",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "ShipKitWings",
      "Item": "asp_shipkit1_wings3",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "ShipKitBumper",
      "Item": "asp_shipkit1_bumper1",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "WeaponColour",
      "Item": "weaponcustomisation_red",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "EngineColour",
      "Item": "enginecustomisation_blue",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    },
    {
      "Slot": "VesselVoice",
      "Item": "voicepack_verity",
      "On": true,
      "Priority": 1,
      "Health": 1.0
    }
  ]
}`
