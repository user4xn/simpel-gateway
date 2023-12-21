package apksetting

import (
	"simpel-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.Use(middleware.AuthBearerServer())
	g.GET("/first", h.FindOne)
	g.POST("/store", h.Store)
}
