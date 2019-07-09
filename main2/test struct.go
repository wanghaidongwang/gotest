package main

import "fmt"

func main ()  {
    wang := cli{
		"root",
		"xxxxxxxxxxxxxxxxxxxxx",
		 "192.168.100.50:22",
	}
    fmt.Println(wang)
}

type  cli struct {
	user    string
	pwd     string
	addr    string
}