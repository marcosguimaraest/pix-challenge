package utils

import (
	"errors"
	"io"
	"mguimara/pixchallenge/internal/objects"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const uriApi = "https://sandbox.asaas.com/api/v3/"

var apiKey string = os.Getenv("ASAASKEY")

func setDefaultHeaders(req *http.Request) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("User-Agent", "pix-challenge/0.0.0-alpha")
	req.Header.Add("access_token", apiKey)
}

func GetDefaultPostRequest(method string, body []byte, endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(method, uriApi+endpoint, strings.NewReader(string(body)))
	setDefaultHeaders(req)
	return req, err
}

func DoDefaultPostRequest(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultClient.Do(req)
	return res, err
}

func ReadResponseBody(res *http.Response, c *gin.Context) ([]byte, error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	return body, err
}

func ResolveResponse(res *http.Response, c *gin.Context) {
	body, _ := ReadResponseBody(res, c)
	if res.StatusCode == 401 {
		CustomError(c, errors.New("algo deu errado"), res.StatusCode)
		return
	}
	if res.StatusCode == 400 {
		errResponse, _ := objects.ByteToErrorResponseArray(body)
		DefaultError(c, errors.New(errResponse.Errors[0].Description))
		return
	}
	qr, err := objects.ByteToQrCode(body)
	if err != nil {
		DefaultError(c, err)
		return
	}
	DefaultResponse(c, objects.QRToH(qr))
}

func DefaultResponse(c *gin.Context, h gin.H) {
	c.JSON(http.StatusOK, h)
}

func DefaultError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func CustomError(c *gin.Context, err error, sc int) {
	c.JSON(sc, gin.H{
		"error": err.Error(),
	})
}
