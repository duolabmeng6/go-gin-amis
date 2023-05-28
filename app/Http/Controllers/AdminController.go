package Controllers

import (
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func (b *AdminController) Index(c *gin.Context) {
	c.HTML(200, "admin.html", gin.H{})
}
