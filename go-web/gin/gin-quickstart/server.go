package main

import (
	"log"
	"time"

	"github.com/pyhita/go-learning/go-web/gin/gin-quickstart/handler/binding"
	"github.com/pyhita/go-learning/go-web/gin/gin-quickstart/handler/render"

	"github.com/gin-gonic/gin"
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
	v1API := r.Group("/v1")
	v2API := r.Group("/v2")
	renderHandler := &render.RenderHandler{}
	renderAPI.GET("/json", renderHandler.JSON)

	// register binding routes
	v1BindingAPI := v1API.Group("/binding")
	bindingHandler := &binding.BindingHandler{}
	v1BindingAPI.GET("/query", Logger(), bindingHandler.BindingQueryString)
	v1BindingAPI.POST("/form", bindingHandler.BindingForm)
	v1BindingAPI.POST("/json", bindingHandler.BindingJson)
	v1BindingAPI.POST("/header", bindingHandler.BindingHeader)
	v1BindingAPI.GET("/path/:name", bindingHandler.BindingPath)

	v2BindingAPI := v2API.Group("/binding")
	v2BindingAPI.GET("/query", bindingHandler.BindingQuryStringV2)
	v2BindingAPI.GET("/path/:name", bindingHandler.BindingPathV2)
	v2BindingAPI.POST("/header", bindingHandler.BindingHeaderV2)
}
