package denJson

import (
	"time"

	"github.com/shopspring/decimal"
)

type Empty struct{}
type TariffData struct {
	Data       []Tariff  `json:"data"`
	StatusCode int       `json:"status_code"`
	Timestamp  time.Time `json:"timestamp"`
}

type Tariff struct {
	ID            string          `json:"id,omitempty"`
	Currency      string          `json:"currency,omitempty"`
	Elements      []TariffElement `json:"elements,omitempty"`
	TariffAltText []AltText       `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  string          `json:"tariff_alt_url,omitempty"`
	EnergyMix     EnergyMix       `json:"energy_mix"`
	LastUpdated   time.Time       `json:"last_updated"`
	/* DEN extensions */
	DriverGroups []DriverGroup `json:"driver_groups,omitempty"`
	Max          MinMax        `json:"max"`
	Min          MinMax        `json:"min"`
	Locations    []string      `json:"locations,omitempty"`
}

type MinMax struct {
	Price    decimal.Decimal `json:"price"`
	StepSize int             `json:"step_size"`
}

type AltText struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

type TariffElement struct {
	PriceComponents []PriceComponent `json:"price_components"`
	Restrictions    Restrictions     `json:"restrictions"`
}

type PriceComponent struct {
	Type     string          `json:"type"`
	Price    decimal.Decimal `json:"price"`
	StepSize decimal.Decimal `json:"step_size"`
}

type Restrictions struct {
	StartTime string          `json:"start_time,omitempty"`
	EndTime   string          `json:"end_time,omitempty"`
	StartDate string          `json:"start_date,omitempty"`
	EndDate   string          `json:"end_date,omitempty"`
	MinKwh    decimal.Decimal `json:"min_kwh"`
	MaxKwh    decimal.Decimal `json:"max_kwh"`
	MinPower  decimal.Decimal `json:"min_power"`
	MaxPower  decimal.Decimal `json:"max_power"`
	DayOfWeek []string        `json:"day_of_week,omitempty"`
	// MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY
	MaxDuration            int `json:"max_duration,omitempty"`
	MinDuration            int `json:"min_duration,omitempty"`
	MinCongestion          int `json:"min_congestion,omitempty"`
	MinCongestionThreshold int `json:"min_congestion_threshold,omitempty"`
	MinVehicleSOC          int `json:"min_vehicle_soc,omitempty"`
}
