package internal

import (
"bytes"
"github.com/dejavuzhou/felix/flx"
"github.com/dejavuzhou/felix/models"
"github.com/dejavuzhou/felix/utils"
"github.com/gin-gonic/gin"
"github.com/gorilla/websocket"
"github.com/sirupsen/logrus"
"net/http"
"strconv"
"time"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handle webSocket connection.
// first,we establish a ssh connection to ssh server when a webSocket comes;
// then we deliver ssh data via ssh connection between browser and ssh server.
// That is, read webSocket data from browser (e.g. 'ls' command) and send data to ssh server via ssh connection;
// the other hand, read returned ssh data from ssh server and write back to browser via webSocket API.
func WsSsh(c *gin.Context) {

	v, ok := c.Get("user")
	if !ok {
		logrus.Error("jwt token can't find auth user")
		return
	}
	userM, ok := v.(*models.User)
	if !ok {
		logrus.Error("context user is not a models.User type obj")
		return
	}
	cols, err := strconv.Atoi(c.DefaultQuery("cols", "120"))
	if wshandleError(c, err) {
		return
	}
	rows, err := strconv.Atoi(c.DefaultQuery("rows", "32"))
	if wshandleError(c, err) {
		return
	}
	idx, err := parseParamID(c)
	if wshandleError(c, err) {
		return
	}
	mc, err := models.MachineFind(idx)
	if wshandleError(c, err) {
		return
	}

	client, err := flx.NewSshClient(mc)
	if wshandleError(c, err) {
		return
	}
	defer client.Close()
	startTime := time.Now()
	ssConn, err := utils.NewSshConn(cols, rows, client)
	if wshandleError(c, err) {
		return
	}
	defer ssConn.Close()
	// after configure, the WebSocket is ok.
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if wshandleError(c, err) {
		return
	}
	defer wsConn.Close()

	quitChan := make(chan bool, 3)

	var logBuff = new(bytes.Buffer)

	// most messages are ssh output, not webSocket input
	go ssConn.ReceiveWsMsg(wsConn, logBuff, quitChan)
	go ssConn.SendComboOutput(wsConn, quitChan)
	go ssConn.SessionWait(quitChan)

	<-quitChan
	//write logs
	xtermLog := models.TermLog{
		EndTime:     time.Now(),
		StartTime:   startTime,
		UserId:      userM.ID,
		Log:         logBuff.String(),
		MachineId:   idx,
		MachineName: mc.Name,
		MachineIp:   mc.Ip,
		MachineHost: mc.Host,
		UserName:    userM.Username,
	}

	err = xtermLog.Create()
	if wshandleError(c, err) {
		return
	}
	logrus.Info("websocket finished")
}
