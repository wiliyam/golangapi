package guest

import (
	cm "golangapi/web/commonmodules"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

//ReqBodyValidate request body tyep
type ReqBodyValidate struct {
	Name     string `json:"name" example:"wiliyam"`
	DeviceID string `json:"deviceId" example:"1234567809"`
}

//SignInHanlder Guest godoc
// @Summary guest login api
// @Description APi use for login as guest user
// @Tags guest
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} Guest
// @Failure 400 {object} cm.HTTPError401
// @Failure 401 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /guest/signIn [post]
func SignInHanlder() gin.HandlerFunc {

	return func(c *gin.Context) {

		//get params
		// id := c.Param("id")
		//get headers
		//authHeader := c.GetHeader("Authorization")

		res := cm.HTTPError401{

			Status:  401,
			Message: "Unauthorized",
		}
		c.JSON(http.StatusUnauthorized, res)

		// c.JSON(http.StatusOK, map[string]string{
		// 	"err":     "",
		// 	"message": "guest get api version 1",
		// })
	}

}
