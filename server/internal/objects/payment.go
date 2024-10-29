package objects

import "mguimara/pixchallenge/internal/objects/utils"

type Payment struct {
	CustomerID  string             `json:"customer"`
	BillingType string             `json:"billingType"`
	Value       float64            `json:"value"`
	DueDate     utils.BirthdayTime `json:"dueDate"`
}
