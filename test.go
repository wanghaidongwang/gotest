package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	var bt bytes.Buffer
	config := &ssh.ClientConfig{
		User: "admin",
		Auth: []ssh.AuthMethod{
			ssh.Password("byd@1024"),
		},
	}
	client, err := ssh.Dial("tcp", "10.4.0.174:22", config)
	if err != nil {
		os.Exit(1)
	}
	session, err := client.NewSession()
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.Stdout = &bt
	if err := session.Run("ls"); err != nil {
		os.Exit(1)
	}
	//fmt.Println(bt.String()) //此处只能一次性读出结果，执行耗时长的命令时，半天无响应
}
