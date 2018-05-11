package service

import (
	"github.com/artify/constant"
	"github.com/artify/model"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start() {

	r := gin.Default()
	r.Run()
}

func initializeEngine() *gin.Engine {
	if constant.AppMode == "debug" {
		gin.SetMode(gin.DebugMode)
		return gin.Default()
	}

	// Release
	gin.SetMode(gin.ReleaseMode)
	return gin.New()
}
func GetEngine(config constant.Config) *gin.Engine {

	app := initializeEngine()
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	// Apply Routes

	// Default route
	app.GET("/", home)

	// Return
	return app
}

func home(c *gin.Context) {
	res := model.Response{
		Code:    http.StatusOK,
		Data:    map[string]string{"Hello": "It's Artify Core"},
		Message: "Success",
	}
	c.JSON(http.StatusOK, res)
}
