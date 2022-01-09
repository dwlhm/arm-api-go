package main

import (
	"arm_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := routes.Setup(gin.Default())

	r.Run()
}
