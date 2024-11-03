package objects

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type CustomerResponse struct {
	CustomerID string `json:"id"`
	Name       string `json:"name"`
	CpfCnpj    string `json:"cpfCnpj"`
}

func ByteToCustomerResponse(b []byte) (CustomerResponse, error) {
	var customerResponse CustomerResponse
	err := json.Unmarshal(b, &customerResponse)
	return customerResponse, err
}

func CustomerResponseToH(customerResponse CustomerResponse) gin.H {
	return gin.H{
		"id":      customerResponse.CustomerID,
		"name":    customerResponse.Name,
		"CpfCnpj": customerResponse.CpfCnpj,
	}
}
