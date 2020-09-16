package guest

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.GET("/guest/list", Gethandler())
	router.GET("/guest", GetOnehandler())
	router.PATCH("/guest", EditHanlder())
	router.DELETE("/guest", DeleteHanlder())
	router.POST("/guest/signIn", SignInHanlder())
}
