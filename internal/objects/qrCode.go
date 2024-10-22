package objects

import "mguimara/pixchallenge/internal/objects/utils"

type QrCode struct {
	AdressKey              string          `json:"addressKey"`
	Description            string          `json:"description"`
	Value                  float64         `json:"value"`
	Format                 string          `json:"format"`
	ExpirationDate         utils.AsaasTime `json:"expirationDate"`
	ExpirationSeconds      int32           `json:"expirationSeconds"`
	AllowsMultiplePayments bool            `json:"allowMultiplePayments"`
	ExternalReference      string          `json:"externalReference"`
}
