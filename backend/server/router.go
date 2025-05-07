package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(corsMiddleware)

	return r
}

func corsMiddleware(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	reffer := c.Request.Header.Get("Referer")

	if origin == "" {
		origin = reffer
	}

	if c.Request.Method == http.MethodOptions || c.Request.Method == http.MethodHead {
		h := c.Writer.Header()
		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD")
		h.Set("Access-Control-Allow-Headers", "authorization,content-type")
		h.Set("Access-Control-Max-Age", "86400")
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	h := c.Writer.Header()
	h.Set("Access-Control-Allow-Origin", origin)
	c.Next()
}
