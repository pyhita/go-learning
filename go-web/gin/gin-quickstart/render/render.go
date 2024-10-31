package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 测试gin 框架渲染不同的内容

type RenderHandler struct {
}

func (r *RenderHandler) StaticHtml(ctx *gin.Context) {

}

func (r *RenderHandler) JSON(ctx *gin.Context) {
	var msg struct {
		Name  string `json:"user"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	msg.Name = "kante.yang"
	msg.Email = "kante.yang@shopee.com"
	msg.Age = 23

	ctx.JSON(http.StatusOK, &msg)
}
