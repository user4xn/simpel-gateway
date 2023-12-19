package server

import (
	"simpel-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) Router(g *gin.RouterGroup) {
	g.Use(middleware.AuthBearerServer())
	g.POST("/store", h.Store)
	g.GET("/list", h.GetServers)
	g.POST("/check-ip", h.CheckServerIp)
}
