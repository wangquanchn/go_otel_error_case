package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 gin 对象实例
	r := gin.New()

	// 健康检查 URL
	r.GET("/healthCheck", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	_ = r.Run(":8088")
}
