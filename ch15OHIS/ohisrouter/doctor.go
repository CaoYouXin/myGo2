package ohisrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	topRegistrations = append(topRegistrations, ping)
}

func ping(engine *gin.Engine) {
	engine.GET("/doc/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "hello doc \n")
	})
}
