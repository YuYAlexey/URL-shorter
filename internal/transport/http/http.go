package http

import (
	"github.com/adYushinW/URL-shorter/internal/app"
	"github.com/gin-gonic/gin"
)

func Service(app *app.App) error {

	c := NewController(app)

	r := gin.Default()

	r.GET("/ping", c.Ping)

	r.POST("/site", c.ShortLinkGen)

	r.GET("/:hash", c.GetLink)

	return r.Run(":8080")
}
