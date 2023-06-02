package http

import (
	"encoding/json"
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

	r.POST("/site", func(c *gin.Context) {
		var i interface{}
		err := json.NewDecoder(c.Request.Body).Decode(&i)
		if err != nil {
			return
		}
		shortHash, err := app.SetHash(i.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		shorterURL := fmt.Sprintf("https://%s/%s", "short.ss", shortHash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, shorterURL)

	})

	r.GET("/:hash", func(c *gin.Context) {
		hash, ok := c.Params.Get("hash")
		if !ok {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		link, err := app.GetHash(hash)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusFound, link)
		c.Redirect(http.StatusFound, link)
	})

	return r.Run(":8080")
}
