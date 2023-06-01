package http

import (
	"fmt"
	"net/http"

	"github.com/adYushinW/URL-shorter/internal/app"
	"github.com/gin-gonic/gin"
)

func Service(app *app.App) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("/short", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.GET("/:hash", func(c *gin.Context) {
		hash, ok := c.Params.Get("hash")
		if !ok {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("я получил хэш: %v", hash))
	})

	return r.Run(":8080")
}
