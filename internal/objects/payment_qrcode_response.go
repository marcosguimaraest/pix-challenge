package objects

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/objects/utils"

	"github.com/gin-gonic/gin"
)

type PaymentQrCodeResponse struct {
	EncodedImage   string          `json:"encodedImage"`
	Payload        string          `json:"payload"`
	ExpirationDate utils.AsaasTime `json:"expirationDate"`
}

func ByteToPaymentQrCodeResponse(b []byte) (PaymentQrCodeResponse, error) {
	var paymentQrCodeResponse PaymentQrCodeResponse
	err := json.Unmarshal(b, &paymentQrCodeResponse)
	return paymentQrCodeResponse, err
}

func PaymentQRToH(pqr PaymentQrCodeResponse) gin.H {
	return gin.H{
		"encodedImage":   pqr.EncodedImage,
		"payload":        pqr.Payload,
		"expirationDate": pqr.ExpirationDate,
	}
}
