package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Service() error {
	r := gin.Default()

	r.GET("/ping", getPing)

	r.POST("/short", postShort)

	r.GET("/:hash", getHash)

	return r.Run(":8080")
}

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func postShort(c *gin.Context) {
	c.JSON(http.StatusOK, "я верну тебе короткую ссылку")
}

func getHash(c *gin.Context) {
	hash, ok := c.Params.Get("hash")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("я получил хэш: %v", hash))
}
