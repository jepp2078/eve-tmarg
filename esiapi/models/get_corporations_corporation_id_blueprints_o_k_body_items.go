// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDBlueprintsOKBodyItems get_corporations_corporation_id_blueprints_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdBlueprintsOKBodyItems
type GetCorporationsCorporationIDBlueprintsOKBodyItems struct {

	// get_corporations_corporation_id_blueprints_item_id
	//
	// Unique ID for this item.
	// Required: true
	ItemID *int64 `json:"item_id"`

	// get_corporations_corporation_id_blueprints_location_flag
	//
	// Type of the location_id
	// Required: true
	// Enum: [AssetSafety AutoFit Bonus Booster BoosterBay Capsule Cargo CorpDeliveries CorpSAG1 CorpSAG2 CorpSAG3 CorpSAG4 CorpSAG5 CorpSAG6 CorpSAG7 CrateLoot Deliveries DroneBay DustBattle DustDatabank FighterBay FighterTube0 FighterTube1 FighterTube2 FighterTube3 FighterTube4 FleetHangar Hangar HangarAll HiSlot0 HiSlot1 HiSlot2 HiSlot3 HiSlot4 HiSlot5 HiSlot6 HiSlot7 HiddenModifiers Implant Impounded JunkyardReprocessed JunkyardTrashed LoSlot0 LoSlot1 LoSlot2 LoSlot3 LoSlot4 LoSlot5 LoSlot6 LoSlot7 Locked MedSlot0 MedSlot1 MedSlot2 MedSlot3 MedSlot4 MedSlot5 MedSlot6 MedSlot7 OfficeFolder Pilot PlanetSurface QuafeBay Reward RigSlot0 RigSlot1 RigSlot2 RigSlot3 RigSlot4 RigSlot5 RigSlot6 RigSlot7 SecondaryStorage ServiceSlot0 ServiceSlot1 ServiceSlot2 ServiceSlot3 ServiceSlot4 ServiceSlot5 ServiceSlot6 ServiceSlot7 ShipHangar ShipOffline Skill SkillInTraining SpecializedAmmoHold SpecializedCommandCenterHold SpecializedFuelBay SpecializedGasHold SpecializedIndustrialShipHold SpecializedLargeShipHold SpecializedMaterialBay SpecializedMediumShipHold SpecializedMineralHold SpecializedOreHold SpecializedPlanetaryCommoditiesHold SpecializedSalvageHold SpecializedShipHold SpecializedSmallShipHold StructureActive StructureFuel StructureInactive StructureOffline SubSystemBay SubSystemSlot0 SubSystemSlot1 SubSystemSlot2 SubSystemSlot3 SubSystemSlot4 SubSystemSlot5 SubSystemSlot6 SubSystemSlot7 Unlocked Wallet Wardrobe]
	LocationFlag *string `json:"location_flag"`

	// get_corporations_corporation_id_blueprints_location_id
	//
	// References a solar system, station or item_id if this blueprint is located within a container.
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_corporations_corporation_id_blueprints_material_efficiency
	//
	// Material Efficiency Level of the blueprint.
	// Required: true
	// Maximum: 25
	// Minimum: 0
	MaterialEfficiency *int32 `json:"material_efficiency"`

	// get_corporations_corporation_id_blueprints_quantity
	//
	// A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet).
	// Required: true
	// Minimum: -2
	Quantity *int32 `json:"quantity"`

	// get_corporations_corporation_id_blueprints_runs
	//
	// Number of runs remaining if the blueprint is a copy, -1 if it is an original.
	// Required: true
	// Minimum: -1
	Runs *int32 `json:"runs"`

	// get_corporations_corporation_id_blueprints_time_efficiency
	//
	// Time Efficiency Level of the blueprint.
	// Required: true
	// Maximum: 20
	// Minimum: 0
	TimeEfficiency *int32 `json:"time_efficiency"`

	// get_corporations_corporation_id_blueprints_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporations corporation Id blueprints o k body items
func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaterialEfficiency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeEfficiency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AssetSafety","AutoFit","Bonus","Booster","BoosterBay","Capsule","Cargo","CorpDeliveries","CorpSAG1","CorpSAG2","CorpSAG3","CorpSAG4","CorpSAG5","CorpSAG6","CorpSAG7","CrateLoot","Deliveries","DroneBay","DustBattle","DustDatabank","FighterBay","FighterTube0","FighterTube1","FighterTube2","FighterTube3","FighterTube4","FleetHangar","Hangar","HangarAll","HiSlot0","HiSlot1","HiSlot2","HiSlot3","HiSlot4","HiSlot5","HiSlot6","HiSlot7","HiddenModifiers","Implant","Impounded","JunkyardReprocessed","JunkyardTrashed","LoSlot0","LoSlot1","LoSlot2","LoSlot3","LoSlot4","LoSlot5","LoSlot6","LoSlot7","Locked","MedSlot0","MedSlot1","MedSlot2","MedSlot3","MedSlot4","MedSlot5","MedSlot6","MedSlot7","OfficeFolder","Pilot","PlanetSurface","QuafeBay","Reward","RigSlot0","RigSlot1","RigSlot2","RigSlot3","RigSlot4","RigSlot5","RigSlot6","RigSlot7","SecondaryStorage","ServiceSlot0","ServiceSlot1","ServiceSlot2","ServiceSlot3","ServiceSlot4","ServiceSlot5","ServiceSlot6","ServiceSlot7","ShipHangar","ShipOffline","Skill","SkillInTraining","SpecializedAmmoHold","SpecializedCommandCenterHold","SpecializedFuelBay","SpecializedGasHold","SpecializedIndustrialShipHold","SpecializedLargeShipHold","SpecializedMaterialBay","SpecializedMediumShipHold","SpecializedMineralHold","SpecializedOreHold","SpecializedPlanetaryCommoditiesHold","SpecializedSalvageHold","SpecializedShipHold","SpecializedSmallShipHold","StructureActive","StructureFuel","StructureInactive","StructureOffline","SubSystemBay","SubSystemSlot0","SubSystemSlot1","SubSystemSlot2","SubSystemSlot3","SubSystemSlot4","SubSystemSlot5","SubSystemSlot6","SubSystemSlot7","Unlocked","Wallet","Wardrobe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum = append(getCorporationsCorporationIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagAssetSafety captures enum value "AssetSafety"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagAssetSafety string = "AssetSafety"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagAutoFit captures enum value "AutoFit"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagAutoFit string = "AutoFit"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagBonus captures enum value "Bonus"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagBonus string = "Bonus"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagBooster captures enum value "Booster"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagBooster string = "Booster"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagBoosterBay captures enum value "BoosterBay"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagBoosterBay string = "BoosterBay"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCapsule captures enum value "Capsule"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCapsule string = "Capsule"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCargo captures enum value "Cargo"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCargo string = "Cargo"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpDeliveries captures enum value "CorpDeliveries"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpDeliveries string = "CorpDeliveries"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG1 captures enum value "CorpSAG1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG1 string = "CorpSAG1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG2 captures enum value "CorpSAG2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG2 string = "CorpSAG2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG3 captures enum value "CorpSAG3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG3 string = "CorpSAG3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG4 captures enum value "CorpSAG4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG4 string = "CorpSAG4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG5 captures enum value "CorpSAG5"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG5 string = "CorpSAG5"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG6 captures enum value "CorpSAG6"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG6 string = "CorpSAG6"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG7 captures enum value "CorpSAG7"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCorpSAG7 string = "CorpSAG7"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCrateLoot captures enum value "CrateLoot"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagCrateLoot string = "CrateLoot"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDeliveries captures enum value "Deliveries"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDeliveries string = "Deliveries"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDroneBay captures enum value "DroneBay"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDroneBay string = "DroneBay"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDustBattle captures enum value "DustBattle"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDustBattle string = "DustBattle"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDustDatabank captures enum value "DustDatabank"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagDustDatabank string = "DustDatabank"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterBay captures enum value "FighterBay"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterBay string = "FighterBay"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube0 captures enum value "FighterTube0"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube0 string = "FighterTube0"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube1 captures enum value "FighterTube1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube1 string = "FighterTube1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube2 captures enum value "FighterTube2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube2 string = "FighterTube2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube3 captures enum value "FighterTube3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube3 string = "FighterTube3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube4 captures enum value "FighterTube4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFighterTube4 string = "FighterTube4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFleetHangar captures enum value "FleetHangar"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagFleetHangar string = "FleetHangar"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHangar captures enum value "Hangar"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHangar string = "Hangar"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHangarAll captures enum value "HangarAll"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHangarAll string = "HangarAll"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot0 captures enum value "HiSlot0"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot0 string = "HiSlot0"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot1 captures enum value "HiSlot1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot1 string = "HiSlot1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot2 captures enum value "HiSlot2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot2 string = "HiSlot2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot3 captures enum value "HiSlot3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot3 string = "HiSlot3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot4 captures enum value "HiSlot4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot4 string = "HiSlot4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot5 captures enum value "HiSlot5"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot5 string = "HiSlot5"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot6 captures enum value "HiSlot6"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot6 string = "HiSlot6"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot7 captures enum value "HiSlot7"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiSlot7 string = "HiSlot7"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiddenModifiers captures enum value "HiddenModifiers"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagHiddenModifiers string = "HiddenModifiers"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagImplant captures enum value "Implant"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagImplant string = "Implant"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagImpounded captures enum value "Impounded"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagImpounded string = "Impounded"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagJunkyardReprocessed captures enum value "JunkyardReprocessed"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagJunkyardReprocessed string = "JunkyardReprocessed"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagJunkyardTrashed captures enum value "JunkyardTrashed"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagJunkyardTrashed string = "JunkyardTrashed"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot0 captures enum value "LoSlot0"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot0 string = "LoSlot0"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot1 captures enum value "LoSlot1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot1 string = "LoSlot1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot2 captures enum value "LoSlot2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot2 string = "LoSlot2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot3 captures enum value "LoSlot3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot3 string = "LoSlot3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot4 captures enum value "LoSlot4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot4 string = "LoSlot4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot5 captures enum value "LoSlot5"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot5 string = "LoSlot5"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot6 captures enum value "LoSlot6"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot6 string = "LoSlot6"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot7 captures enum value "LoSlot7"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLoSlot7 string = "LoSlot7"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLocked captures enum value "Locked"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagLocked string = "Locked"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot0 captures enum value "MedSlot0"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot0 string = "MedSlot0"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot1 captures enum value "MedSlot1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot1 string = "MedSlot1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot2 captures enum value "MedSlot2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot2 string = "MedSlot2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot3 captures enum value "MedSlot3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot3 string = "MedSlot3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot4 captures enum value "MedSlot4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot4 string = "MedSlot4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot5 captures enum value "MedSlot5"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot5 string = "MedSlot5"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot6 captures enum value "MedSlot6"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot6 string = "MedSlot6"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot7 captures enum value "MedSlot7"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagMedSlot7 string = "MedSlot7"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagOfficeFolder captures enum value "OfficeFolder"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagOfficeFolder string = "OfficeFolder"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagPilot captures enum value "Pilot"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagPilot string = "Pilot"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagPlanetSurface captures enum value "PlanetSurface"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagPlanetSurface string = "PlanetSurface"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagQuafeBay captures enum value "QuafeBay"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagQuafeBay string = "QuafeBay"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagReward captures enum value "Reward"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagReward string = "Reward"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot0 captures enum value "RigSlot0"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot0 string = "RigSlot0"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot1 captures enum value "RigSlot1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot1 string = "RigSlot1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot2 captures enum value "RigSlot2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot2 string = "RigSlot2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot3 captures enum value "RigSlot3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot3 string = "RigSlot3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot4 captures enum value "RigSlot4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot4 string = "RigSlot4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot5 captures enum value "RigSlot5"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot5 string = "RigSlot5"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot6 captures enum value "RigSlot6"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot6 string = "RigSlot6"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot7 captures enum value "RigSlot7"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagRigSlot7 string = "RigSlot7"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSecondaryStorage captures enum value "SecondaryStorage"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSecondaryStorage string = "SecondaryStorage"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot0 captures enum value "ServiceSlot0"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot0 string = "ServiceSlot0"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot1 captures enum value "ServiceSlot1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot1 string = "ServiceSlot1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot2 captures enum value "ServiceSlot2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot2 string = "ServiceSlot2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot3 captures enum value "ServiceSlot3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot3 string = "ServiceSlot3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot4 captures enum value "ServiceSlot4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot4 string = "ServiceSlot4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot5 captures enum value "ServiceSlot5"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot5 string = "ServiceSlot5"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot6 captures enum value "ServiceSlot6"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot6 string = "ServiceSlot6"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot7 captures enum value "ServiceSlot7"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagServiceSlot7 string = "ServiceSlot7"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagShipHangar captures enum value "ShipHangar"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagShipHangar string = "ShipHangar"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagShipOffline captures enum value "ShipOffline"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagShipOffline string = "ShipOffline"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSkill captures enum value "Skill"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSkill string = "Skill"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSkillInTraining captures enum value "SkillInTraining"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSkillInTraining string = "SkillInTraining"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedAmmoHold captures enum value "SpecializedAmmoHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedAmmoHold string = "SpecializedAmmoHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedCommandCenterHold captures enum value "SpecializedCommandCenterHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedCommandCenterHold string = "SpecializedCommandCenterHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedFuelBay captures enum value "SpecializedFuelBay"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedFuelBay string = "SpecializedFuelBay"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedGasHold captures enum value "SpecializedGasHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedGasHold string = "SpecializedGasHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedIndustrialShipHold captures enum value "SpecializedIndustrialShipHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedIndustrialShipHold string = "SpecializedIndustrialShipHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedLargeShipHold captures enum value "SpecializedLargeShipHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedLargeShipHold string = "SpecializedLargeShipHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedMaterialBay captures enum value "SpecializedMaterialBay"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedMaterialBay string = "SpecializedMaterialBay"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedMediumShipHold captures enum value "SpecializedMediumShipHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedMediumShipHold string = "SpecializedMediumShipHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedMineralHold captures enum value "SpecializedMineralHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedMineralHold string = "SpecializedMineralHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedOreHold captures enum value "SpecializedOreHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedOreHold string = "SpecializedOreHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedPlanetaryCommoditiesHold captures enum value "SpecializedPlanetaryCommoditiesHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedPlanetaryCommoditiesHold string = "SpecializedPlanetaryCommoditiesHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedSalvageHold captures enum value "SpecializedSalvageHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedSalvageHold string = "SpecializedSalvageHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedShipHold captures enum value "SpecializedShipHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedShipHold string = "SpecializedShipHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedSmallShipHold captures enum value "SpecializedSmallShipHold"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSpecializedSmallShipHold string = "SpecializedSmallShipHold"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureActive captures enum value "StructureActive"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureActive string = "StructureActive"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureFuel captures enum value "StructureFuel"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureFuel string = "StructureFuel"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureInactive captures enum value "StructureInactive"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureInactive string = "StructureInactive"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureOffline captures enum value "StructureOffline"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagStructureOffline string = "StructureOffline"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemBay captures enum value "SubSystemBay"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemBay string = "SubSystemBay"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot0 captures enum value "SubSystemSlot0"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot0 string = "SubSystemSlot0"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot1 captures enum value "SubSystemSlot1"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot1 string = "SubSystemSlot1"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot2 captures enum value "SubSystemSlot2"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot2 string = "SubSystemSlot2"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot3 captures enum value "SubSystemSlot3"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot3 string = "SubSystemSlot3"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot4 captures enum value "SubSystemSlot4"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot4 string = "SubSystemSlot4"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot5 captures enum value "SubSystemSlot5"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot5 string = "SubSystemSlot5"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot6 captures enum value "SubSystemSlot6"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot6 string = "SubSystemSlot6"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot7 captures enum value "SubSystemSlot7"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot7 string = "SubSystemSlot7"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagUnlocked captures enum value "Unlocked"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagUnlocked string = "Unlocked"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagWallet captures enum value "Wallet"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagWallet string = "Wallet"

	// GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagWardrobe captures enum value "Wardrobe"
	GetCorporationsCorporationIDBlueprintsOKBodyItemsLocationFlagWardrobe string = "Wardrobe"
)

// prop value enum
func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateLocationFlagEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateLocationFlag(formats strfmt.Registry) error {

	if err := validate.Required("location_flag", "body", m.LocationFlag); err != nil {
		return err
	}

	// value enum
	if err := m.validateLocationFlagEnum("location_flag", "body", *m.LocationFlag); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateMaterialEfficiency(formats strfmt.Registry) error {

	if err := validate.Required("material_efficiency", "body", m.MaterialEfficiency); err != nil {
		return err
	}

	if err := validate.MinimumInt("material_efficiency", "body", int64(*m.MaterialEfficiency), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("material_efficiency", "body", int64(*m.MaterialEfficiency), 25, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.MinimumInt("quantity", "body", int64(*m.Quantity), -2, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateRuns(formats strfmt.Registry) error {

	if err := validate.Required("runs", "body", m.Runs); err != nil {
		return err
	}

	if err := validate.MinimumInt("runs", "body", int64(*m.Runs), -1, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateTimeEfficiency(formats strfmt.Registry) error {

	if err := validate.Required("time_efficiency", "body", m.TimeEfficiency); err != nil {
		return err
	}

	if err := validate.MinimumInt("time_efficiency", "body", int64(*m.TimeEfficiency), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("time_efficiency", "body", int64(*m.TimeEfficiency), 20, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDBlueprintsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDBlueprintsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
