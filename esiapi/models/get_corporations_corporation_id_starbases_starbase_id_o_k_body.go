// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDStarbasesStarbaseIDOKBody get_corporations_corporation_id_starbases_starbase_id_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdStarbasesStarbaseIdOKBody
type GetCorporationsCorporationIDStarbasesStarbaseIDOKBody struct {

	// get_corporations_corporation_id_starbases_starbase_id_allow_alliance_members
	//
	// allow_alliance_members boolean
	// Required: true
	AllowAllianceMembers *bool `json:"allow_alliance_members"`

	// get_corporations_corporation_id_starbases_starbase_id_allow_corporation_members
	//
	// allow_corporation_members boolean
	// Required: true
	AllowCorporationMembers *bool `json:"allow_corporation_members"`

	// get_corporations_corporation_id_starbases_starbase_id_anchor
	//
	// Who can anchor starbase (POS) and its structures
	// Required: true
	// Enum: [alliance_member config_starbase_equipment_role corporation_member starbase_fuel_technician_role]
	Anchor *string `json:"anchor"`

	// get_corporations_corporation_id_starbases_starbase_id_attack_if_at_war
	//
	// attack_if_at_war boolean
	// Required: true
	AttackIfAtWar *bool `json:"attack_if_at_war"`

	// get_corporations_corporation_id_starbases_starbase_id_attack_if_other_security_status_dropping
	//
	// attack_if_other_security_status_dropping boolean
	// Required: true
	AttackIfOtherSecurityStatusDropping *bool `json:"attack_if_other_security_status_dropping"`

	// get_corporations_corporation_id_starbases_starbase_id_attack_security_status_threshold
	//
	// Starbase (POS) will attack if target's security standing is lower than this value
	AttackSecurityStatusThreshold float32 `json:"attack_security_status_threshold,omitempty"`

	// get_corporations_corporation_id_starbases_starbase_id_attack_standing_threshold
	//
	// Starbase (POS) will attack if target's standing is lower than this value
	AttackStandingThreshold float32 `json:"attack_standing_threshold,omitempty"`

	// get_corporations_corporation_id_starbases_starbase_id_fuel_bay_take
	//
	// Who can take fuel blocks out of the starbase (POS)'s fuel bay
	// Required: true
	// Enum: [alliance_member config_starbase_equipment_role corporation_member starbase_fuel_technician_role]
	FuelBayTake *string `json:"fuel_bay_take"`

	// get_corporations_corporation_id_starbases_starbase_id_fuel_bay_view
	//
	// Who can view the starbase (POS)'s fule bay. Characters either need to have required role or belong to the starbase (POS) owner's corporation or alliance, as described by the enum, all other access settings follows the same scheme
	// Required: true
	// Enum: [alliance_member config_starbase_equipment_role corporation_member starbase_fuel_technician_role]
	FuelBayView *string `json:"fuel_bay_view"`

	// get_corporations_corporation_id_starbases_starbase_id_fuels
	//
	// Fuel blocks and other things that will be consumed when operating a starbase (POS)
	// Max Items: 20
	Fuels []*GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelsItems `json:"fuels"`

	// get_corporations_corporation_id_starbases_starbase_id_offline
	//
	// Who can offline starbase (POS) and its structures
	// Required: true
	// Enum: [alliance_member config_starbase_equipment_role corporation_member starbase_fuel_technician_role]
	Offline *string `json:"offline"`

	// get_corporations_corporation_id_starbases_starbase_id_online
	//
	// Who can online starbase (POS) and its structures
	// Required: true
	// Enum: [alliance_member config_starbase_equipment_role corporation_member starbase_fuel_technician_role]
	Online *string `json:"online"`

	// get_corporations_corporation_id_starbases_starbase_id_unanchor
	//
	// Who can unanchor starbase (POS) and its structures
	// Required: true
	// Enum: [alliance_member config_starbase_equipment_role corporation_member starbase_fuel_technician_role]
	Unanchor *string `json:"unanchor"`

	// get_corporations_corporation_id_starbases_starbase_id_use_alliance_standings
	//
	// True if the starbase (POS) is using alliance standings, otherwise using corporation's
	// Required: true
	UseAllianceStandings *bool `json:"use_alliance_standings"`
}

