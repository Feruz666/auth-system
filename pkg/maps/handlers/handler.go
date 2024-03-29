package maps

import (
	"log"

	"github.com/Feruz666/auth-system/util"
	"github.com/gin-gonic/gin"
)

var url = GetMapsConfig()

func CreateWorkspace(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/0.0/workspaces/workspace", ctx)
}

func DeleteWorkspace(ctx *gin.Context) {
	util.DeleteGateWayUrl(url+"/0.0/workspaces/workspace", ctx)
}

func DeleteLayers(ctx *gin.Context) {
	util.DeleteGateWayUrl(url+"/0.0/layers/layer", ctx)
}

func CreateDatastore(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/0.0/datastores/datastore", ctx)
}

func DeleteDatastore(ctx *gin.Context) {
	util.DeleteGateWayUrl(url+"/0.0/datastores/datastore", ctx)
}

func DeleteFeatureType(ctx *gin.Context) {
	util.DeleteGateWayUrl(url+"/0.0/featuretypes/featuretype", ctx)
}

func GetMapsConfig() string {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	return config.MAPS_SYSTEM_ADDRESS
}

func CreateCoverageStore(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/0.0/coverage/coveragestore", ctx)
}

func PublishCoverageLayer(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/0.0/coverage/publish", ctx)
}

func MirrorGET(ctx *gin.Context) {
	requestURL := ctx.Request.URL.String()
	util.GetGateWayUrl(url+"/0.0"+requestURL[5:], ctx)
}

func MirrorWMS(ctx *gin.Context) {
	requestURL := ctx.Request.URL.String()
	util.GetWMSGateWayURL(url+"/0.0"+requestURL[5:], ctx)
}
