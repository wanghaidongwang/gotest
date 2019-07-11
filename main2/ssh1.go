package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	//"os"
)

func shh1() {
	var user, password, address string
	user = "admin"
	password = "byd@1024"
	address = "10.4.0.174:22"
	//	fmt.Println("请输入,服务器IP:端口,账号,密码")
	//	fmt.Scanln(&address, &user, &password)
	//fmt.Scanln(&address)
	//fmt.Scanln(&user)
	//fmt.Scanln(&password)
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}
	client, err := ssh.Dial("tcp", address, &ssh.ClientConfig{
		//ssh是用TCP22端口?
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	ce(err, "dial")
	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()
	//fmt.Println(os.Stdout)
	//fmt.Println(session.Stdout)
	//var s string ="ls"
	//var str string = "ls"

	//var data []byte = []byte(str)
	var b bytes.Buffer

	//cmd := exec.Command("ls")
	session.Stdout = &b
	//fmt.Println(session.Stdout)
	session.Stderr = os.Stderr
	//session.Stdin = os.Stdin
	//fmt.Println(session.Stdout)

	//data1 = os.Stdout
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("linux", 32, 160, modes)
	_ = session.Run("ls ")
	ce(err, "request pty")
	err = session.Shell()
	ce(err, "start shell")
	err = session.Wait()
	ce(err, "return")

}
func main() {
	//var  user , password , address string
	//user = "admin"
	//password = "byd@1024"
	//address = "10.4.0.174:22"
	////	fmt.Println("请输入,服务器IP:端口,账号,密码")
	////	fmt.Scanln(&address, &user, &password)
	////fmt.Scanln(&address)
	////fmt.Scanln(&user)
	////fmt.Scanln(&password)
	//ce := func(err error, msg string) {
	//	if err != nil {
	//		log.Fatalf("%s error: %v", msg, err)
	//	}
	//}
	//client, err := ssh.Dial("tcp",address, &ssh.ClientConfig{
	//	User: user,
	//	Auth: []ssh.AuthMethod{ssh.Password(password)},
	//	//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	//})
	//ce(err, "dial")
	//session, err := client.NewSession()
	//ce(err, "new session")
	//defer session.Close()
	//
	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr
	//session.Stdin = os.Stdin
	////session.Run("ls /")
	////data1 = os.Stdout
	//modes := ssh.TerminalModes{
	//	ssh.ECHO: 1,
	//	ssh.TTY_OP_ISPEED: 14400,
	//	ssh.TTY_OP_OSPEED: 14400,
	//}
	//err = session.RequestPty("linux", 32, 160, modes)
	//ce(err, "request pty")
	//err = session.Shell()
	//ce(err, "start shell")
	//err = session.Wait()
	//ce(err, "return")
	//go  shh()
	shh1()
	//time.Sleep(5 * time.Second)
	// 当有请求访问ws时，执行此回调方法
	//http.HandleFunc("/wswang",wsHandler)
	//// 监听127.0.0.1:7777
	//err = http.ListenAndServe("0.0.0.0:8088", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe", err.Error())
	//}
	//go shh()

}