// Validate validates this get corporations corporation Id starbases starbase Id o k body
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowAllianceMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowCorporationMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnchor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttackIfAtWar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttackIfOtherSecurityStatusDropping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFuelBayTake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFuelBayView(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFuels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnanchor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseAllianceStandings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateAllowAllianceMembers(formats strfmt.Registry) error {

	if err := validate.Required("allow_alliance_members", "body", m.AllowAllianceMembers); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateAllowCorporationMembers(formats strfmt.Registry) error {

	if err := validate.Required("allow_corporation_members", "body", m.AllowCorporationMembers); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeAnchorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance_member","config_starbase_equipment_role","corporation_member","starbase_fuel_technician_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeAnchorPropEnum = append(getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeAnchorPropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorAllianceMember captures enum value "alliance_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorAllianceMember string = "alliance_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorConfigStarbaseEquipmentRole captures enum value "config_starbase_equipment_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorConfigStarbaseEquipmentRole string = "config_starbase_equipment_role"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorCorporationMember captures enum value "corporation_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorCorporationMember string = "corporation_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorStarbaseFuelTechnicianRole captures enum value "starbase_fuel_technician_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyAnchorStarbaseFuelTechnicianRole string = "starbase_fuel_technician_role"
)

// prop value enum
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateAnchorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeAnchorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateAnchor(formats strfmt.Registry) error {

	if err := validate.Required("anchor", "body", m.Anchor); err != nil {
		return err
	}

	// value enum
	if err := m.validateAnchorEnum("anchor", "body", *m.Anchor); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateAttackIfAtWar(formats strfmt.Registry) error {

	if err := validate.Required("attack_if_at_war", "body", m.AttackIfAtWar); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateAttackIfOtherSecurityStatusDropping(formats strfmt.Registry) error {

	if err := validate.Required("attack_if_other_security_status_dropping", "body", m.AttackIfOtherSecurityStatusDropping); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayTakePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance_member","config_starbase_equipment_role","corporation_member","starbase_fuel_technician_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayTakePropEnum = append(getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayTakePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeAllianceMember captures enum value "alliance_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeAllianceMember string = "alliance_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeConfigStarbaseEquipmentRole captures enum value "config_starbase_equipment_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeConfigStarbaseEquipmentRole string = "config_starbase_equipment_role"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeCorporationMember captures enum value "corporation_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeCorporationMember string = "corporation_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeStarbaseFuelTechnicianRole captures enum value "starbase_fuel_technician_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayTakeStarbaseFuelTechnicianRole string = "starbase_fuel_technician_role"
)

// prop value enum
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateFuelBayTakeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayTakePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateFuelBayTake(formats strfmt.Registry) error {

	if err := validate.Required("fuel_bay_take", "body", m.FuelBayTake); err != nil {
		return err
	}

	// value enum
	if err := m.validateFuelBayTakeEnum("fuel_bay_take", "body", *m.FuelBayTake); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayViewPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance_member","config_starbase_equipment_role","corporation_member","starbase_fuel_technician_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayViewPropEnum = append(getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayViewPropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewAllianceMember captures enum value "alliance_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewAllianceMember string = "alliance_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewConfigStarbaseEquipmentRole captures enum value "config_starbase_equipment_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewConfigStarbaseEquipmentRole string = "config_starbase_equipment_role"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewCorporationMember captures enum value "corporation_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewCorporationMember string = "corporation_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewStarbaseFuelTechnicianRole captures enum value "starbase_fuel_technician_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyFuelBayViewStarbaseFuelTechnicianRole string = "starbase_fuel_technician_role"
)

// prop value enum
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateFuelBayViewEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeFuelBayViewPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateFuelBayView(formats strfmt.Registry) error {

	if err := validate.Required("fuel_bay_view", "body", m.FuelBayView); err != nil {
		return err
	}

	// value enum
	if err := m.validateFuelBayViewEnum("fuel_bay_view", "body", *m.FuelBayView); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateFuels(formats strfmt.Registry) error {

	if swag.IsZero(m.Fuels) { // not required
		return nil
	}

	iFuelsSize := int64(len(m.Fuels))

	if err := validate.MaxItems("fuels", "body", iFuelsSize, 20); err != nil {
		return err
	}

	for i := 0; i < len(m.Fuels); i++ {
		if swag.IsZero(m.Fuels[i]) { // not required
			continue
		}

		if m.Fuels[i] != nil {
			if err := m.Fuels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fuels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOfflinePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance_member","config_starbase_equipment_role","corporation_member","starbase_fuel_technician_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOfflinePropEnum = append(getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOfflinePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineAllianceMember captures enum value "alliance_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineAllianceMember string = "alliance_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineConfigStarbaseEquipmentRole captures enum value "config_starbase_equipment_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineConfigStarbaseEquipmentRole string = "config_starbase_equipment_role"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineCorporationMember captures enum value "corporation_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineCorporationMember string = "corporation_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineStarbaseFuelTechnicianRole captures enum value "starbase_fuel_technician_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOfflineStarbaseFuelTechnicianRole string = "starbase_fuel_technician_role"
)

// prop value enum
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateOfflineEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOfflinePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateOffline(formats strfmt.Registry) error {

	if err := validate.Required("offline", "body", m.Offline); err != nil {
		return err
	}

	// value enum
	if err := m.validateOfflineEnum("offline", "body", *m.Offline); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOnlinePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance_member","config_starbase_equipment_role","corporation_member","starbase_fuel_technician_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOnlinePropEnum = append(getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOnlinePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineAllianceMember captures enum value "alliance_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineAllianceMember string = "alliance_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineConfigStarbaseEquipmentRole captures enum value "config_starbase_equipment_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineConfigStarbaseEquipmentRole string = "config_starbase_equipment_role"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineCorporationMember captures enum value "corporation_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineCorporationMember string = "corporation_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineStarbaseFuelTechnicianRole captures enum value "starbase_fuel_technician_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyOnlineStarbaseFuelTechnicianRole string = "starbase_fuel_technician_role"
)

// prop value enum
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateOnlineEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeOnlinePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateOnline(formats strfmt.Registry) error {

	if err := validate.Required("online", "body", m.Online); err != nil {
		return err
	}

	// value enum
	if err := m.validateOnlineEnum("online", "body", *m.Online); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeUnanchorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance_member","config_starbase_equipment_role","corporation_member","starbase_fuel_technician_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeUnanchorPropEnum = append(getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeUnanchorPropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorAllianceMember captures enum value "alliance_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorAllianceMember string = "alliance_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorConfigStarbaseEquipmentRole captures enum value "config_starbase_equipment_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorConfigStarbaseEquipmentRole string = "config_starbase_equipment_role"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorCorporationMember captures enum value "corporation_member"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorCorporationMember string = "corporation_member"

	// GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorStarbaseFuelTechnicianRole captures enum value "starbase_fuel_technician_role"
	GetCorporationsCorporationIDStarbasesStarbaseIDOKBodyUnanchorStarbaseFuelTechnicianRole string = "starbase_fuel_technician_role"
)

// prop value enum
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateUnanchorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdStarbasesStarbaseIdOKBodyTypeUnanchorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateUnanchor(formats strfmt.Registry) error {

	if err := validate.Required("unanchor", "body", m.Unanchor); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnanchorEnum("unanchor", "body", *m.Unanchor); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) validateUseAllianceStandings(formats strfmt.Registry) error {

	if err := validate.Required("use_alliance_standings", "body", m.UseAllianceStandings); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStarbasesStarbaseIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDStarbasesStarbaseIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
