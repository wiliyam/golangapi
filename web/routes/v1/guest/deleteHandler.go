package guest

import (
	"context"
	"fmt"
	"golangapi/library/mongodb"
	dbmodel "golangapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gin-gonic/gin"
)

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
func DeleteHanlder() gin.HandlerFunc {

	return func(c *gin.Context) {

		//get params
		// id := c.Param("id")
		//get headers
		//authHeader := c.GetHeader("Authorization")
		//get body data

		var body dbmodel.GusetUserModelGet
		c.BindJSON(&body)
		q := c.Request.URL.Query()
		id := q.Get("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id is in not object id formate in query params",
			})
			return
		}
		fmt.Println(objectId)

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		result, err := mongodb.GetDb().Collection("guest").DeleteOne(ctx, bson.M{"_id": objectId})

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
