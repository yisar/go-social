package service

import (
	"github.com/gin-gonic/gin"
	"github.com/cliclitv/htwxc/helper"
	"github.com/cliclitv/htwxc/model"
	"log"
	"net/http"
	"time"
)

func InsertChapter(c *gin.Context) {
	json := model.Chapter{}
 	c.BindJSON(&json)
 	log.Printf("%v",&json)
	if json.Title == "" || json.Content == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "都是必填的！",
		})
		return
	}
	cnt, err := model.GetChapterCountByName(json.Title)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "重名了",
		})
		return
	}

	chapter := &model.Chapter{
		Title: json.Title,
		Content: json.Content,
		Time: time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04"),
	}

	err2 := model.UpdateChapter(chapter)

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
		"msg":  "投递成功",
	})
}

func ChapterDetail(c *gin.Context) {
	u, _ := c.Get("user")
	log.Printf("[DB ERROR]:%v\n", u)
	uc := u.(*helper.UserClaims)
	author, err := model.GetChapterByIdentity(uc.Identity)
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