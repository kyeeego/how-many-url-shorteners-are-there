package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kyeeego/how-many-url-shorteners-are-there/domain"
)

func (h *Handler) HandleShorten(c *gin.Context) {
	var body domain.ShortenDto
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, map[string]error{"error": err})
		return
	}

	model, err := h.services.UrlService.Shorten(body.UrlToShorten)
	if err != nil {
		c.JSON(500, map[string]error{"error": err})
		return
	}

	c.JSON(201, map[string]string{"shortened_url": fmt.Sprintf("https://kyeeego.dev/%s\n", model.Token)})
}

func (h *Handler) HandleRedirect(c *gin.Context) {
	token := c.Param("url")

	model, err := h.services.UrlService.GetUrlByToken(token)
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
	}

	model.Views++
	err = h.services.UrlService.Update(model)
	if err != nil {
		c.JSON(500, map[string]error{"error": err})
		return
	}

	c.Redirect(301, model.OriginalUrl)
}

func (h *Handler) HandleGetViews(c *gin.Context) {
	token := c.Param("url")

	model, err := h.services.UrlService.GetUrlByToken(token)
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
	}

	c.JSON(200, map[string]int{"views": model.Views})
}
