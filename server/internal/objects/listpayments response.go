package objects

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ListPaymentResponse struct {
	Payments []PaymentResponse `json:"data"`
}

func ByteToListPaymentResponse(b []byte) (ListPaymentResponse, error) {
	var listPaymentResponse ListPaymentResponse
	err := json.Unmarshal(b, &listPaymentResponse)
	return listPaymentResponse, err
}

func ListPaymentResponseToH(lpr ListPaymentResponse) gin.H {
	return gin.H{
		"payments": lpr.Payments,
	}
}
