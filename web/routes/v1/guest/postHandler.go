package guest

import (
	"context"
	"fmt"
	"golangapi/library/mongodb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
// @Summary login as guest user
// @Description this api is use for login as guest in app
// @Tags Guest
// @Accept  json
// @Produce  json
// @Success 200 {object} Guest
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /guest [post]

func PostHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		//insert data into database
		result, err := mongodb.GetDb().Collection("dummy").InsertOne(ctx, bson.D{
			{"title", "dummmfsddfsd"},
			{"tag", "ccfsdfsdfdf"},
		})

		if err != nil {
			panic(err)
		}

		fmt.Println(result)
		c.JSON(http.StatusOK, map[string]string{
			"err":     "",
			"message": "guest get api version 1",
		})
	}

}
