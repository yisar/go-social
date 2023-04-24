package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yisar/footsie/svc"
	"net/http"
	"embed"
	"io/fs"
)


//go:embed fre/dist
var embededFiles embed.FS

//go:embed fre/dist/index.html
var html string

var whiteOrigins = [5]string{
	"http://localhost:3000",
}

var whiteOriginsSet = make(map[string]bool)

func initMiddleware(c *gin.Context) {
	origin := c.GetHeader("Origin")
	if whiteOriginsSet[origin] {
		c.Header("Access-Control-Allow-Origin", origin)
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, token")
	if c.Request.Method == "OPTIONS" {
        c.JSON(http.StatusOK,"ok")
    }
	c.Next()
}


func Router() *gin.Engine {
	r := gin.Default()

	r.Use(initMiddleware)

	fsys, _ := fs.Sub(embededFiles, "fre/dist")

	r.StaticFS("/assets", http.FS(fsys))

	r.NoRoute(func(c *gin.Context) {
        c.Header("Content-Type", "text/html; charset=utf-8")
    	c.String(200, html)
    })

	r.POST("/user/login", service.Login)
	r.POST("/user/register", service.Register)
	r.POST("/user/sendcode", service.SendCode)
	r.GET("/user/detail/:id", service.UserDetail)
	r.GET("/users", service.GetUsers)

	return r
}


func main() {
	for _, s := range whiteOrigins {
		whiteOriginsSet[s] = true
	}
	e := Router()
	e.Run(":5000")
}
