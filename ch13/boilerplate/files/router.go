package files

// ROUTER is the boilerplate code of router package
const ROUTER = `import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter initials all routers
func InitRouter(engine *gin.Engine) http.Handler {

	return engine
}`
