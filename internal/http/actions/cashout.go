package actions

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestTransfer(transfer objects.Transfer, c *gin.Context) {
	transferBytes, _ := json.Marshal(transfer)
	req, err := utils.GetDefaultPostRequest("POST", transferBytes, "transfers")
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	res, err := utils.DoDefaultPostRequest(req)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	utils.ResolveTransferResponse(res, c)
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
