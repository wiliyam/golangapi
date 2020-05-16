package v1

import (
	"golangapi/web/routes/v1/guest"
	"golangapi/web/routes/v1/user"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	guest.Routes(router)
	user.Routes(router)
}
