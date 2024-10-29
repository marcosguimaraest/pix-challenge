package objects

import "mguimara/pixchallenge/internal/objects/utils"

type WebhookTransfer struct {
	ID          string           `json:"id"`
	Event       string           `json:"event"`
	DateCreated utils.AsaasTime  `json:"dateCreated"`
	Transfer    TransferResponse `json:"transfer"`
}
