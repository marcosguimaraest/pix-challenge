package objects

type Transfer struct {
	Value             string      `json:"value"`
	BankAccount       BankAccount `json:"bankAccount"`
	OperationType     string      `json:"operationType"`
	PixAddressKey     string      `json:"pixAddressKey"`
	PixAddressKeyType string      `json:"pixAddressKeyType"`
	Description       string      `json:"description"`
	ScheduleDate      string      `json:"scheduleDate"`
	ExternalReference string      `json:"externalReference"`
}
