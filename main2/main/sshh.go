
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
    "net/http"
	"sync"
	"time"
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
const (
	wsMsgCmd    = "cmd"
	wsMsgResize = "resize"
)


type wsMsg struct {
	Type string `json:"type"`
	Cmd  string `json:"cmd"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

// connect to ssh server using ssh session.
type SshConn struct {
	// calling Write() to write data into ssh server
	StdinPipe  io.WriteCloser
	// Write() be called to receive data from ssh server
	ComboOutput *wsBufferWriter
	Session     *ssh.Session
}
// write data to WebSocket
// the data comes from ssh server.
type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}
func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

func NewSshConn(cols, rows int, sshClient *ssh.Client) (*SshConn, error) {
	sshSession, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	// we set stdin, then we can write data to ssh server via this stdin.
	// but, as for reading data from ssh server, we can set Session.Stdout and Session.Stderr
	// to receive data from ssh server, and write back to somewhere.
	stdinP, err := sshSession.StdinPipe()
	if err != nil {
		return nil, err
	}

	comboWriter := new(wsBufferWriter)
	//ssh.stdout and stderr will write output into comboWriter
	sshSession.Stdout = comboWriter
	sshSession.Stderr = comboWriter

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := sshSession.RequestPty("xterm", rows, cols, modes); err != nil {
		return nil, err
	}
	// Start remote shell
	if err := sshSession.Shell(); err != nil {
		return nil, err
	}
	return &SshConn{StdinPipe: stdinP, ComboOutput: comboWriter, Session: sshSession}, nil
}
func (ssConn *SshConn) SendComboOutput(wsConn *websocket.Conn, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)

	//every 120ms write combine output bytes into websocket response
	tick := time.NewTicker(time.Millisecond * time.Duration(120))
	//for range time.Tick(120 * time.Millisecond){}
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			//write combine output bytes into websocket response
			if err := flushComboOutput(ssConn.ComboOutput, wsConn); err != nil {
				logrus.WithError(err).Error("ssh sending combo output to webSocket failed")
				return
			}
		case <-exitCh:
			return
		}
	}
}

//ReceiveWsMsg  receive websocket msg do some handling then write into ssh.session.stdin
func (ssConn *SshConn) ReceiveWsMsg(wsConn *websocket.Conn, logBuff *bytes.Buffer, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)
	for {
		select {
		case <-exitCh:
			return
		default:
			//read websocket msg
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				logrus.WithError(err).Error("reading webSocket message failed")
				return
			}
			//unmashal bytes into struct
			msgObj := wsMsg{}
			if err := json.Unmarshal(wsData, &msgObj); err != nil {
				logrus.WithError(err).WithField("wsData", string(wsData)).Error("unmarshal websocket message failed")
			}
			switch msgObj.Type {
			case wsMsgCmd:
				//handle xterm.js stdin
				decodeBytes, err := base64.StdEncoding.DecodeString(msgObj.Cmd)
				if err != nil {
					logrus.WithError(err).Error("websock cmd string base64 decoding failed")
				}
				if _, err := ssConn.StdinPipe.Write(decodeBytes); err != nil {
					logrus.WithError(err).Error("ws cmd bytes write to ssh.stdin pipe failed")
				}
				//write input cmd to log buffer
				if _, err := logBuff.Write(decodeBytes); err != nil {
					logrus.WithError(err).Error("write received cmd into log buffer failed")
				}
			}
		}
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


func setQuit(ch chan bool) {
	ch <- true
}
