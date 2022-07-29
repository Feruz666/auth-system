package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	util "github.com/Feruz666/auth-system/util/maps"
	"github.com/gin-gonic/gin"
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
	ctx.Header("Access-Control-Allow-Origin", "*")
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

// Get WMS
func GetWMSGateWayURL(url string, ctx *gin.Context) {

	urlParams := &util.WMSParams{}

	if err := ctx.ShouldBind(urlParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "URL parametr unmarshal error, please check url parametrs",
		})

		return
	}

	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "request to wms service error",
		})

		return
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

	query := req.URL.Query()

	query.Add("service", urlParams.Service)
	query.Add("version", urlParams.Version)
	query.Add("request", urlParams.Request)
	query.Add("layers", urlParams.Layers)
	query.Add("bbox", urlParams.BBox)
	query.Add("width", urlParams.Width)
	query.Add("height", urlParams.Height)
	query.Add("srs", urlParams.SRS)
	query.Add("styles", urlParams.Styles)
	query.Add("format", urlParams.Format)
	query.Add("transparent", urlParams.Transparent)

	req.URL.RawQuery = query.Encode()

	req.Header.Add("Accept", "text/html") // добавляем заголовок Accept

	resp, err := client.Do(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "request to wms service error",
		})

		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "request to wms service error",
		})

		return
	}

	ctx.Header("Content-Type", "image/png")
	ctx.Header("Content-Length", strconv.Itoa(len(body)))
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(body) 
}

// ErrorResponce ...
func ErrorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}



