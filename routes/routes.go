package routes

import (
	"go-gin-amis/app/Http/Controllers"

	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Admin AdminRouter
	Web   WebRouter
}
type WebRouter struct {
	Home Controllers.HomeController
}
type AdminRouter struct {
	Login     Controllers.LoginController
	Admin     Controllers.AdminController
	Articles  Controllers.ArticlesController
	Users     Controllers.UsersController
	AmisPages Controllers.AmisPagesController
}

var RouterGroupApp = new(RouterGroup)

func Init(Router *gin.Engine) {
	RouterGroupApp.Admin.Articles.Init()
	RouterGroupApp.Admin.Users.Init()

	Router.GET("/", RouterGroupApp.Web.Home.Index)
	{
		Router.GET("/admin/login", RouterGroupApp.Admin.Login.GetLogin)
		Router.POST("/admin/login", RouterGroupApp.Admin.Login.PostLogin)
		Router.GET("/admin/logout", RouterGroupApp.Admin.Login.GetLogout)

		Router.GET("/page/:id", RouterGroupApp.Admin.AmisPages.Show) //详情

		Auth := Router.Group("/admin")
		// Auth := Router.Group("/admin").Use(Middleware.JwtVerifyMiddleware("/admin/login"))
		{
			Auth.GET("/", RouterGroupApp.Admin.Admin.Index)
			Auth.GET("/userinfo", RouterGroupApp.Admin.Login.GetUserInfo)
			//页面管理路由
			Auth.GET("/amis-pages", RouterGroupApp.Admin.AmisPages.Index)          //列表
			Auth.GET("/amis-pages/create", RouterGroupApp.Admin.AmisPages.Create)  //创建页面
			Auth.POST("/amis-pages", RouterGroupApp.Admin.AmisPages.Store)         //保存数据
			Auth.GET("/amis-pages/:id", RouterGroupApp.Admin.AmisPages.Show)       //详情
			Auth.GET("/amis-pages/:id/edit", RouterGroupApp.Admin.AmisPages.Edit)  //编辑页面
			Auth.PUT("/amis-pages/:id", RouterGroupApp.Admin.AmisPages.Update)     //更新数据
			Auth.DELETE("/amis-pages/:id", RouterGroupApp.Admin.AmisPages.Destroy) //删除数据
			// 文章管理路由
			Auth.GET("/articles", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.Index))                         //列表
			Auth.GET("/articles/create", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.Create))                 //创建页面
			Auth.POST("/articles", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.Store))                        //保存数据
			Auth.GET("/articles/:id", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.Show))                      //详情
			Auth.GET("/articles/:id/edit", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.Edit))                 //编辑页面
			Auth.PUT("/articles/:id", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.Update))                    //更新数据
			Auth.DELETE("/articles/:id", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.Destroy))                //删除数据
			Auth.DELETE("/articles/bulkDelete/:ids", egin.AutoVerifyHandler(RouterGroupApp.Admin.Articles.BulkDelete)) //删除数据
			//用户管理路由
			// 文章管理路由
			Auth.GET("/users", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.Index))                         //列表
			Auth.GET("/users/create", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.Create))                 //创建页面
			Auth.POST("/users", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.Store))                        //保存数据
			Auth.GET("/users/:id", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.Show))                      //详情
			Auth.GET("/users/:id/edit", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.Edit))                 //编辑页面
			Auth.PUT("/users/:id", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.Update))                    //更新数据
			Auth.DELETE("/users/:id", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.Destroy))                //删除数据
			Auth.DELETE("/users/bulkDelete/:ids", egin.AutoVerifyHandler(RouterGroupApp.Admin.Users.BulkDelete)) //删除数据
		}
	}

}
