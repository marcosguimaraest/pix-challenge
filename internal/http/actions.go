package http

import (
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAsaasQrCode() {

}
func PostGenerateQrCode(c *gin.Context) {
	var qrCode objects.QrCode

	if err := c.BindJSON(&qrCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
}
