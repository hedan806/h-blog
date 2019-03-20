package routers

import (
	"github.com/gin-gonic/gin"
	ctl "h-blog/server/controllers"
)

func SetRouters(app *gin.Engine) {
	// 前端CMS页面路由.
	setCmsRouter(app.Group("/"))

	// 后台OM管理路由.
	//setOmRouter(app.Group("/om"))

	// Api接口相关路由.
	//setApiRouter(app.Group("/api"))
}

func setCmsRouter(g *gin.RouterGroup) {
	g.GET("/", ctl.CmsDefault)
	g.GET("/about", ctl.About)
	g.GET("/golang", ctl.Golang)
	//g.GET("/articles", ctl.CmsArticles)
	//g.GET("/articles/:categoryId", ctl.CmsArticles)
	//g.GET("/search/articles", ctl.CmsSearch)
	//g.GET("/article/:articleId", ctl.CmsDetail)
}
