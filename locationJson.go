package denjson

import (
	"time"
)

// multiple location response from API
type LocationData struct {
	Data       []Location `json:"data,omitempty"`
	StatusCode int        `json:"status_code,omitempty"`
	Timestamp  time.Time  `json:"timestamp,omitempty"`
}

// OCPI 2.1.1 section 8.4.13
type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// OCPI 2.1.1 section 8.3.3
type Connector struct {
	ID                 string    `json:"id"`
	Standard           string    `json:"standard"`
	Format             string    `json:"format"`
	PowerType          string    `json:"power_type"`
	Voltage            int       `json:"voltage"`
	Amperage           int       `json:"amperage"`
	TariffID           string    `json:"tariff_id,omitempty"`
	TermsAndConditions string    `json:"terms_and_conditions,omitempty"`
	LastUpdated        time.Time `json:"last_updated"`
	// non-OCPI for DEN(///
	MaxPower         float64 `json:"max_power,omitempty"`
	MaxSessionLength int     `json:"max_session_length,omitempty"`
	AvailablePower   float64 `json:"available_power,omitempty"`
	PowerStatus      string  `json:"power_status,omitempty"`
}

// OCPI 2.1.1 section 8.3.2
type Evses struct {
	UID                string           `json:"uid"`
	EvseID             string           `json:"evse_id,omitempty"`
	PartnerEvseID      string           `json:"partner_evse_id,omitempty"`
	Status             string           `json:"status"`
	StatusSchedule     []StatusSchedule `json:"status_schedule,omitempty"`
	Capabilities       []string         `json:"capabilities,omitempty"`
	Connectors         []Connector      `json:"connectors"`
	FloorLevel         string           `json:"floor_level,omitempty"`
	Coordinates        GeoLocation      `json:"coordinates,omitempty"`
	PhysicalReference  string           `json:"physical_reference,omitempty"`
	Directions         []DisplayText    `json:"directions,omitempty"`
	ParkingRestriction []string         `json:"parking_restriction,omitempty"`
	Images             []Image          `json:"images,omitempty"`
	LastUpdated        time.Time        `json:"last_updated"`
}

// OCPI 2.1.1 section 14.3
type DisplayText struct {
	Language string `json:"language"` // ISO 639-1
	Text     string `json:"text"`
}

// OCPI 2.1.1 section 8.4.22
type StatusSchedule struct {
	PeriodBegin time.Time `json:"period_begin"`
	PeriodEnd   time.Time `json:"period_end,omitempty"`
	Status      string    `json:"status"`
}

// OCPI 2.1.1 section 8.4.2
type BusinessDetails struct {
	Name    string `json:"name"`
	Website string `json:"website,omitempty"`
	Logo    Image  `json:"logo,omitempty"`
}

// OCPI 2.1.1 section 8.4.11
type ExceptionalPeriod struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

// OCPI 2.1.1 section 8.4.20
type RegularHours struct {
	Weekday     int    `json:"weekday"`
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

// OCPI 2.1.1 section 8.4.14
type Hours struct {
	// EITHER `RegularHours` OR `Twentyfourseven` should be present, not both
	Twentyfourseven     bool                `json:"twentyfourseven,omitempty"`
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
	ExceptionalOpenings []ExceptionalPeriod `json:"exceptional_openings,omitempty"`
	ExceptionalClosings []ExceptionalPeriod `json:"exceptional_closings,omitempty"`
}

// DEN
type DriverGroup struct {
	ID           int    `json:"id,omitempty"`
	TariffID     string `json:"tariff_id,omitempty"`
	OpeningTimes Hours  `json:"opening_times,omitempty"`
}

// DEN
type Hotline struct {
	Phone string `json:"phone,omitempty"`
	Name  string `json:"name,omitempty"`
}

// OCPI 2.1.1
type AdditionalGeoLocation struct {
	Latitude  string      `json:"latitude"`
	Longitude string      `json:"longitude"`
	Name      DisplayText `json:"name,omitempty"`
}

// OCPI 2.1.1
type Location struct {
	ID                 string                `json:"id"`
	Type               string                `json:"type"`
	Name               string                `json:"name,omitempty"`
	Address            string                `json:"address"`
	City               string                `json:"city"`
	PostalCode         string                `json:"postal_code"`
	Country            string                `json:"country"`
	Coordinates        GeoLocation           `json:"coordinates"`
	RelatedLocations   AdditionalGeoLocation `json:"related_locations,omitempty"`
	Evses              []Evses               `json:"evses,omitempty"`
	Directions         []DisplayText         `json:"directions,omitempty"`
	Operator           BusinessDetails       `json:"operator,omitempty"`
	Suboperator        BusinessDetails       `json:"suboperator,omitempty"`
	Owner              BusinessDetails       `json:"owner,omitempty"`
	Facilities         []string              `json:"facilities,omitempty"`
	TimeZone           string                `json:"time_zone,omitempty"`
	OpeningTimes       Hours                 `json:"opening_times,omitempty"`
	ChargingWhenClosed bool                  `json:"charging_when_closed,omitempty"`
	Images             []Image               `json:"images,omitempty"`
	EnergyMix          EnergyMix             `json:"energy_mix,omitempty"`
	LastUpdated        time.Time             `json:"last_updated"`
	// Extended DEN parameters/values/information
	Province               string        `json:"province,omitempty"`                 // non-OCPI
	DriverGroups           []DriverGroup `json:"driver_groups,omitempty"`            // non-OCPI
	AuthIDType             string        `json:"auth_id_type,omitempty"`             // non-OCPI
	PaymentMethods         []string      `json:"payment_methods,omitempty"`          // non-OCPI
	Hotline                Hotline       `json:"hotline,omitempty"`                  // non-OCPI
	PluginDependency       string        `json:"plugin_dependency,omitempty"`        // non-OCPI
	AccessRestriction      string        `json:"access_restriction,omitempty"`       // non-OCPI
	UsageStatus            string        `json:"usage_status,omitempty"`             // non-OCPI
	IsPublic               bool          `json:"is_public,omitempty"`                // non-OCPI
	StationGroupId         int64         `json:"station_group_id,omitempty"`         // non-OCPI
	Tax                    []Tax         `json:"tax,omitempty"`                      // non-OCPI
	ChargingAccessibility  string        `json:"charging_accessibility,omitempty"`   // non-OCPI
	ReservableDriverGroups []int         `json:"reservable_driver_groups,omitempty"` // non-OCPI
}

type Tax struct {
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	AppliedOn string  `json:"applied_on"`
}

// OCPI 2.1.1
type EnergyMix struct {
	IsGreenEnergy     bool                  `json:"is_green_energy"`
	EnergySources     []EnergySource        `json:"energy_sources,omitempty"`
	EnvironImpact     []EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName      string                `json:"supplier_name,omitempty"`
	EnergyProductName string                `json:"energy_product_name,omitempty"`
}

// OCPI 2.1.1 section 8.4.8
type EnvironmentalImpact struct {
	Source string  `json:"source"`
	Amount float64 `json:"amount"`
}

// OCPI 2.1.1 section 8.4.7
type EnergySource struct {
	Source     string  `json:"source"`
	Percentage float64 `json:"float"`
}

// OCPI 2.1.1 section 8.4.15
type Image struct {
	Url       string `json:"url,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Category  string `json:"category,omitempty"`
	Type      string `json:"type,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
}
