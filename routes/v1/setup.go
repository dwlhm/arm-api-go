package v1

import (
	"github.com/gin-gonic/gin"
)

var calendar *gin.RouterGroup

func Setup(g *gin.RouterGroup) {

	calendar = g.Group("/calendar")

	calendarRoute()

}
