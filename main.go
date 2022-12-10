package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cliclitv/htwxc/service"
	"github.com/cliclitv/htwxc/helper"
	"net/http"
	"embed"
	"io/fs"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证不通过",
			})
			return
		}
		c.Set("user", userClaims)
		c.Next()
	}
}

//go:embed fre/dist
var embededFiles embed.FS

//go:embed fre/dist/index.html
var html string

var whiteOrigins = [5]string{
	"https://www.htwxc.com",
	"https://htwxc.com",
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
        c.AbortWithStatus(http.StatusNoContent)
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

	auth := r.Group("/author", AuthCheck())
	auth.GET("/detail", service.UserDetail)

	return r
}


func main() {
	for _, s := range whiteOrigins {
		whiteOriginsSet[s] = true
	}
	e := Router()
	e.Run(":5000")
}
