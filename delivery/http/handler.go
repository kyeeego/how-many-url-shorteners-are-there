package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kyeeego/how-many-url-shorteners-are-there/service"
)

type Handler struct {
	services *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	a := router.Group("/api")
	{
		a.POST("/shorten", h.HandleShorten)
		a.GET("/:url", h.HandleRedirect)
		a.GET("/:url/views", h.HandleGetViews)
	}
}
