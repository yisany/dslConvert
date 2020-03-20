package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yisany.top/milu/dslConvert/apps"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 设置静态资源地址
	router.StaticFS("/static", http.Dir("static"))
	// 设置页面文件
	router.LoadHTMLGlob("templates/*")

	router.GET("/", apps.IndexApi)

	v1 := router.Group("/v1")
	{
		v1.POST("/convert", apps.SQLConvert)
	}

	return router
}
