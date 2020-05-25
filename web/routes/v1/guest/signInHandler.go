package guest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

// Guest godoc
// @Summary guest login api
// @Description APi use for login as guest user
// @Tags guest
// @Accept  json
// @Produce  json
// @Success 200 {object} Guest
// @Failure 400 {object} HTTPError
// @Failure 401 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /guest/signIn [post]
func SignInHanlder() gin.HandlerFunc {

	return func(c *gin.Context) {
		// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		// //insert data into database
		// result, err := mongodb.GetDb().Collection("dummy").InsertOne(ctx, bson.D{
		// 	{"title", "dummmfsddfsd"},
		// 	{"tag", "ccfsdfsdfdf"},
		// })

		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Println(result)

		c.JSON(http.StatusOK, map[string]string{
			"err":     "",
			"message": "guest get api version 1",
		})
	}

}
