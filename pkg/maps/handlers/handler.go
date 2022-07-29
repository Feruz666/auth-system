package maps

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/Feruz666/auth-system/util"
	"github.com/gin-gonic/gin"
)

var url = GetMapsConfig()

// Workspace handlers
func GetWorkspaces(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/workspaces", ctx)
}

func CreateWorkspace(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/0.0/workspaces/workspace", ctx)
}

func GetWorkspace(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/workspaces/workspace", ctx)
}

func DeleteWorkspace(ctx *gin.Context) {
	util.DeleteGateWayUrl(url+"/0.0/workspaces/workspace", ctx)
}

// Styles handler
func GetStyles(ctx *gin.Context) {
	workspace := ctx.Query("workspace")
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("{\"workspace\":\"" + workspace + "\"}")))
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Request.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "POST, DELETE, GET, PUT")
	// body, _ := ctx.Request.GetBody()
	// ctx.Request.Header.Set("Content-Length"))
	util.GetGateWayUrl(url+"/0.0/styles", ctx)
}

// Layers handlers
func GetLayers(ctx *gin.Context) {
	workspace := ctx.Query("workspace")
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("{\"workspace\":\"" + workspace + "\"}")))
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Request.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "POST, DELETE, GET, PUT")
	// body, _ := ctx.Request.GetBody()
	// ctx.Request.Header.Set("Content-Length"))
	util.GetGateWayUrl(url+"/0.0/layers", ctx)
}

func GetLayer(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/layers/layer", ctx)
}

func DeleteLayers(ctx *gin.Context) {
	util.DeleteGateWayUrl(url+"/0.0/layers/layer", ctx)
}

// Datastore handlers
func GetDatastores(ctx *gin.Context) {
	workspace := ctx.Query("workspace")
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("{\"workspace\":\"" + workspace + "\"}")))
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Request.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "POST, DELETE, GET, PUT")
	util.GetGateWayUrl(url+"/0.0/datastores", ctx)
}

func GetDatastore(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/datastores/datastore", ctx)
}

func DatastoreExists(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/datastores/datastore/exists", ctx)
}

func CreateDatastore(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/0.0/datastores/datastore", ctx)
}

func DeleteDatastore(ctx *gin.Context) {
	util.DeleteGateWayUrl(url+"/0.0/datastores/datastore", ctx)
}

// LayerGroups handlers
func GetLayerGroups(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/layergroups", ctx)
}

func GetLayerGroup(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/layergroups/layergroup", ctx)
}

// FeatureType handlers
func GetFeatureTypes(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/featuretypes", ctx)
}

func GetFeatureType(ctx *gin.Context) {
	util.GetGateWayUrl(url+"/0.0/featuretypes/featuretype", ctx)
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

func GetWMS(ctx *gin.Context) {
	util.GetWMSGateWayURL(url+"/0.0/wms", ctx)
}


func GetUsersLayers(ctx gin.Context) {
	
}