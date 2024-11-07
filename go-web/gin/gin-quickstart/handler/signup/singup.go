package signup

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpReq struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type SignupHandler struct {
}

func (s *SignupHandler) Signup(c *gin.Context) {
	var req SignUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store in db
	c.JSON(http.StatusOK, "success")
}
