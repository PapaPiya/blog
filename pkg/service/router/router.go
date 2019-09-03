package router

import (
	"beegolearn/pkg/service/controller"
	"beegolearn/pkg/service/midderware"
	"github.com/gin-gonic/gin"
)

// New 初始化路由
func New(e *gin.Engine) {
	// 注册全局的中间件
	e.Use(gin.Logger(), midderware.Recovery)

	e.POST("/webhook", controller.Webhook)

	// 前台页面组，添加右侧标签的中间件
	front := e.Group("/", midderware.Navigation, midderware.Tags)
	{
		// 出现错误的页面
		e.GET("/errors", midderware.Errors)

		// 404 页面
		e.NoRoute(midderware.NotFound)

		// 首页
		front.GET("/", controller.Home)

		// about 页
		front.GET("/about", controller.About)

		// 博客文章详情
		front.GET("/posts/:path", controller.PostByPath)

		// 根据分类显示文章
		front.GET("/categories/:route", controller.PostsByCategory)

		// 根据标签显示文章
		front.GET("/tag/:name", controller.PostsByTag)
	}
}
