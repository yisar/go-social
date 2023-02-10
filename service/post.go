package service

import (
	"fmt"
	"github.com/cliclitv/htwxc/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strconv"
	"time"
)

func InsertPost(c *gin.Context) {
	json := model.Post{}
	c.BindJSON(&json)
	if json.Title == "" || json.Content == "" || json.Tid == "" || json.Oid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "都是必填的！",
		})
		return
	}
	cnt, err := model.GetPostCountByName(json.Title)
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

	post := &model.Post{
		Oid:     json.Oid,
		Tid:     json.Tid,
		Status:  json.Status,
		Title:   json.Title,
		Content: json.Content,
		Time:    time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04"),
	}

	if json.Identity.Hex() == "000000000000000000000000" {
		err = model.InsertPost(post)
	} else {
		err = model.UpdatePost(post, json.Identity)
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
		"msg":  "投递成功",
	})
}

func PostDetail(c *gin.Context) {
	id := c.Param("id")
	oid, _ := primitive.ObjectIDFromHex(id)
	post, err := model.GetThreadByIdentity(oid)
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
		"data": post,
	})
}

func GetPosts(c *gin.Context) {
	tid := c.Query("tid")
	// oid, _ := primitive.ObjectIDFromHex(id)

	pageIndex, _ := strconv.ParseInt(c.Query("page"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	skip := (pageIndex - 1) * pageSize
	posts, err := model.GetPosts(&pageSize, &skip, tid)

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
		"data": posts,
	})
}
