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

func InsertThread(c *gin.Context) {
	json := model.Thread{}
	c.BindJSON(&json)
	log.Printf("%v", &json)
	if json.Title == "" || json.Content == "" || json.Sort == "" || json.Tag == "" || json.Uid == "" || json.Status == "" || json.Size == "" || json.Aptitude == "" || json.Bio == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "都是必填的！",
		})
		return
	}
	cnt, err := model.GetThreadCountByName(json.Title)
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

	thread := &model.Thread{
		Title:    json.Title,
		Content:  json.Content,
		Sort:     json.Sort,
		Status:   json.Status,
		Tag:      json.Tag,
		Uid:      json.Uid,
		Bio:      json.Bio,
		Size:     json.Size,
		Aptitude: json.Aptitude,
		Thumb:    json.Thumb,
		Time:     time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04"),
	}

	if json.Identity.Hex() == "000000000000000000000000" {
		err = model.InsertThread(thread)
	} else {
		token := c.GetHeader("token")
		err = Auth(json.Uid, token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("%s", err),
			})
			return
		}
		err = model.UpdateThread(thread, json.Identity)
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

func ThreadDetail(c *gin.Context) {
	id := c.Param("id")
	// uc := u.(*helper.UserClaims)
	oid, _ := primitive.ObjectIDFromHex(id)
	thread, err := model.GetThreadByIdentity(oid)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": thread,
	})
}

func GetThreads(c *gin.Context) {
	sort := c.Query("sort")
	// oid, _ := primitive.ObjectIDFromHex(id)

	pageIndex, _ := strconv.ParseInt(c.Query("page"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	skip := (pageIndex - 1) * pageSize
	threads, err := model.GetThreads(&pageSize, &skip, sort)

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
		"data": threads,
	})
}
