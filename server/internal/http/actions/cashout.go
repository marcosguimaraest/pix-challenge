package actions

import (
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

const cashoutEndpoint = "transfers"

func ResolveTransferResponse(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	transfer, err := objects.ByteToTransferResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	if errResponse != nil {
		utils.DefaultResponse(c, objects.TransferToH(transfer))
	}
}

func RequestTransfer(transfer objects.Transfer, c *gin.Context) {
	res := utils.ResolveRequest(transfer, c, cashoutEndpoint)
	ResolveTransferResponse(res, c)
}

func ParseTransfer(c *gin.Context) *objects.Transfer {
	var transfer objects.Transfer

	if err := c.BindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	return &transfer
}

func CashoutAction(c *gin.Context) {
	transfer := ParseTransfer(c)
	if transfer != nil {
		RequestTransfer(*transfer, c)
	}
}
