package web

import (
	"golangapi/dotenv"
	"golangapi/web/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//init server method
func InitServer() {

	apibaseurlversion1 := dotenv.GetVariableValue("API_BASE_URL_VERSION_1")
	apibaseurlversion2 := dotenv.GetVariableValue("API_BASE_URL_VERSION_2")
	fmt.Printf("godotenv : %s = %s \n", "baseurl version 1", apibaseurlversion1)
	fmt.Printf("godotenv : %s = %s \n", "baseurl version 2", apibaseurlversion2)

	//set to production mode
	if os.Getenv("PRODUCTION_MODE") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	//initialize ginn router
	r := gin.Default()

	//register html templates
	r.LoadHTMLGlob("templates/*")

	//register version 1 routes
	v1 := r.Group(apibaseurlversion1)
	routes.Routesversion1(v1)

	//register version 2 routes
	v2 := r.Group(apibaseurlversion2)
	routes.Routesversion2(v2)

	//render html on route
	r.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})
	// r.GET("/ping", routes.PingGet())
	r.Run()

}
