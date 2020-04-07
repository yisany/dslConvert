package apps

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	elasticsql "yisany.top/milu/dslConvert/convert"
)

type SqlRequtest struct {
	Sql string `json:"sql"`
}

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func SQLConvert(c *gin.Context) {
	var req SqlRequtest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 失败
		send(false, nil, err, c)
	} else {
		// 成功
		sql := strings.ReplaceAll(req.Sql, "\n", " ")
		dsl, _, err := elasticsql.ConvertPretty(sql)
		if err != nil {
			send(false, nil, err, c)
		} else {
			send(true, dsl, err, c)
		}
	}

}

func send(isSuccess bool, m interface{}, e error, c *gin.Context) {
	if isSuccess {
		c.JSON(http.StatusOK, Response{
			Status:  http.StatusOK,
			Data:    m,
			Message: "ok~",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: e.Error(),
		})
	}
}
