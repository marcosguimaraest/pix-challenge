package objects

import (
	"time"
)

type QrCode struct {
	AdressKey              string    `json:"adressKey"`
	Description            string    `json:"description"`
	Value                  float64   `json:"value"`
	Format                 string    `json:"format"`
	ExpirationDate         time.Time `json:"expirationDate"`
	ExpirationSeconds      int32     `json:"expirationSeconds"`
	AllowsMultiplePayments bool      `json:"allowMultiplePayments"`
	ExternalReference      string    `json:"externalReference"`
}
