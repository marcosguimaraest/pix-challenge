package objects

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/objects/utils"
)

type QrCode struct {
	AdressKey              string          `json:"addressKey"`
	Description            string          `json:"description"`
	Value                  float64         `json:"value"`
	Format                 string          `json:"format"`
	ExpirationDate         utils.AsaasTime `json:"expirationDate,omitempty"`
	ExpirationSeconds      int32           `json:"expirationSeconds,omitempty"`
	AllowsMultiplePayments bool            `json:"allowMultiplePayments"`
	ExternalReference      string          `json:"externalReference"`
}

func ByteToQrCode(b []byte) (QrCodeResponse, error) {
	var qrCodeResponse QrCodeResponse
	err := json.Unmarshal(b, &qrCodeResponse)
	return qrCodeResponse, err
}
