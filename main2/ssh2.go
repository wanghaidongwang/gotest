package main
import (
	"bytes"
	"fmt"
	 "golang.org/x/crypto/ssh"
)
func main() {
	fmt.Println("helloworld")
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
		User: "admin",
		Auth: []ssh.AuthMethod{
			ssh.Password("byd@1024"),
		},
	}
	client, err := ssh.Dial("tcp", "10.4.0.174:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer  session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer

	session.Stdout = &b
	if err := session.Run("who"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}
