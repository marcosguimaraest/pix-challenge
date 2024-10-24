package objects

import "mguimara/pixchallenge/internal/objects/utils"

type BankAccount struct {
	Bank            Bank               `json:"bank"`
	AccountName     string             `json:"accountName"`
	OwnerName       string             `json:"ownerName"`
	OwnerBirthDate  utils.BirthdayTime `json:"ownerBirthDate"`
	CpfCnpj         string             `json:"cpfCnpj"`
	Agency          string             `json:"agency"`
	Account         string             `json:"account"`
	AccountDigit    string             `json:"accountDigit"`
	BankAccountType string             `json:"bankAccountType"`
	Ispb            string             `json:"ispb"`
}
