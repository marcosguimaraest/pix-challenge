package actions

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestQrCode(qrCode objects.QrCode, c *gin.Context) {
	qrCodeBytes, _ := json.Marshal(qrCode)
	req, err := utils.GetDefaultPostRequest("POST", qrCodeBytes, "pix/qrCodes/static")
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	res, err := utils.DoDefaultPostRequest(req)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	utils.ResolveResponse(res, c)
}

func ParseQrCode(c *gin.Context) *objects.QrCode {
	var qrCode objects.QrCode

	if err := c.BindJSON(&qrCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	return &qrCode
}

func QrCodeAction(c *gin.Context) {
	qrCode := ParseQrCode(c)
	if qrCode != nil {
		RequestQrCode(*qrCode, c)
	}
}
