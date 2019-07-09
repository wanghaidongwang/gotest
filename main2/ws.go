package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main () {
	//当有请求访问ws时，执行此回调方法
	http.HandleFunc("/wswang", wsHandler)
	// 监听127.0.0.1:7777
	err := http.ListenAndServe("0.0.0.0:8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}


var (
	upgrader = websocket.Upgrader {
		// 读取存储空间大小
		ReadBufferSize:1024,
		// 写入存储空间大小
		WriteBufferSize:1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wbsCon *websocket.Conn
		err error
		data []byte
	)

	// 完成http应答，在httpheader中放下如下参数
	if wbsCon, err = upgrader.Upgrade(w, r, nil);err != nil {
		return // 获取连接失败直接返回
	}

	for {

		// 只能发送Text, Binary 类型的数据,下划线意思是忽略这个变量.
		if _, data, err = wbsCon.ReadMessage();err != nil {
			goto ERR // 跳转到关闭连接
		}
		if err = wbsCon.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR // 发送消息失败，关闭连接
		}

	}

ERR:
	// 关闭连接
	wbsCon.Close()
}