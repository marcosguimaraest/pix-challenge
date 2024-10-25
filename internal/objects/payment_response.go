package objects

import "mguimara/pixchallenge/internal/objects/utils"

type PaymentResponse struct {
	ID          string             `json:"id"`
	CustomerID  string             `json:"customer"`
	BillingType string             `json:"billingType"`
	Value       float64            `json:"value"`
	DueDate     utils.BirthdayTime `json:"dueDate"`
	Status      string             `json:"status"`
}
