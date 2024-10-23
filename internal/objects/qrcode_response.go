package objects

import (
	"mguimara/pixchallenge/internal/objects/utils"

	"github.com/gin-gonic/gin"
)

type QrCodeResponse struct {
	ID                    string          `json:"id"`
	EncodedImage          string          `json:"encodedImage"`
	Payload               string          `json:"payload"`
	AllowMultiplePayments bool            `json:"allowMultiplePayments"`
	ExpirationDate        utils.AsaasTime `json:"expirationDate"`
	ExternalReference     string          `json:"externalReference"`
}

func QRToH(qr QrCodeResponse) gin.H {
	return gin.H{
		"id":                     qr.ID,
		"encodedImage":           qr.EncodedImage,
		"payload":                qr.Payload,
		"allowsMultiplePayments": qr.AllowMultiplePayments,
		"expirationDate":         qr.ExpirationDate,
		"externalReference":      qr.ExternalReference,
	}
}
