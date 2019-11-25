package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter initials all routers
func InitRouter(engine *gin.Engine) http.Handler {
	index(engine)

	// Gin Examples
	somejson(engine)
	xml(engine)
	yaml(engine)
	protobuf(engine)
	formpost(engine)
	upload(engine)
	download(engine)
	auth(engine)

	return engine
}
