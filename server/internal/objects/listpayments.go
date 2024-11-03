package objects

type ListPayments struct {
	Limit  int    `json:"limit"`
	Status string `json:"status"`
}
