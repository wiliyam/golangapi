package guest

import (
	"context"
	"fmt"
	"golangapi/library/mongodb"
	dbmodel "golangapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

//ReqBodyValidate request body tyep
type ReqBodyValidate struct {
	// Name     string `json:"name" example:"wiliyam"`
	DeviceID   string `json:"deviceId" example:"1234567809"`
	DeviceType string `json:"deviceType" example:"0"`
}

//SignInHanlder Guest godoc
// @Summary guest login api
// @Description APi use for login as guest user
// @Tags guest
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} Guest
// @Failure 400 {object} cm.HTTPError400
// @Failure 401 {object} cm.HTTPError401
// @Failure 404 {object} cm.HTTPError404
// @Failure 500 {object} cm.HTTPError500
// @Router /guest/signIn [post]
func SignInHanlder() gin.HandlerFunc {

	return func(c *gin.Context) {

		//get params
		// id := c.Param("id")
		//get headers
		//authHeader := c.GetHeader("Authorization")
		//get body data

		var body dbmodel.GusetUserModelPost
		c.BindJSON(&body)
		deviceid := body.DeviceId
		devicetype := body.DeviceType
		ipAddress := body.IpAddress

		var devicetypeMsg string

		switch devicetype {
		default:
			devicetypeMsg = "unknown"
		case 0:
			devicetypeMsg = "unknown"
		case 1:
			devicetypeMsg = "web"
		case 2:
			devicetypeMsg = "android"
		case 3:
			devicetypeMsg = "ios"
		}

		fmt.Println("body", body)
		fmt.Println("deviceid", deviceid)
		fmt.Println("devicetype", devicetype)
		fmt.Println("devicetypeMsg", devicetypeMsg)

		insertdata := dbmodel.GusetUserModelPost{
			DeviceId:      deviceid,
			DeviceType:    devicetype,
			DeviceTypeMsg: devicetypeMsg,
			IpAddress:     ipAddress,
			LoginAt:       time.Now(),
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		result, err := mongodb.GetDb().Collection("guest").InsertOne(ctx, insertdata)

		if err != nil {
			println("err--->", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong",
			})
			return
		}
		fmt.Println("result--->>", result)
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    "",
		})

	}

}
