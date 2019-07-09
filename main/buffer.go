package main

import (
	"bytes"
	"fmt"
)

func main()  {
	/*
		s := []byte(" world")
		buf := bytes.NewBufferString("hello")
		fmt.Println(buf.String()) //buf转整形
		buf.Write(s)              //将s这个slice写到buf的尾部
		fmt.Println(buf.String())
		buf.WriteString(s) //将s这个string 写到buffer的尾部
	*/

	/*
		var s rune = '好'
		buf := bytes.NewBufferString("hello")
		fmt.Println(buf.String())
		buf.WriteRune(s)
		fmt.Println(buf.String())
	*/

	/*
		file, _ := os.Create("chenghuan.txt")
		buf := bytes.NewBufferString("水水水水谁谁谁水水水水")
		buf.WriteTo(file)
		fmt.Fprintf(file, buf.String())
	*/

	/*
		s1 := []byte("hello")
		buff := bytes.NewBuffer(s1)
		s2 := []byte(" world")
		buff.Write(s2)
		fmt.Println(buff.String())
		s3 := make([]byte, 3)
		buff.Read(s3)
		fmt.Println(buff.String())
		fmt.Println(string(s3))
		buff.Read(s3)
		fmt.Println(buff.String())
		fmt.Println(string(s3))
	*/
	/*
		buf := bytes.NewBufferString("hello")
		b, _ := buf.ReadByte()    //读取第一个byte，赋值给b
		fmt.Println(buf.String()) //打印ello，缓冲器头部第一个h被拿掉
		fmt.Println(string(b))    //打印h
	*/
	//ReadBytes和ReadByte根本就不是一回事，ReadBytes需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为byte类型的slice，返回后，缓冲器也会空掉一部分
	/*var d byte = "e"
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())
	b, _ := buf.ReadBytes(d)  //读到分隔符，并返回给b
	fmt.Println(buf.String()) //打印llo，缓冲器被取走一些数据
	fmt.Println(string(b))    //打印he
	*/
	/*file, _ := os.Open("chenghuan.txt")
	buff := bytes.NewBufferString("hello ")
	buff.ReadFrom(file) //将text里面的内容追加到缓冲器的尾部
	fmt.Println(buff.String())*/

	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String())
	b := buf.Next(2) //重头开始，取两个
	fmt.Println(buf.String())
	fmt.Println(string(b))

}
