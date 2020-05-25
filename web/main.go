package web

import (
	"fmt"
	"golangapi/dotenv"
	"golangapi/web/routes"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "golangapi/web/docs"
)

// @title glang boiler plate api
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email wiliyambhadani@gmail.com

// @host localhost:9191
// @BasePath /api/v1

// @schemes http https

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

	//register swagger api documentation
	if os.Getenv("PRODUCTION_MODE") == "true" {
		color.Cyan("Disabling api documentation...")
		r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DisablingWrapHandler"))
	} else {
		color.Cyan("Enabling api documentation...")
		swaggerUrl := "http://localhost:" + os.Getenv("PORT") + "/doc/doc.json"
		url := ginSwagger.URL(swaggerUrl) // The url pointing to API definition
		r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	r.Run()

}
