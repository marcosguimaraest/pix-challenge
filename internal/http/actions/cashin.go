package actions

import (
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResolveQrCodeResponse(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	qr, err := objects.ByteToQrCodeResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	if errResponse == nil {
		utils.DefaultResponse(c, objects.QRToH(qr))
	}
}

func RequestQrCode(qrCode objects.QrCode, c *gin.Context) {
	res := utils.ResolveRequest(qrCode, c)
	ResolveQrCodeResponse(res, c)
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
