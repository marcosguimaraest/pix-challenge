package objects

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/objects/utils"

	"github.com/gin-gonic/gin"
)

type TransferResponse struct {
	Object                string             `json:"object"`
	Id                    string             `json:"id"`
	Type                  string             `json:"type"`
	DateCreated           utils.BirthdayTime `json:"dateCreated"`
	Value                 float64            `json:"value"`
	NetValue              float64            `json:"netValue"`
	Status                string             `json:"status"`
	TransferFee           float64            `json:"transferFee"`
	EffectiveDate         utils.BirthdayTime `json:"effectiveDate"`
	ScheduleDate          utils.BirthdayTime `json:"scheduleDate"`
	EndToEndIdentifier    string             `json:"endToEndIdentifier"`
	Authorized            bool               `json:"authorized"`
	FailReason            string             `json:"failReason"`
	ExternalReference     string             `json:"externalReference"`
	TransactionReceiptUrl string             `json:"transactionReceiptUrl"`
	OperationType         string             `json:"operationType"`
	Description           string             `json:"description"`
}

func ByteToTransferResponse(b []byte) (TransferResponse, error) {
	var transferResponse TransferResponse
	err := json.Unmarshal(b, &transferResponse)
	return transferResponse, err
}

func TransferToH(tr TransferResponse) gin.H {
	return gin.H{
		"object":                tr.Object,
		"id":                    tr.Id,
		"type":                  tr.Type,
		"dateCreated":           tr.DateCreated,
		"value":                 tr.Value,
		"netValue":              tr.NetValue,
		"status":                tr.Status,
		"transferFee":           tr.TransferFee,
		"effectiveDate":         tr.EffectiveDate,
		"scheduleDate":          tr.ScheduleDate,
		"endToEndIdentifier":    tr.EndToEndIdentifier,
		"authorized":            tr.Authorized,
		"failReason":            tr.FailReason,
		"externalReference":     tr.ExternalReference,
		"transactionReceiptUrl": tr.TransactionReceiptUrl,
		"operationType":         tr.OperationType,
		"description":           tr.Description,
	}
}
