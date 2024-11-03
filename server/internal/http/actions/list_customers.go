package actions

import (
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

const listCustomersEndpoint = "customers"

func GetListCustomersEndpoint() string {
	return listCustomersEndpoint + "?limit=100"
}

func RequestListCustomers(c *gin.Context) {
	endpoint := GetListCustomersEndpoint()
	res := utils.ResolveGetRequest(c, endpoint)
	ResolveListCustomer(res, c)
}

func ResolveListCustomer(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	listCustomerResponse, err := objects.ByteToListCustomerResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	if errResponse == nil {
		utils.DefaultResponse(c, objects.ListCustomerResponseToH(listCustomerResponse))
	}
}

func ListCustomerAction(c *gin.Context) {
	RequestListCustomers(c)
}
