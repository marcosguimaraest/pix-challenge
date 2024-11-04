package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	hutils "mguimara/pixchallenge/internal/http/utils"
	"mguimara/pixchallenge/internal/objects"
	"mguimara/pixchallenge/internal/objects/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func HundredReqs(c *gin.Context) {
	i := 0
	var paymentObjects [100]objects.Payment
	for i < 100 {
		paymentObjects[i].BillingType = "PIX"
		paymentObjects[i].CustomerID = "cus_000006317497"
		paymentObjects[i].Value = 5.01 + rand.Float64()*(1000-5.01)
		paymentObjects[i].DueDate = utils.BirthdayTime{time.Now()}
		body, _ := json.Marshal(paymentObjects[i])
		req, err := hutils.GetDefaultRequest("POST", body, "payments")
		if err != nil {
			hutils.CustomError(c, err, 404)
			return
		}
		res, err := hutils.DoDefaultRequest(req)
		if err != nil {
			hutils.CustomError(c, err, res.StatusCode)
			return
		}
		defer res.Body.Close()
		body, errr := io.ReadAll(res.Body)
		if err != nil {
			hutils.CustomError(c, errr, res.StatusCode)
			return
		}
		fmt.Println(string(body))
		i++
	}
}
