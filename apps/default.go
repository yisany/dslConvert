package apps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "dslConvert",
	})
}
