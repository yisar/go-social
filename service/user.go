package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cliclitv/htwxc/helper"
	"github.com/cliclitv/htwxc/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *gin.Context) {
	json := model.User{}
	c.BindJSON(&json)

	if json.Name == "" || json.Pwd == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	user, err := model.GetUserByAccountPassword(json.Name, helper.GetMd5(json.Pwd))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := helper.GenerateToken(user.Identity.Hex(), user.Name, user.Level)
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
			"user": gin.H{
				"name":  user.Name,
				"email": user.Email,
				"_id":   user.Identity,
			},
		},
	})
}

func Register(c *gin.Context) {
	json := model.User{}
	c.ShouldBind(&json)
	if json.Name == "" || json.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "都是必填的！",
		})
		return
	}

	if json.Pwd == "" && json.Identity.Hex() != "000000000000000000000000" {

		// 编辑状态
		user, err := model.GetUserByIdentity(json.Identity)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("%s", err),
			})
			return
		}
		token := c.GetHeader("token")
		err = Auth(user.Identity.Hex(), token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("%s", err),
			})
			return
		}
		json.Pwd = user.Pwd

		err = model.UpdateUser(&model.User{
			Name:  json.Name,
			Pwd:   helper.GetMd5(json.Pwd),
			Email: json.Email,
			Level: 0,
		}, json.Identity)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("%s", err),
			})
			return
		}

	} else {
		//注册状态
		cnt, err := model.GetUserCountByEmail(json.Email)
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

		cnt2, err := model.GetUserCountByName(json.Name)
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

		err = model.InsertUser(&model.User{
			Name:  json.Name,
			Pwd:   helper.GetMd5(json.Pwd),
			Email: json.Email,
			Level: 0,
		})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("%s", err),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}

func UserDetail(c *gin.Context) {
	id := c.Param("id")
	// uc := u.(*helper.UserClaims)
	oid, _ := primitive.ObjectIDFromHex(id)
	user, err := model.GetUserByIdentity(oid)
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
			"name":  user.Name,
			"email": user.Email,
			"_id":   user.Identity,
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
	cnt, err := model.GetUserCountByEmail(email)
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
