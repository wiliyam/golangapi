package guest

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.GET("/guest", Gethandler())
	router.POST("/guest/signIn", SignInHanlder())
}
