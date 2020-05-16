package routes

import (
	v1 "golangapi/web/routes/v1"
	v2 "golangapi/web/routes/v2"

	"github.com/gin-gonic/gin"
)

func Routesversion1(router *gin.RouterGroup) {
	// router.GET("/", PingGet())

	//register all routes here
	v1.Routes(router)
}

func Routesversion2(router *gin.RouterGroup) {
	// router.GET("/", PingGet())

	//register all routes here
	v2.Routes(router)
}
