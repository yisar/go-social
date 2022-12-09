package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yisar/nugei/service"
	"github.com/yisar/nugei/helper"
	"net/http"
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


func Router() *gin.Engine {
	r := gin.Default()

	// 用户登录
	r.POST("/login", service.Login)
	r.POST("/register", service.Regisiter)
	// 发送验证码
	r.POST("/sendcode", service.SendCode)

	auth := r.Group("/author", AuthCheck())

	// 用户详情
	auth.GET("/detail", service.UserDetail)

	return r
}


func main() {
	e := Router()
	e.Run(":8080")
}
