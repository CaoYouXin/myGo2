package ohisrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	topRegistrations = append(topRegistrations, validOrderInit)
}

func validOrderInit(engine *gin.Engine) {
	engine.GET("/auth", func(c *gin.Context) {
		c.String(http.StatusOK, "auth successed! \n")
	})
}
