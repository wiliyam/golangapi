package v2

import (
	"golangapi/web/routes/v2/guest"
	"golangapi/web/routes/v2/user"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	guest.Routes(router)
	user.Routes(router)
}
