package utils

import (
	"encoding/json"
	"errors"
	"fmt"
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

func GetDefaultRequest(method string, body []byte, endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(method, uriApi+endpoint, strings.NewReader(string(body)))
	setDefaultHeaders(req)
	return req, err
}

func DoDefaultRequest(req *http.Request) (*http.Response, error) {
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

func ResolveResponse(res *http.Response, c *gin.Context) ([]byte, error) {
	body, err := ReadResponseBody(res, c)
	if res.StatusCode == 401 {
		err := errors.New("algo deu errado")
		CustomError(c, err, res.StatusCode)
		return body, err
	}
	if res.StatusCode == 400 {
		errResponse, _ := objects.ByteToErrorResponseArray(body)
		ResponseError(c, errResponse.Errors[0], res.StatusCode)
		return body, errors.New("algo deu errado")
	}
	if err != nil {
		DefaultError(c, err)
	}
	return body, err
}
func ResolveGetRequest(c *gin.Context, endpoint string) *http.Response {
	req, err := GetDefaultRequest("GET", []byte(""), endpoint)
	if err != nil {
		DefaultError(c, err)
		return nil
	}
	res, err := DoDefaultRequest(req)
	if err != nil {
		DefaultError(c, err)
		return nil
	}
	return (res)
}
func ResolveRequest(obj any, c *gin.Context, endpoint string) *http.Response {
	qrCodeBytes, _ := json.Marshal(obj)
	fmt.Println(string(qrCodeBytes))
	req, err := GetDefaultRequest("POST", qrCodeBytes, endpoint)
	if err != nil {
		DefaultError(c, err)
		return nil
	}
	res, err := DoDefaultRequest(req)
	if err != nil {
		DefaultError(c, err)
		return res
	}
	return res
}

func DefaultResponse(c *gin.Context, h gin.H) {
	c.JSON(http.StatusOK, h)
}

func DefaultError(c *gin.Context, err error) {
	fmt.Println(string(err.Error()))
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func CustomError(c *gin.Context, err error, sc int) {
	fmt.Println(string(err.Error()))
	c.JSON(sc, gin.H{
		"error": err.Error(),
	})
}

func ResponseError(c *gin.Context, errResponse objects.ErrorResponse, sc int) {
	c.JSON(sc, gin.H{
		"error": gin.H{
			"code":        errResponse.Code,
			"description": errResponse.Description,
		},
	})
}
