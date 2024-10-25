package actions

import (
	"encoding/json"
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResolvePaymentResponse(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	payment, err := objects.ByteToPaymentResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	if errResponse == nil {
		utils.DefaultResponse(c, objects.PaymentToH(payment))
	}
}

func RequestPayment(payment objects.Payment, c *gin.Context) {
	paymentBytes, _ := json.Marshal(payment)
	req, err := utils.GetDefaultPostRequest("POST", paymentBytes, "payments")
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	res, err := utils.DoDefaultPostRequest(req)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	ResolvePaymentResponse(res, c)
}

func ParsePayment(c *gin.Context) *objects.Payment {
	var payment objects.Payment

	if err := c.BindJSON(&payment); err != nil {
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
