package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adYushinW/URL-shorter/internal/app"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	app *app.App
}

func NewController(app *app.App) *Controller {
	return &Controller{
		app: app,
	}
}

func (c *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

func (c *Controller) ShortLinkGen(ctx *gin.Context) {
	var i interface{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&i)
	if err != nil {
		return
	}
	shortHash, err := c.app.SetHash(i.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	shorterURL := fmt.Sprintf("https://%s/%s", "short.ss", shortHash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, shorterURL)

}

func (c *Controller) GetLink(ctx *gin.Context) {
	hash, ok := ctx.Params.Get("hash")
	if !ok {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	link, err := c.app.GetHash(hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusFound, link)
	ctx.Redirect(http.StatusFound, link)
}
