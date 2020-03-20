package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"yisany.top/milu/dslConvert/routers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := routers.InitRouter()
	endless.ListenAndServe(":3232", router)
}
