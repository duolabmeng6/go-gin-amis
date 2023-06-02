package routes

import (
	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
	"go-gin-amis/app/Http/Controllers"
)

func Init(Router *gin.Engine) {
	var HomeController Controllers.HomeController
	var LoginController Controllers.LoginController
	var AdminController Controllers.AdminController

	Router.GET("/", HomeController.Index)
	{
		Router.GET("/admin/login", LoginController.GetLogin)
		Router.POST("/admin/login", LoginController.PostLogin)
		Router.GET("/admin/logout", LoginController.GetLogout)

		var AutoController Controllers.AutoController
		AutoController.Init()
		Router.GET("/auto/get/:table_name", AutoController.Get)
		Router.POST("/auto/created", egin.AutoVerifyHandler(AutoController.Store))
		Router.GET("/auto/created", AutoController.Store)
		Router.GET("/auto/get_all_table_name", AutoController.GetAllTableName)

		Auth := Router.Group("/admin")
		// Auth := Router.Group("/admin").Use(Middleware.JwtVerifyMiddleware("/admin/login"))
		{
			Auth.GET("/", AdminController.Index)
			Auth.GET("/userinfo", LoginController.GetUserInfo)
			//页面管理路由
			var AmisPagesController Controllers.AmisPagesController
			Auth.GET("/amis-pages", AmisPagesController.Index)          //列表
			Auth.GET("/amis-pages/create", AmisPagesController.Create)  //创建页面
			Auth.POST("/amis-pages", AmisPagesController.Store)         //保存数据
			Auth.GET("/amis-pages/:id", AmisPagesController.Show)       //详情
			Auth.GET("/amis-pages/:id/edit", AmisPagesController.Edit)  //编辑页面
			Auth.PUT("/amis-pages/:id", AmisPagesController.Update)     //更新数据
			Auth.DELETE("/amis-pages/:id", AmisPagesController.Destroy) //删除数据
			Router.GET("/page/:id", AmisPagesController.Show)           //详情

			// 文章管理路由
			var ArticlesController Controllers.ArticlesController
			ArticlesController.Init()
			Auth.GET("/articles", egin.AutoVerifyHandler(ArticlesController.Index))                         //列表
			Auth.GET("/articles/create", egin.AutoVerifyHandler(ArticlesController.Create))                 //创建页面
			Auth.POST("/articles", egin.AutoVerifyHandler(ArticlesController.Store))                        //保存数据
			Auth.GET("/articles/:id", egin.AutoVerifyHandler(ArticlesController.Show))                      //详情
			Auth.GET("/articles/:id/edit", egin.AutoVerifyHandler(ArticlesController.Edit))                 //编辑页面
			Auth.PUT("/articles/:id", egin.AutoVerifyHandler(ArticlesController.Update))                    //更新数据
			Auth.DELETE("/articles/:id", egin.AutoVerifyHandler(ArticlesController.Destroy))                //删除数据
			Auth.DELETE("/articles/bulkDelete/:ids", egin.AutoVerifyHandler(ArticlesController.BulkDelete)) //删除数据

			//用户管理路由
			var UsersController Controllers.UsersController
			UsersController.Init()
			Auth.GET("/users", egin.AutoVerifyHandler(UsersController.Index))                         //列表
			Auth.GET("/users/create", egin.AutoVerifyHandler(UsersController.Create))                 //创建页面
			Auth.POST("/users", egin.AutoVerifyHandler(UsersController.Store))                        //保存数据
			Auth.GET("/users/:id", egin.AutoVerifyHandler(UsersController.Show))                      //详情
			Auth.GET("/users/:id/edit", egin.AutoVerifyHandler(UsersController.Edit))                 //编辑页面
			Auth.PUT("/users/:id", egin.AutoVerifyHandler(UsersController.Update))                    //更新数据
			Auth.DELETE("/users/:id", egin.AutoVerifyHandler(UsersController.Destroy))                //删除数据
			Auth.DELETE("/users/bulkDelete/:ids", egin.AutoVerifyHandler(UsersController.BulkDelete)) //删除数据
		}
	}

}
