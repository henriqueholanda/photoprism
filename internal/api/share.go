package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/service"
)

// GET /s/:token/...
func Shares(router *gin.RouterGroup) {
	router.GET("/:token", func(c *gin.Context) {
		conf := service.Config()

		shareToken := c.Param("token")

		links := entity.FindValidLinks(shareToken, "")

		if len(links) == 0 {
			log.Warn("share: invalid token")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		clientConfig := conf.GuestConfig()

		c.HTML(http.StatusOK, "share.tmpl", gin.H{"config": clientConfig, "previewUrl": clientConfig.SitePreview})
	})

	router.GET("/:token/:uid", func(c *gin.Context) {
		conf := service.Config()

		shareToken := c.Param("token")
		shareUID := c.Param("uid")

		links := entity.FindValidLinks(shareToken, shareUID)

		if len(links) != 1 {
			log.Warn("share: invalid token or uid")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		clientConfig := conf.GuestConfig()
		sharePreview := fmt.Sprintf("%ss/%s/%s/preview", clientConfig.SiteUrl, shareToken, shareUID)

		c.HTML(http.StatusOK, "share.tmpl", gin.H{"config": clientConfig, "previewUrl": sharePreview})
	})
}
