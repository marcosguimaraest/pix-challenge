package http

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.POST("/pix/qrcode/generate", GenerateQrCode)
}
