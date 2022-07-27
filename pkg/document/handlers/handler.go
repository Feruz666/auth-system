package document

import (
	"net/http"

	"github.com/Feruz666/auth-system/util"
	"github.com/gin-gonic/gin"
)

func Example(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"resp": "HIIII",
	})
}

func GetExample(ctx *gin.Context) {
	util.GetGateWayUrl("https://dummy.restapiexample.com/api/v1/employees", ctx)
}

func DeleteExample(ctx *gin.Context) {
	util.DeleteGateWayUrl("https://jsonplaceholder.typicode.com/posts", ctx)
}
