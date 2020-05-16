package user

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.GET("/user", Gethandler())
}
