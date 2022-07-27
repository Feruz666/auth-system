package sensors

import (
	"log"

	"github.com/Feruz666/auth-system/util"
	"github.com/gin-gonic/gin"
)

var url = GetSensorConfig()

func CreateSensor(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/api/sensors", ctx)
}

func GetSensors(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/api/sensors", ctx)
}

func GetSensorsCharts(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/api/sensors/charts", ctx)
}

func GetSensorConfig() string {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	return config.SENSORS_SYSTEM_ADDRESS
}
