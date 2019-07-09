package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

func Connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth   []ssh.AuthMethod
		addr   string
		clientConfig *ssh.ClientConfig
		sshClient *ssh.Client
		sftpClient *sftp.Client
		err   error
	)
	// 将密码穿到验证方法切片里
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	//配置项
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		Timeout: 30 * time.Second,
		//这各参数是验证服务端的，返回nil可以不做验证，如果不设置会报错
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	//连接ip和端口
	addr = fmt.Sprintf("%s:%d", host, port)
	//通过tcp协议,连接ssh
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	//创建sftp服务对象
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	//返回sftp服务对象
	return sftpClient, nil
}

func main()  {
	Connect("admin","byd@1024","10.4.0.174",22)
}