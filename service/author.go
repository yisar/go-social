package service

import (
	"github.com/gin-gonic/gin"
	"github.com/cliclitv/htwxc/helper"
	"github.com/cliclitv/htwxc/model"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	json := model.Author{}
 	c.BindJSON(&json)
 	
	if json.Name == "" || json.Pwd == "" {
		c.JSON(http.StatusOK, gin.H {
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	author, err := model.GetAuthorByAccountPassword(json.Name, helper.GetMd5(json.Pwd))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := helper.GenerateToken(author.Identity, author.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func Register(c *gin.Context) {
	json := model.Author{}
 	c.BindJSON(&json)
 	log.Printf("%v",&json)
	if json.Name == "" || json.Pwd == "" || json.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "都是必填的！",
		})
		return
	}
	cnt, err := model.GetAuthorCountByEmail(json.Email)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已被注册",
		})
		return
	}

	cnt2, err := model.GetAuthorCountByName(json.Name)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return
	}
	if cnt2 > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前笔名已被注册",
		})
		return
	}

	ub := &model.Author{
		Name: json.Name,
		Pwd: json.Pwd,
		Email: json.Email,
	}

	err2 := model.UpdateAuthor(ub)

	if err2 != nil {
		log.Printf("[DB ERROR]:%v\n", err2)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func UserDetail(c *gin.Context) {
	u, _ := c.Get("user")
	log.Printf("[DB ERROR]:%v\n", u)
	uc := u.(*helper.UserClaims)
	author, err := model.GetAuthorByIdentity(uc.Identity)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": author,
	})
}

func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	cnt, err := model.GetAuthorCountByEmail(email)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已被注册",
		})
		return
	}
	err = helper.SendCode(email, "666666")
	if err != nil {
		log.Printf("[ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}
