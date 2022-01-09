package routes

import (
	v1 "arm_go/routes/v1"

	"github.com/gin-gonic/gin"
)

func Setup(e *gin.Engine) *gin.Engine {

	v1.Setup(e.Group("/v1"))

	return e
}
