package guest

import (
	"context"
	"fmt"
	"golangapi/library/mongodb"
	dbmodel "golangapi/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		//get query params
		q := c.Request.URL.Query()
		//pagination
		// var pagenum int
		// var pagesize int
		pagenum, err := strconv.ParseInt(q.Get("pagenum"), 6, 12)
		pagesize, err := strconv.ParseInt(q.Get("pagesize"), 6, 12)

		skip := pagesize * pagenum

		fmt.Println(pagenum)
		fmt.Println(pagesize)

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		// res := []Todo{}

		guests := []dbmodel.GusetUserModelGet{}

		options := options.Find()

		// Sort by `_id` field descending
		options.SetSort(bson.D{{"_id", -1}})

		// Limit by 10 documents only
		options.SetLimit(pagesize)
		options.SetSkip(skip)

		// DB.C("collection").Find(nil, filter, &opts)

		cursor, err := mongodb.GetDb().Collection("guest").Find(ctx, bson.D{}, options)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}

		// Iterate through the returned cursor.
		for cursor.Next(ctx) {
			var guest dbmodel.GusetUserModelGet
			cursor.Decode(&guest)
			guests = append(guests, guest)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    guests,
		})
	}

}
func GetOnehandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		// res := []Todo{}
		q := c.Request.URL.Query()
		id := q.Get("id")
		fmt.Println(id)
		// id = "5b9223c86486b341ea76910c"
		// convert id string to ObjectId
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id is in not object id formate",
			})
			return
		}
		fmt.Println(objectId)
		// guests := []dbmodel.GusetUserModelGet{}/
		var guest dbmodel.GusetUserModelGet

		err = mongodb.GetDb().Collection("guest").FindOne(ctx, bson.M{"_id": objectId}).Decode(&guest)

		if err != nil {
			fmt.Println("error->", err.Error())
			c.JSON(http.StatusNotFound, gin.H{
				"message": "data not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    guest,
		})

		// Iterate through the returned cursor.
		// for cursor.Next(ctx) {
		// 	cursor.Decode(&guest)
		// 	guests = append(guests, guest)
		// }
	}

}
