package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pyhita/go-learning/go-web/gin/gin-quickstart/binding"
	"github.com/pyhita/go-learning/go-web/gin/gin-quickstart/render"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	RegisterRouters(r)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run("127.0.0.1:8888")
}

func RegisterRouters(r *gin.Engine) {
	// register render routes
	renderAPI := r.Group("/render")
	renderHandler := &render.RenderHandler{}
	renderAPI.GET("/json", renderHandler.JSON)

	// register binding routes
	bindingAPI := r.Group("/binding")
	bindingHandler := &binding.BindingHandler{}
	bindingAPI.GET("/query", bindingHandler.BindingQuryString)
	bindingAPI.GET("/queryv2", bindingHandler.BindingQueryStringV2)
}
