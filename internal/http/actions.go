package http

import (
	"encoding/json"
	"fmt"
	"io"
	"mguimara/pixchallenge/internal/objects"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var uriApi string = "https://sandbox.asaas.com/api/v3/pix/qrCodes/static"
var apiKey string = os.Getenv("ASAASKEY")

func PostQrCode(qrCode objects.QrCode, c *gin.Context) {
	qrCodeBytes, _ := json.Marshal(qrCode)
	fmt.Println(string(qrCodeBytes))
	req, err := http.NewRequest("POST", uriApi, strings.NewReader(string(qrCodeBytes)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("User-Agent", "pix-challenge/0.0.0-alpha")
	req.Header.Add("access_token", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(string(body))
}

func GenerateQrCode(c *gin.Context) {
	var qrCode objects.QrCode

	if err := c.BindJSON(&qrCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	PostQrCode(qrCode, c)
}
