package actions

import (
	"fmt"
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

const listPaymentsEndpoint = "payments"

func GetListPaymentsEndpoint(lp objects.ListPayments) string {
	return listPaymentsEndpoint + "?status=" + lp.Status + "&limit=100"
}

func RequestListPayments(lp objects.ListPayments, c *gin.Context) {
	endpoint := GetListPaymentsEndpoint(lp)
	println(endpoint)
	res := utils.ResolveGetRequest(c, endpoint)
	ResolveListPayments(res, c)
}

func ResolveListPayments(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	listPaymentResponse, err := objects.ByteToListPaymentResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	fmt.Println(listPaymentResponse)
	if errResponse == nil {
		utils.DefaultResponse(c, objects.ListPaymentResponseToH(listPaymentResponse))
	}
}

func ParseListPayments(c *gin.Context) *objects.ListPayments {
	var lp objects.ListPayments
	lp.Status = c.Query("status")
	return &lp
}

func ListPaymentsAction(c *gin.Context) {
	payment := ParseListPayments(c)
	if payment != nil {
		RequestListPayments(*payment, c)
	}
}
