package objects

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ListCustomerResponse struct {
	Customers []CustomerResponse `json:"data"`
}

func ByteToListCustomerResponse(b []byte) (ListCustomerResponse, error) {
	var listCustomerResponse ListCustomerResponse
	err := json.Unmarshal(b, &listCustomerResponse)
	return listCustomerResponse, err
}

func ListCustomerResponseToH(listCustomerResponse ListCustomerResponse) gin.H {
	return gin.H{
		"customers": listCustomerResponse.Customers,
	}
}
