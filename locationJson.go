package denJson

import (
	"time"

	"github.com/shopspring/decimal"
)

// multiple location response from API
type LocationData struct {
	Data       []Location `json:"data,omitempty"`
	StatusCode int        `json:"status_code,omitempty"`
	Timestamp  time.Time  `json:"timestamp,omitempty"`
}

// OCPI 2.1.1 section 8.4.13
type GeoLocation struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

// OCPI 2.1.1 section 8.3.3
type Connector struct {
	ID                 string    `json:"id,omitempty"`
	Standard           string    `json:"standard,omitempty"`
	Format             string    `json:"format,omitempty"`
	PowerType          string    `json:"power_type,omitempty"`
	Voltage            int       `json:"voltage,omitempty"`
	Amperage           int       `json:"amperage,omitempty"`
	TariffID           string    `json:"tariff_id,omitempty"`
	TermsAndConditions string    `json:"terms_and_conditions,omitempty"`
	LastUpdated        time.Time `json:"last_updated,omitempty"`
	// non-OCPI for DEN(///
	MaxPower         decimal.Decimal `json:"max_power,omitempty"`
	MaxSessionLength int             `json:"max_session_length,omitempty"`
	AvailablePower   decimal.Decimal `json:"available_power,omitempty"`
	PowerStatus      string          `json:"power_status,omitempty"`
	MaxElectricPower decimal.Decimal `json:"max_electric_power,omitempty"`
}

// OCPI 2.1.1 section 8.3.2
type Evses struct {
	UID                string           `json:"uid,omitempty"`
	EvseID             string           `json:"evse_id,omitempty"`
	PartnerEvseID      string           `json:"partner_evse_id,omitempty"`
	Status             string           `json:"status,omitempty"`
	StatusSchedule     []StatusSchedule `json:"status_schedule,omitempty"`
	Capabilities       []string         `json:"capabilities,omitempty"`
	Connectors         []Connector      `json:"connectors,omitempty"`
	FloorLevel         string           `json:"floor_level,omitempty"`
	Coordinates        GeoLocation      `json:"coordinates,omitempty"`
	PhysicalReference  string           `json:"physical_reference,omitempty"`
	Directions         []DisplayText    `json:"directions,omitempty"`
	ParkingRestriction []string         `json:"parking_restriction,omitempty"`
	Images             []Image          `json:"images,omitempty"`
	LastUpdated        time.Time        `json:"last_updated,omitempty"`
}

// OCPI 2.1.1 section 14.3
type DisplayText struct {
	Language string `json:"language,omitempty"` // ISO 639-1
	Text     string `json:"text,omitempty"`
}

// OCPI 2.1.1 section 8.4.22
type StatusSchedule struct {
	PeriodBegin time.Time `json:"period_begin,omitempty"`
	PeriodEnd   time.Time `json:"period_end,omitempty"`
	Status      string    `json:"status,omitempty"`
}

// OCPI 2.1.1 section 8.4.2
type BusinessDetails struct {
	Name    string `json:"name,omitempty"`
	Website string `json:"website,omitempty"`
	Logo    Image  `json:"logo,omitempty"`
}

// OCPI 2.1.1 section 8.4.11
type ExceptionalPeriod struct {
	PeriodBegin string `json:"period_begin,omitempty"`
	PeriodEnd   string `json:"period_end,omitempty"`
}

// OCPI 2.1.1 section 8.4.20
type RegularHours struct {
	Weekday     int    `json:"weekday,omitempty"`
	PeriodBegin string `json:"period_begin,omitempty"`
	PeriodEnd   string `json:"period_end,omitempty"`
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
	Latitude  string      `json:"latitude,omitempty"`
	Longitude string      `json:"longitude,omitempty"`
	Name      DisplayText `json:"name,omitempty,omitempty"`
}

// OCPI 2.1.1
type Location struct {
	ID                 string                `json:"id,omitempty"`
	Type               string                `json:"type,omitempty"`
	Name               string                `json:"name,omitempty"`
	Address            string                `json:"address,omitempty"`
	City               string                `json:"city,omitempty"`
	PostalCode         string                `json:"postal_code,omitempty"`
	Country            string                `json:"country,omitempty"`
	Coordinates        GeoLocation           `json:"coordinates,omitempty"`
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
	LastUpdated        time.Time             `json:"last_updated,omitempty"`
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
	Name      string          `json:"name,omitempty"`
	Amount    decimal.Decimal `json:"amount,omitempty"`
	AppliedOn string          `json:"applied_on,omitempty"`
}

// OCPI 2.1.1
type EnergyMix struct {
	IsGreenEnergy     bool                  `json:"is_green_energy,omitempty"`
	EnergySources     []EnergySource        `json:"energy_sources,omitempty"`
	EnvironImpact     []EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName      string                `json:"supplier_name,omitempty"`
	EnergyProductName string                `json:"energy_product_name,omitempty"`
}

// OCPI 2.1.1 section 8.4.8
type EnvironmentalImpact struct {
	Source string          `json:"source,omitempty"`
	Amount decimal.Decimal `json:"amount,omitempty"`
}

// OCPI 2.1.1 section 8.4.7
type EnergySource struct {
	Source     string          `json:"source,omitempty"`
	Percentage decimal.Decimal `json:"percentage,omitempty"`
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
