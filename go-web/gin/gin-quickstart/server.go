package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	// add middleware
	r.Use(func(c *gin.Context) {
		println("this is the first middleware")
	}, func(c *gin.Context) {
		println("this is the second middleware")
	})
	r.Use(Logger())

	// test custom middleware
	r.GET("/middleware", func(c *gin.Context) {
		str := c.MustGet("example").(string)
		log.Println(str)
	})

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/healthy", func(context *gin.Context) {
		// c.JSON：返回JSON格式的数据
		context.JSON(200, gin.H{
			"result": "success",
		})
	})

	// gin 参数路由
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.Writer.WriteString("username is " + name)
	})

	// gin 查询参数
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.Query("firstname")
		lastname := c.DefaultQuery("lastname", "Gopher")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run("127.0.0.1:8888")
}
