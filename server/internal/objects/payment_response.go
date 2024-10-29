package objects

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/objects/utils"

	"github.com/gin-gonic/gin"
)

type PaymentResponse struct {
	ID          string             `json:"id"`
	CustomerID  string             `json:"customer"`
	BillingType string             `json:"billingType"`
	Value       float64            `json:"value"`
	DueDate     utils.BirthdayTime `json:"dueDate"`
	Status      string             `json:"status"`
}

func ByteToPaymentResponse(b []byte) (PaymentResponse, error) {
	var paymentResponse PaymentResponse
	err := json.Unmarshal(b, &paymentResponse)
	return paymentResponse, err
}

func PaymentToH(payment PaymentResponse) gin.H {
	return gin.H{
		"id":          payment.ID,
		"customer":    payment.CustomerID,
		"billingType": payment.BillingType,
		"value":       payment.Value,
		"dueDate":     payment.DueDate,
		"status":      payment.Status,
	}
}
