package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogeriofontes/api-go-gin/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

// NewGinRouter all the routes are defined here

// NewGinRouter godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /healthcheck [get]
func NewGinRouter() GinRouter {

	httpRouter := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	httpRouter.GET("/api/v1/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Polls Up and Running..."})
	})

	//url := ginSwagger.URL("http://localhost:8000/docs/swagger.json")

	//url := ginSwagger.URL("http://localhost:8000/docs/swagger.json") // The url pointing to API definition
	//httpRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	httpRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return GinRouter{
		Gin: httpRouter,
	}

}
