package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ReverseProxy(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		target := viper.GetString("services."+serviceName)
		if target == "" {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Unknown service"})
			return
		}

		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
			return
		}

		// Strip only `/api` prefix, keep serviceName
		originalPath := c.Request.URL.Path
		strippedPath := strings.TrimPrefix(originalPath, "/api")
		if strippedPath == "" {
			strippedPath = "/" // fallback if empty
		}
		c.Request.URL.Path = strippedPath

		c.Request.Host = remote.Host
		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
