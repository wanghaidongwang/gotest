package main

import (
	"bytes"
	 gossh "golang.org/x/crypto/ssh"
	"net"
)

type sssh struct {
	user    string
	pwd     string
	addr    string
	client  *gossh.Client
	session *gossh.Session
}


func (s *ssh) Connect() (*ssh, error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = s.user
	config.Auth = []gossh.AuthMethod{gossh.Password(s.pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key gossh.PublicKey) error { return nil }
	client, err := gossh.Dial("tcp", s.addr, config)
	if nil != err {
		return nil, err
	}
	s.client = client
	return s, nil
}


func (s *ssh) Exec(cmd string) (string, error) {
	var buf bytes.Buffer
	session, err := s.client.NewSession()
	if nil != err {
		return "", err
	}
	session.Stdout = &buf
	session.Stderr = &buf
	err = session.Run(cmd)
	if err != nil {
		return "", err
	}
	defer session.Close()
	stdout := buf.String()
	return stdout, nil
}