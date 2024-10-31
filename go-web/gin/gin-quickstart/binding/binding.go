package binding

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 测试 gin 绑定参数

type BindingHandler struct {
}

type Person struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

// ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据
func (b *BindingHandler) BindingQuryString(c *gin.Context) {
	var person Person
	c.ShouldBindQuery(&person)

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": person.Name,
		"address":  person.Address,
	})
}

// // 示例 URL： /welcome?firstname=Jane&lastname=Doe
func (b *BindingHandler) BindingQueryStringV2(c *gin.Context) {
	req := &Person{}
	req.FirstName = c.DefaultQuery("firstname", "kante")
	req.LastName = c.Query("lastname")

	c.JSON(http.StatusOK, req)
}
