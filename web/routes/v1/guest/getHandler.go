package guest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Guest struct {
	Err     string `json:"err" example:""`
	Message string `json:"message" example:"some message"`
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// Gethandler godoc
// @Summary Show an guest user
// @Description get all guest users
// @Tags Guest
// @Accept  json
// @Produce  json
// @Success 200 {object} Guest
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /guest [get]

func Gethandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"err":     "",
			"message": "guest get api version 1",
		})
	}

}
