package ohisrouter

import (
	"github.com/gin-gonic/gin"
)

type ginRegistrations func(*gin.Engine)

var topRegistrations = make([]ginRegistrations, 0, 10)

// InitRouter initials all routers
func InitRouter(engine *gin.Engine) {
	for _, rs := range topRegistrations {
		rs(engine)
	}
}
