package http

import (
	"mguimara/pixchallenge/internal/http/actions"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	//g.POST("/pix/qrcode/generate", actions.QrCodeAction)
	g.POST("/payment/", actions.PaymentAction)
	g.GET("/payment", actions.ListPaymentsAction)
	g.POST("/customer/", actions.CustomerAction)
	g.GET("/customer", actions.ListCustomerAction)
	g.POST("/transfer/", actions.CashoutAction)
	g.GET("/qrcode", actions.PaymentQrCodeAction)
	g.GET("/testhundredcashin", actions.HundredReqs)
}
