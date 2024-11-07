package binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

// 绑定form 表单里面的请求参数
func (b *BindingHandler) BindingForm(c *gin.Context) {
	var req Person

	req.Name = c.PostForm("name")
	req.Address = c.PostForm("address")

	c.JSON(http.StatusOK, req)
}

type UriParam struct {
	Name string `uri:"name"`
}

// 绑定path 参数 /binding/path/:name
func (b *BindingHandler) BindingPath(c *gin.Context) {
	var req UriParam = UriParam{}
	req.Name = c.Param("name")

	c.JSON(http.StatusOK, req)
}

// 绑定path 参数 /binding/path/:name
func (b *BindingHandler) BindingPathV2(c *gin.Context) {
	var req UriParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, req)
}

type FormParam struct {
	Name      string `form:"name"`
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
}

// // 示例 URL： /welcome?firstname=Jane&lastname=Doe
func (b *BindingHandler) BindingQueryString(c *gin.Context) {
	req := &FormParam{}
	req.FirstName = c.DefaultQuery("firstname", "kante")
	req.LastName = c.Query("lastname")

	c.JSON(http.StatusOK, req)
}

// ShouldBindQuery 绑定查询字符串
func (b *BindingHandler) BindingQuryStringV2(c *gin.Context) {
	var formParam FormParam

	if err := c.ShouldBindQuery(&formParam); err != nil {
		c.ShouldBindQuery(formParam)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "ok",
		"firstname": formParam.FirstName,
		"lastname":  formParam.LastName,
	})
}

type HeaderParam struct {
	Firstname string `header:"firstname"`
	Lastname  string `header:"lastname"`
}

func (b *BindingHandler) BindingHeader(c *gin.Context) {
	req := &HeaderParam{}
	req.Firstname = c.GetHeader("firstname")
	req.Lastname = c.GetHeader("lastname")

	c.JSON(http.StatusOK, req)
}

// ShouldBindQuery 绑定查询字符串
func (b *BindingHandler) BindingHeaderV2(c *gin.Context) {
	var headerParam HeaderParam

	if err := c.ShouldBindHeader(&headerParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "ok",
		"firstname": headerParam.Firstname,
		"lastname":  headerParam.Lastname,
	})
}

// 绑定json
func (b *BindingHandler) BindingJson(c *gin.Context) {
	var req Person

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, req)
}

// shouldBind
func (b *BindingHandler) BindingShould(c *gin.Context) {
	//c.ShouldBindJSON()
	//c.ShouldBindQuery()
	//c.ShouldBindUri()
	//c.ShouldBindHeader()
}
