package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
)

func PostGateWayUrl(url string, ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponce(err))
	}

	responseBody := bytes.NewBuffer(data)

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponce(err))
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponce(err))
	}

	jsonData := make(map[string]interface{})
	json.Unmarshal([]byte(string(b)), &jsonData)

	ctx.JSON(resp.StatusCode, jsonData)
}

func GetGateWayUrl(url string, ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponce(err))
	}
	responseBody := bytes.NewBuffer(data)
	if responseBody == nil {
		responseBody = nil
	}
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, responseBody,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/html") // добавляем заголовок Accept

	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponce(err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponce(err))
	}
	jsonData := make(map[string]interface{})
	json.Unmarshal([]byte(body), &jsonData)
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	ctx.JSON(resp.StatusCode, jsonData)
}

func DeleteGateWayUrl(url string, ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponce(err))
	}
	responseBody := bytes.NewBuffer(data)
	if responseBody == nil {
		responseBody = nil
	}
	client := &http.Client{}
	req, err := http.NewRequest(
		"DELETE", url, responseBody,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/html") // добавляем заголовок Accept

	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponce(err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponce(err))
	}
	jsonData := make(map[string]interface{})
	json.Unmarshal([]byte(body), &jsonData)
	ctx.JSON(resp.StatusCode, jsonData)
}

// ErrorResponce ...
func ErrorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}
