package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Service() error {
	r := gin.Default()

	gin.Default().GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("/short")
	r.GET("/:hash")

	return r.Run(":8080")
}
