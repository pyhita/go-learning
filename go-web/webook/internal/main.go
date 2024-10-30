package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pyhita/webbook/internal/web"
)

func main() {

	server := gin.Default()

	user := &web.UserHandler{}
	user.RegisterRoutes(server)

	server.Run(":8888")
}
