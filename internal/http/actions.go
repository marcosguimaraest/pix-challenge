package http

import (
	"encoding/json"
	"fmt"
	"mguimara/pixchallenge/internal/objects"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var uriApi string = "https://sandbox.asaas.com/api/v3/pix/qrCodes/static"
var apiKey string = os.Getenv("KEYASAAS")

func PostQrCode(qrCode objects.QrCode) {

	qrCodeBytes, _ := json.Marshal(qrCode)
	fmt.Println(string(qrCodeBytes))
	req, _ := http.NewRequest("POST", uriApi, strings.NewReader(string(qrCodeBytes)))
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("User-Agent", "mguimara/pix-challenge")
	req.Header.Add("access_token", apiKey)

	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.Body)
}
func GenerateQrCode(c *gin.Context) {
	var qrCode objects.QrCode

	if err := c.BindJSON(&qrCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	PostQrCode(qrCode)
}
