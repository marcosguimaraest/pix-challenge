package actions

import (
	"errors"
	"fmt"
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

const paymentEndpoint = "payments"

func GetQrCodeEndpoint(id string) string {
	return paymentEndpoint + "/" + id + "/" + "pixQrCode"
}

func GetQrCodeBody(id string) []byte {
	return []byte("{\"id\":\"" + id + "\"}")
}

func RequestPaymentQrCode(pr objects.PaymentResponse, c *gin.Context) *[]byte {
	req, err := utils.GetDefaultRequest("GET", []byte(""), GetQrCodeEndpoint(pr.ID))
	if err != nil {
		utils.DefaultError(c, err)
		return nil
	}
	res, err := utils.DoDefaultRequest(req)
	if err != nil {
		utils.DefaultError(c, err)
		return nil
	}
	body, _ := utils.ResolveResponse(res, c)
	return &body
}

func ResolvePaymentResponse(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	payment, err := objects.ByteToPaymentResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	if errResponse == nil {
		bodyQr := RequestPaymentQrCode(payment, c)
		if bodyQr != nil {
			pqr, err := objects.ByteToPaymentQrCodeResponse(*bodyQr)
			if err != nil {
				utils.DefaultError(c, err)
				return
			}
			utils.DefaultResponse(c, objects.PaymentQRToH(pqr))
			return
		} else {
			utils.DefaultError(c, errors.New("algo deu errado"))
			return
		}
	}
}

func RequestPayment(payment objects.Payment, c *gin.Context) {
	res := utils.ResolveRequest(payment, c, paymentEndpoint)
	ResolvePaymentResponse(res, c)
}

func ParsePayment(c *gin.Context) *objects.Payment {
	var payment objects.Payment

	if err := c.BindJSON(&payment); err != nil {
		fmt.Println(string(err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	return &payment
}

func PaymentAction(c *gin.Context) {
	payment := ParsePayment(c)
	if payment != nil {
		RequestPayment(*payment, c)
	}
}
