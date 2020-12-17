package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
}

// InitRoutes ...
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		images := api.Group("/images")
		{
			images.POST("/", h.createASCII)
		}
	}

	return router
}
