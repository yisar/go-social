package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cliclitv/htwxc/service"
	"net/http"
	"embed"
	"io/fs"
)


//go:embed fre/dist
var embededFiles embed.FS

//go:embed fre/dist/index.html
var html string

var whiteOrigins = [5]string{
	"https://www.cuipiya.net",
	"https://cuipiya.net",
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

	r.POST("/login", service.Login)
	r.POST("/register", service.Register)
	r.POST("/sendcode", service.SendCode)
	r.POST("/thread/add", service.InsertThread)
	r.POST("/post/add", service.InsertPost)
	
	r.GET("/thread/detail/:id", service.ThreadDetail)
	r.GET("/author/detail/:id", service.UserDetail)
	r.GET("/post/detail/:id", service.PostDetail)
	r.GET("/threads", service.GetThreads)
	r.GET("/posts", service.GetPosts)

	return r
}


func main() {
	for _, s := range whiteOrigins {
		whiteOriginsSet[s] = true
	}
	e := Router()
	e.Run(":5000")
}
