package http

import (
	Server "simpel-gateway/internal/app/server"
	"simpel-gateway/internal/factory"
	"simpel-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Here we define route function for user Handlers that accepts gin.Engine and factory parameters
func NewHttp(g *gin.Engine, f *factory.Factory) {

	Index(g)
	// Here we use logger middleware before the actual API to catch any api call from clients
	g.Use(gin.Logger())
	// Here we use the recovery middleware to catch a panic, if panic occurs recover the application witohut shutting it off
	g.Use(gin.Recovery())

	g.Use(middleware.CORSMiddleware())

	// Here we define a router group
	v1 := g.Group("/api/v1")
	Server.NewHandler(f).Router(v1.Group("/server"))

}

func Index(g *gin.Engine) {
	g.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			Name string `json:"name"`
		}{
			Name: "Simpel Gateway",
		})
	})
}
