package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type pix struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
}

var uriApi string = "https://sandbox.asaas.com/api/v3"

func main() {
	router := gin.Default()

	fmt.Println(os.Getenv("PATH"))
	router.Run("localhost:8080")
}
