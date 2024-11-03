package actions

import (
	"mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

const customerEndpoint = "customers"

func ResolveCustomerResponse(res *http.Response, c *gin.Context) {
	body, errResponse := utils.ResolveResponse(res, c)
	customer, err := objects.ByteToCustomerResponse(body)
	if err != nil {
		utils.DefaultError(c, err)
		return
	}
	if errResponse == nil {
		utils.DefaultResponse(c, objects.CustomerResponseToH(customer))
	}
}

func RequestCustomer(customer objects.Customer, c *gin.Context) {
	res := utils.ResolveRequest(customer, c, customerEndpoint)
	ResolveCustomerResponse(res, c)
}

func ParseCustomer(c *gin.Context) *objects.Customer {
	var customer objects.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	return &customer
}

func CustomerAction(c *gin.Context) {
	customer := ParseCustomer(c)
	if customer != nil {
		RequestCustomer(*customer, c)
	}
}
