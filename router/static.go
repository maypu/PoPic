package router

import (
	"PoPic/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Static(r *gin.Engine) *gin.Engine {
	// static
	r.Static("/static", "./web/static")
	r.Static("/dist", "./web/dist")

	//html template
	r.LoadHTMLGlob("./web/view/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": utils.GetConfig("common.name"),
		})
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"Title": utils.GetConfig("common.name"),
		})
	})

	return r
}
