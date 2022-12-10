package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cliclitv/htwxc/service"
	"github.com/cliclitv/htwxc/helper"
	"net/http"
	"embed"
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
	"https://admin.clicli.cc",
	"https://www.clicli.cc",
	"https://clicli.cc",
	"http://localhost:3000",
}

var whiteOriginsSet = make(map[string]bool)


func Router() *gin.Engine {
	r := gin.Default()

	r.StaticFS("assets", http.FS(embededFiles))

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
