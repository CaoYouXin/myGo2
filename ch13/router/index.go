package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// index is the "/" router
func index(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server\n")
	})
}
