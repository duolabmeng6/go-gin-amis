package Controllers

import (
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func (b HomeController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
