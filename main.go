package main

import (
	"go-gin-amis/config"
	"go-gin-amis/routes"

	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载.env文件
	ecore.E加载环境变量_从文件(".env")
	config.InitDB()

	gin.SetMode(gin.DebugMode)

	// 创建app结构体对象
	r := gin.New()
	// 注册日志中间件
	//logger.E初始化Gin日志类("./log/elog.log", "debug")
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// r.Use(Middleware.JsonFormatMiddleware(false, "/admin"))
	// 静态文件
	egin.E绑定静态文件目录(r, "./public")
	// 注册模板引擎
	r.LoadHTMLGlob("./resources/views/*")
	//中间件
	//r.Use(Middlewre.LoggerMiddleware())
	//	//路由注册a
	routes.Init(r)

	r.Run("0.0.0.0:" + ecore.Env("APP_PORT", "8080"))
}
