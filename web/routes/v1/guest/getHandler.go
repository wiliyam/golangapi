package guest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Gethandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"err":     "",
			"message": "guest get api version 1",
		})
	}

}
