package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Ostrovski/sxtant/sxapi"
)

func SiteHandler(sitesFetcher *sxapi.SitesFetcher) gin.HandlerFunc {
	return func(c *gin.Context) {
		// GET /sites
		// GET /info

		site := c.Param("site")
		if site == "" {
			site = "stackoverflow"
		}

		c.HTML(http.StatusOK, "site.tmpl", gin.H{
			"title": "SXtant - StackExchange instant notifier",
			"sites": sitesFetcher.Sites(),
		})
	}
}
