package objects

import "mguimara/pixchallenge/internal/objects/utils"

type WebhookPayment struct {
	ID          string          `json:"id"`
	Event       string          `json:"event"`
	DateCreated utils.AsaasTime `json:"dateCreated"`
	Payment     PaymentResponse `json:"payment"`
}
