package Controllers

import (
	"github.com/duolabmeng6/goefun/egin"
	"github.com/duolabmeng6/goefun/egin/jwt"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func (b *LoginController) GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func (b *LoginController) PostLogin(c *gin.Context) {
	username := egin.I(c, "username", "")
	password := egin.I(c, "password", "")
	if username == "admin" && password == "123456" {
		//创建 jwt token
		token := jwt.GenerateToken(1)
		//发送 cookie jwt token
		c.SetCookie("jwt", token, 3600, "/", "", false, true)

		c.JSON(200, gin.H{
			"status": 0,
			"msg":    "登录成功",
			"data": gin.H{
				"jwt": token,
			},
		})
	} else {
		c.JSON(401, gin.H{
			"status": 1,
			"msg":    "登录失败",
		})
	}

}

func (b *LoginController) GetLogout(c *gin.Context) {
	//清除cookie 退出登录 jwt
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	//跳转到登录页面
	c.Redirect(302, "/admin/login")

}

func (b *LoginController) GetUserInfo(c *gin.Context) {
	//获取jwt中的id输出
	uid, _ := c.Get("uid")

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "获取成功",
		"data": gin.H{
			"id": uid,
		},
	})
}
