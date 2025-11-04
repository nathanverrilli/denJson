package denjson

import (
	"time"

	"github.com/shopspring/decimal"
)

type SessionData struct {
	Data       []Session `json:"data"`
	StatusCode int64     `json:"statuscode"`
	Timestamp  string    `json:"timestamp"`
}

type Session struct {
	AuthId          string           `json:"auth_id"`
	AuthMethod      string           `json:"auth_method"`
	Currency        string           `json:"currency"`
	Id              string           `json:"id"`
	Kwh             decimal.Decimal  `json:"kwh"`
	LastUpdated     time.Time        `json:"last_updated"`
	Location        Location         `json:"location"`
	Power           decimal.Decimal  `json:"power"`
	Status          string           `json:"status"`
	TotalCost       decimal.Decimal  `json:"total_cost"`
	MeterId         string           `json:"meter_id"`
	ChargingPeriods []ChargingPeriod `json:"charging_periods"`
	StartDateTime   time.Time        `json:"start_datetime"`
}

type ChargingPeriod struct {
	StartDateTime time.Time `json:"start_date_time"`
	Dimensions    []CdrDimension
}

type CdrDimension struct {
	Type   string          `json:"type"`
	Volume decimal.Decimal `json:"volume"`
}
