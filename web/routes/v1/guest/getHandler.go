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

// Gethandler godoc
// @Summary Show all guest user
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
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

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
