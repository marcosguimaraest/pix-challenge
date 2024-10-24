package objects

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/objects/utils"
)

type Transfer struct {
	Value             float64            `json:"value"`
	BankAccount       BankAccount        `json:"bankAccount"`
	OperationType     string             `json:"operationType"`
	PixAddressKey     string             `json:"pixAddressKey"`
	PixAddressKeyType string             `json:"pixAddressKeyType"`
	Description       string             `json:"description"`
	ScheduleDate      utils.BirthdayTime `json:"scheduleDate"`
	ExternalReference string             `json:"externalReference"`
}

func ByteToTransfer(b []byte) (Transfer, error) {
	var transfer Transfer
	err := json.Unmarshal(b, &transfer)
	return transfer, err
}
