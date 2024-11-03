package actions

import (
	"fmt"
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

const qrcodeEndpoint = "payments"

func GetQrCodeEndpoint(id string) string {
	return qrcodeEndpoint + "/" + id + "/" + "pixQrCode"
}

func GetQrCodeBody(id string) []byte {
	return []byte("{\"id\":\"" + id + "\"}")
}

func RequestPaymentQrCode(pr objects.PaymentQrCode, c *gin.Context) {
	res := utils.ResolveGetRequest(c, GetQrCodeEndpoint(pr.PaymentID))
	ResolvePaymentQrCodeResponse(res, c)
}

func ResolvePaymentQrCodeResponse(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	paymentQrCodeResponse, err := objects.ByteToPaymentQrCodeResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	fmt.Println(paymentQrCodeResponse)
	if errResponse == nil {
		utils.DefaultResponse(c, objects.PaymentQRCodeResponseToH(paymentQrCodeResponse))
	}
}

func ParsePaymentQrCode(c *gin.Context) *objects.PaymentQrCode {
	var pqr objects.PaymentQrCode

	if pqr.PaymentID = c.Query("id"); pqr.PaymentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Id inválido ouj inexistente!",
		})
		return nil
	}
	fmt.Println(pqr.PaymentID)
	return &pqr
}

func PaymentQrCodeAction(c *gin.Context) {
	payment := ParsePaymentQrCode(c)
	if payment != nil {
		RequestPaymentQrCode(*payment, c)
	}
}
