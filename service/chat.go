package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/fatih/set"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	CMD_SINGLE_MSG = 1
	CMD_ROOM_MSG   = 2
	CMD_HEART      = 0
)

type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"`           //消息ID
	Uid     string  `json:"uid,omitempty" form:"uid"`         //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Tid     string  `json:"tid,omitempty" form:"tid"`       //对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
}

//本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

//映射关系表
var clientMap map[string]*Node = make(map[string]*Node)

//读写锁
var rwlocker sync.RWMutex

//
// ws://127.0.0.1/chat?id=1&token=xxxx
func Chat(c *gin.Context) {

	uid := c.Query("uid")
	token := c.Query("token")
	
	isvalida := checkToken(uid, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//todo 获取用户全部群Id
	// comIds := contactService.SearchComunityIds(userId) // 群组
	// for _,v:=range comIds{
	// 	node.GroupSets.Add(v)
	// }
	//todo userid和node形成绑定关系
	rwlocker.Lock()
	clientMap[uid] = node
	rwlocker.Unlock()
	//todo 完成发送逻辑,con
	go sendproc(node)
	//todo 完成接收逻辑
	go recvproc(node)
	//
	sendMsg(uid, []byte("ok"))
}

//发送协程
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

//todo 添加新的群ID到用户的groupset中
func AddGroupId(userId, gid string) {
	//取得node
	rwlocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	//clientMap[userId] = node
	rwlocker.Unlock()
	//添加gid到set
}

//接收协程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//todo 对data进一步处理
		dispatch(data)
		fmt.Printf("recv<=%s", data)
	}
}

//后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		sendMsg(msg.Tid, data)
	case CMD_ROOM_MSG:
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.Tid) {
				v.DataQueue <- data
			}
		}
	case CMD_HEART:
		sendMsg(msg.Uid, []byte("ok"))
		//todo 一般啥都不做
	}
}

func sendMsg(userId string, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

//检测是否有效
func checkToken(userId string, token string) bool {
	//从数据库里面查询并比对
	return true
}
