package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error


func main(){
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	// 下面这个表示我们读取一行，最后是以\n 为分割的，\n表示换行
	input,err = inputReader.ReadString('\n')
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("the input was:%s\n",input)
}
