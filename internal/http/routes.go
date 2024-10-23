package http

import (
	"mguimara/pixchallenge/internal/http/actions"
	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.POST("/pix/qrcode/generate", actions.QrCodeAction)
}
