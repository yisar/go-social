package service

import (
	"fmt"
	"github.com/cliclitv/htwxc/helper"
	"github.com/cliclitv/htwxc/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	json := model.Author{}
	c.BindJSON(&json)

	if json.Name == "" || json.Pwd == "" {
		c.JSON(http.StatusOK, gin.H{
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
	token, err := helper.GenerateToken(author.Identity.Hex(), author.Level)
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
			"author": gin.H{
				"name":  author.Name,
				"email": author.Email,
				"_id":   author.Identity,
			},
		},
	})
}

func Register(c *gin.Context) {
	json := model.Author{}
	c.BindJSON(&json)
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
		Name:  json.Name,
		Pwd:   helper.GetMd5(json.Pwd),
		Email: json.Email,
		Level: 0,
	}

	
	if json.Identity.Hex() == "000000000000000000000000" {
		err = model.InsertAuthor(ub)
	} else {

		err = model.UpdateAuthor(ub, json.Identity)
	}


	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func UserDetail(c *gin.Context) {
	id := c.Param("id")
	// uc := u.(*helper.UserClaims)
	oid, _ := primitive.ObjectIDFromHex(id)
	author, err := model.GetAuthorByIdentity(oid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": gin.H{
			"name":  author.Name,
			"email": author.Email,
			"_id":   author.Identity,
		},
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
