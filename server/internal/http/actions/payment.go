package actions

import (
	"fmt"
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

const paymentEndpoint = "payments"

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
