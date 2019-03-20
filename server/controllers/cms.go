package controllers

import "github.com/gin-gonic/gin"

func CmsDefault(c *gin.Context) {
	//module  := modules.NewArticle()
	// get the article list data.
	//_, list := module.GetList(models.M{"state": "pass"}, 1, 8, "created desc", true)

	// get the hot article list data.
	//_, hotList := module.GetList(models.M{"state": "pass"}, 1, 3, "commented desc, liked desc", false)

	params := gin.H{
		"title":     "登录管理后台",
		"customCss": "cms/css/index.css",
		"breadcrumbs": []gin.H{
			{"title": "首页", "link": ""},
		},
		//"list": list,
		//"hotList": hotList,
	}
	c.HTML(200, "index.html", params)
}

func About(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{})
}

func Golang(c *gin.Context) {
	c.HTML(200, "golang.html", gin.H{})
}
