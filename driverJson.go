package denjson

import "time"

type DriverData struct {
	Data          Driver    `json:"data"`
	StatusCode    int       `json:"status_code"`
	StatusMessage string    `json:"status_message"`
	Timestamp     time.Time `json:"timestamp"`
}

// Driver represents the driver data payload.
type Driver struct {
	DriverID        string           `json:"driver_id"`
	FirstName       string           `json:"first_name"`
	LastName        string           `json:"last_name"`
	DriverGroups    []int64          `json:"driver_groups"`
	AccountState    string           `json:"account_state"`
	Credentials     []Credential     `json:"credentials"`
	BusinessPayment *BusinessPayment `json:"business_payment,omitempty"`
	LastUpdated     time.Time        `json:"last_updated"`
}

// Credential represents a single credential record.
type Credential struct {
	CredentialSerialNumber string `json:"credential_serial_number"`
	CredentialType         string `json:"credential_type"`
}

// CredentialType is an enum-like type for credential types.
/*************
type CredentialType string
const (
	CredentialTypeRFID        CredentialType = "RFID"
	CredentialTypeTapToCharge CredentialType = "TAP_TO_CHARGE"
	CredentialTypeOther       CredentialType = "OTHER"
)
 **********/

// BusinessPayment represents business payment details for a driver.
type BusinessPayment struct {
	PaymentStatus    string `json:"payment_status"`
	SuspensionReason string `json:"suspension_reason"`
}

// PaymentStatus is an enum-like type for business payment status.
/***************
type PaymentStatus string
const (
	PaymentStatusSuspended PaymentStatus = "SUSPENDED"
)
 **************/
