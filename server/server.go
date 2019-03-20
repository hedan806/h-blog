package server

import (
	"github.com/gin-gonic/gin"
	"h-blog/server/middlewares"
	"h-blog/server/routers"
	"net/http"
)

func Start() {
	r := gin.Default()

	r.Static("/assets", "server/views/assets")
	r.StaticFS("/more_static", http.Dir("assets"))
	r.StaticFile("/favicon.ico", "./favicon.ico")

	middlewares.SetEnv(r)

	routers.SetRouters(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
