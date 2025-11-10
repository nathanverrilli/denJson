package denJson

import (
	"time"

	"github.com/shopspring/decimal"
)

type CdrData struct {
	Data       []Cdr `json:"data"`
	StatusCode int   `json:"status_code"`
	Timestamp  time.Time
}

type Cdr struct {
	Id               string           `json:"id"`
	StartDateTime    time.Time        `json:"start_date_time"`
	StopDateTime     time.Time        `json:"stop_date_time"`
	AuthID           string           `json:"auth_id"`
	AuthMethod       string           `json:"auth_method"`
	AuthorizationID  string           `json:"authorization_id"`
	Location         Location         `json:"location"`
	Tariffs          []Tariff         `json:"tariffs"`
	ChargingPeriods  []ChargingPeriod `json:"charging_periods"`
	TotalCost        decimal.Decimal  `json:"total_cost"`
	TotalEnergy      decimal.Decimal  `json:"total_energy"`
	LastUpdated      time.Time        `json:"last_updated"`
	TotalTime        decimal.Decimal  `json:"total_time"`
	TotalParkingTime decimal.Decimal  `json:"total_parking_time"`
	Taxes            []Tax            `json:"taxes"`
	TotalTax         decimal.Decimal  `json:"total_tax"`
	AveragePower     decimal.Decimal  `json:"average_power"`
	FundingSource    FundingSource    `json:"funding_source"`
	MaxPower         decimal.Decimal  `json:"max_power"`
}

type FundingSource struct {
	SourceType string `json:"source_type"`
}
