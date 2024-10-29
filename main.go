package main

import (
	"fmt"
	"mguimara/pixchallenge/internal/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default(http.SetRoutes)

	fmt.Println(os.Getenv("ASAASKEY"))
	router.Run("localhost:8080")
}
