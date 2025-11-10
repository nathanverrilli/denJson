package denJson

import (
	"github.com/shopspring/decimal"
	"time"
)

type CdrData struct {
	Data       []Cdr `json:"data"`
	StatusCode int   `json:"statusCode"`
	Timestamp  time.Time
}

type Cdr struct {
	Id                   string           `json:"id"`
	StartDateTime        time.Time        `json:"start_date_time"`
	StopDateTime         time.Time        `json:"stop_date_time"`
	AuthID               string           `json:"auth_id"`
	AuthMethod           string           `json:"auth_method"`
	AuthorizationID      string           `json:"authorization_id"`
	Location             Location         `json:"location"`
	Tariffs              []Tariff         `json:"tariffs"`
	ChargingPeriods      []ChargingPeriod `json:"charging_periods"`
	TotalCost            decimal.Decimal  `json:"total_cost"`
	TotalEnergy          decimal.Decimal  `json:"total_energy"`
	LastUpdated          time.Time        `json:"last_updated"`
	TotalTime            decimal.Decimal  `json:"total_time"`
	TotalParkingTime     decimal.Decimal  `json:"total_parking_time"`
	Taxes                []CdrTax         `json:"taxes"`
	TotalTax             decimal.Decimal  `json:"total_tax"`
	AveragePower         decimal.Decimal  `json:"average_power"`
	FundingSource        FundingSource    `json:"funding_source"`
	MaxPower             decimal.Decimal  `json:"max_power"`
	PreTaxTotal          decimal.Decimal  `json:"pre_tax_total"`
	StationTotalCost     decimal.Decimal  `json:"station_total_cost"`
	RoamingPartnerRawCdr string           `json:"roaming_partner_raw_cdr"`
	Currency             string           `json:"currency"`
	StationCurrency      string           `json:"station_currency"`
}

type CdrTax struct {
	Description string          `json:"description"`
	Amount      decimal.Decimal `json:"amount"`
}
type FundingSource struct {
	SourceType  string `json:"source_type"`
	PartnerName string `json:"partner_name"`
}
